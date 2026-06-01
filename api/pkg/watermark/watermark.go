package watermark

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func AddTextToImage(data []byte, text string) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, "", fmt.Errorf("decode image: %w", err)
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	// 计算水印位置：右下角
	face := basicfont.Face7x13
	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(color.RGBA{200, 200, 200, 160}),
		Face: face,
	}

	width := font.MeasureString(face, text).Round()
	x := bounds.Dx() - width - 10
	y := bounds.Dy() - 10
	if x < 0 {
		x = 10
	}

	d.Dot = fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y),
	}
	d.DrawString(text)

	var buf bytes.Buffer
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, rgba, &jpeg.Options{Quality: 85})
	default:
		err = png.Encode(&buf, rgba)
	}
	if err != nil {
		return nil, "", fmt.Errorf("encode image: %w", err)
	}

	return buf.Bytes(), format, nil
}
