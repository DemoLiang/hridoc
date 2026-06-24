package watermark

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	pdffont "github.com/pdfcpu/pdfcpu/pkg/font"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

const (
	ModeHorizontal = "horizontal"
	ModeDiagonal   = "diagonal"
)

var (
	defaultFontPath string
	fontFaces       = map[int]font.Face{}
	fontDataCache   []byte
)

type Options struct {
	Text     string
	Mode     string
	Color    string
	Opacity  float64
	FontSize int
}

func (o *Options) normalize() {
	if o.Mode == "" {
		o.Mode = ModeHorizontal
	}
	if o.Color == "" {
		o.Color = "#D0E0FF"
	}
	if o.Opacity <= 0 || o.Opacity > 1 {
		o.Opacity = 0.05
	}
	if o.FontSize <= 0 {
		o.FontSize = 44
	}
}

func InitFont(fontPath string) error {
	if fontPath == "" {
		fontPath = "./fonts/simhei.ttf"
	}
	defaultFontPath = fontPath
	data, err := os.ReadFile(fontPath)
	if err != nil {
		return fmt.Errorf("read font file: %w", err)
	}
	fontDataCache = data
	return loadFace(44)
}

func loadFace(size int) error {
	if fontDataCache == nil {
		return fmt.Errorf("font data not loaded")
	}
	ft, err := opentype.Parse(fontDataCache)
	if err != nil {
		return fmt.Errorf("parse font: %w", err)
	}
	face, err := opentype.NewFace(ft, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return fmt.Errorf("create font face: %w", err)
	}
	fontFaces[size] = face
	return nil
}

func getFontFace(size int) font.Face {
	if size <= 0 {
		size = 44
	}
	if face, ok := fontFaces[size]; ok {
		return face
	}
	if fontDataCache == nil {
		_ = InitFont("")
	}
	if fontDataCache != nil {
		_ = loadFace(size)
		if face, ok := fontFaces[size]; ok {
			return face
		}
	}
	return nil
}

func parseHexColor(hex string) (color.RGBA, error) {
	hex = strings.TrimSpace(hex)
	hex = strings.TrimPrefix(hex, "#")
	switch len(hex) {
	case 3:
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	case 4:
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2], hex[3], hex[3]})
	case 6, 8:
		// ok
	default:
		return color.RGBA{}, fmt.Errorf("invalid hex color: %s", hex)
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	var a uint64 = 255
	if len(hex) == 8 {
		a, err = strconv.ParseUint(hex[6:8], 16, 8)
		if err != nil {
			return color.RGBA{}, err
		}
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}, nil
}

func applyOpacity(c color.RGBA, opacity float64) color.RGBA {
	if opacity < 0 {
		opacity = 0
	}
	if opacity > 1 {
		opacity = 1
	}
	a := float64(c.A) * opacity
	if a < 0 {
		a = 0
	}
	if a > 255 {
		a = 255
	}
	alpha := a / 255.0
	return color.RGBA{
		R: uint8(float64(c.R) * alpha),
		G: uint8(float64(c.G) * alpha),
		B: uint8(float64(c.B) * alpha),
		A: uint8(a),
	}
}

func nrgbaToRGBA(src *image.NRGBA) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dst.Set(x, y, src.NRGBAAt(x, y))
		}
	}
	return dst
}

func measureText(face font.Face, text string) (width, height int) {
	width = font.MeasureString(face, text).Round()
	metrics := face.Metrics()
	height = (metrics.Ascent + metrics.Descent).Round()
	if height < 1 {
		height = 1
	}
	if width < 1 {
		width = 1
	}
	return
}

func drawTextTile(face font.Face, text string, src image.Image, wmColor color.RGBA) *image.RGBA {
	bounds := src.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, src, bounds.Min, draw.Src)

	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(wmColor),
		Face: face,
	}
	width, height := measureText(face, text)

	stepX := int(float64(width) * 3.5)
	stepY := int(float64(height) * 5.0)
	if stepX < width*2 {
		stepX = width * 2
	}
	if stepY < height*3 {
		stepY = height * 3
	}

	for y := -height; y < bounds.Dy()+height; y += stepY {
		row := (y + height) / stepY
		for x := -width; x < bounds.Dx()+width; x += stepX {
			dx := x
			if row%2 != 0 {
				dx += stepX / 2
			}
			d.Dot = fixed.Point26_6{
				X: fixed.I(dx),
				Y: fixed.I(y + face.Metrics().Ascent.Round()),
			}
			d.DrawString(text)
		}
	}
	return rgba
}

func drawDiagonalTile(face font.Face, text string, src image.Image, wmColor color.RGBA) *image.RGBA {
	bounds := src.Bounds()
	wmLayer := image.NewRGBA(bounds)

	width, height := measureText(face, text)
	pad := int(float64(width) * 0.6)
	if pad < 20 {
		pad = 20
	}
	tileW := width + pad*2
	tileH := height + pad*2

	textCanvas := image.NewRGBA(image.Rect(0, 0, tileW, tileH))
	d := &font.Drawer{
		Dst:  textCanvas,
		Src:  image.NewUniform(wmColor),
		Face: face,
	}
	d.Dot = fixed.Point26_6{
		X: fixed.I(pad),
		Y: fixed.I(pad + face.Metrics().Ascent.Round()),
	}
	d.DrawString(text)

	rotatedNRGBA := imaging.Rotate(textCanvas, -45, color.Transparent)
	rotated := nrgbaToRGBA(rotatedNRGBA)
	rotBounds := rotated.Bounds()

	stepX := int(float64(rotBounds.Dx()) * 1.6)
	stepY := int(float64(rotBounds.Dy()) * 1.6)
	if stepX < rotBounds.Dx()+20 {
		stepX = rotBounds.Dx() + 20
	}
	if stepY < rotBounds.Dy()+20 {
		stepY = rotBounds.Dy() + 20
	}

	for y := -rotBounds.Dy(); y < bounds.Dy()+rotBounds.Dy(); y += stepY {
		row := (y + rotBounds.Dy()) / stepY
		for x := -rotBounds.Dx(); x < bounds.Dx()+rotBounds.Dx(); x += stepX {
			dx := x
			if row%2 != 0 {
				dx += stepX / 2
			}
			draw.Draw(wmLayer, image.Rect(dx, y, dx+rotBounds.Dx(), y+rotBounds.Dy()), rotated, rotBounds.Min, draw.Over)
		}
	}

	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, src, bounds.Min, draw.Src)
	draw.Draw(rgba, bounds, wmLayer, bounds.Min, draw.Over)
	return rgba
}

func AddTextToImage(data []byte, opts *Options) ([]byte, string, error) {
	opts.normalize()

	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, "", fmt.Errorf("decode image: %w", err)
	}

	face := getFontFace(opts.FontSize)
	if face == nil {
		return data, format, nil
	}

	c, err := parseHexColor(opts.Color)
	if err != nil {
		c = color.RGBA{208, 224, 255, 255}
	}
	c = applyOpacity(c, opts.Opacity)

	var rgba *image.RGBA
	if opts.Mode == ModeDiagonal {
		rgba = drawDiagonalTile(face, opts.Text, img, c)
	} else {
		rgba = drawTextTile(face, opts.Text, img, c)
	}

	var buf bytes.Buffer
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, rgba, &jpeg.Options{Quality: 92})
	default:
		err = png.Encode(&buf, rgba)
	}
	if err != nil {
		return nil, "", fmt.Errorf("encode image: %w", err)
	}

	return buf.Bytes(), format, nil
}

func AddTextToPDF(data []byte, opts *Options) ([]byte, error) {
	opts.normalize()

	conf := model.NewDefaultConfiguration()

	fontPath := defaultFontPath
	if fontPath == "" {
		fontPath = "./fonts/simhei.ttf"
	}
	if _, err := os.Stat(fontPath); err == nil {
		fontData, err := os.ReadFile(fontPath)
		if err == nil {
			_ = pdffont.InstallFontFromBytes(pdffont.UserFontDir, "simhei", fontData)
		}
	}

	color := opts.Color
	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}
	opacity := opts.Opacity
	if opacity <= 0 || opacity > 1 {
		opacity = 0.05
	}

	desc := fmt.Sprintf("fontname:simhei, points:%d, diagonal:1, opacity:%.2f, color:%s", opts.FontSize, opacity, color)
	if opts.Mode == ModeHorizontal {
		desc = fmt.Sprintf("fontname:simhei, points:%d, opacity:%.2f, color:%s, position:c", opts.FontSize, opacity, color)
	}
	wm, err := api.TextWatermark(opts.Text, desc, true, false, types.POINTS)
	if err != nil {
		return nil, fmt.Errorf("create pdf watermark: %w", err)
	}

	in := bytes.NewReader(data)
	var out bytes.Buffer
	if err := api.AddWatermarks(in, &out, nil, wm, conf); err != nil {
		return nil, fmt.Errorf("add pdf watermark: %w", err)
	}
	return out.Bytes(), nil
}

func ApplyIfNeeded(data []byte, fileType string, opts *Options) ([]byte, string, error) {
	if opts == nil {
		return data, fileType, nil
	}
	opts.normalize()
	if opts.Text == "" {
		return data, fileType, nil
	}
	switch fileType {
	case "image", "jpeg", "jpg", "png":
		return AddTextToImage(data, opts)
	case "pdf":
		out, err := AddTextToPDF(data, opts)
		return out, fileType, err
	default:
		return data, fileType, nil
	}
}

func PreviewSample(opts *Options) ([]byte, error) {
	opts.normalize()
	img := image.NewRGBA(image.Rect(0, 0, 1200, 900))
	for y := 0; y < 900; y++ {
		for x := 0; x < 1200; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 95}); err != nil {
		return nil, err
	}
	out, _, err := AddTextToImage(buf.Bytes(), opts)
	return out, err
}

func ReadAll(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
