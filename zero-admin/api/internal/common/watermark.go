// package main

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/fogleman/gg"
// )

// func MakeWaterMarker(imgPath string, waterDesc string, outPath string) bool {
// 	im, err := gg.LoadImage(imgPath)
// 	if err != nil {
// 		fmt.Print(err)
// 		return false
// 	}
// 	w := im.Bounds().Size().X
// 	h := im.Bounds().Size().Y
// 	fmt.Print(w, h)
// 	dc := gg.NewContext(w, h)
// 	rd := w / 375
// 	if rd == 0 {
// 		rd = 1
// 	}
// 	fontFile := "/Library/Fonts/Arial Unicode.ttf"

// 	if err := dc.LoadFontFace(fontFile, float64(rd*24)); err != nil {
// 		fmt.Print(err)
// 		return false
// 	}
// 	dc.DrawImage(im, 0, 0)
// 	sw, sh := dc.MeasureString(waterDesc)
// 	chs := strings.Split(waterDesc, "")
// 	dc.SetRGBA(99, 99, 99, 0.6)
// 	for idx := 0; idx < len(chs); idx++ {
// 		for j := 0; j < h; j += (int(sh + float64(30*rd))) {
// 			for i := 0; i < w; i += (int(sw + float64(24*rd))) {
// 				dc.Push()
// 				// for idx := 0; idx < len(chs); idx++ {
// 				// dc.DrawString(chs[0], float64(i), float64(j))
// 				dc.DrawString(chs[idx], float64(i), float64(j))

// 				// }
// 				dc.Pop()
// 			}
// 		}
// 	}
// 	err = dc.SavePNG(outPath)
// 	fmt.Printf("err:%v", err)
// 	return true
// }

// func main() {
// 	MakeWaterMarker("./22.jpg", "测试水印", "./abc.jpg")
// }

package common

import (
	"fmt"
	"strings"

	"github.com/fogleman/gg"
)

func MakeWaterMarker(imgPath string, waterDesc string, outPath string) bool {
	im, err := gg.LoadImage(imgPath)
	if err != nil {
		fmt.Print(err)
		return false
	}
	// 图片宽
	w := im.Bounds().Size().X
	// 图片高
	h := im.Bounds().Size().Y
	fmt.Println("图片宽：", w, "图片高：", h)

	dc := gg.NewContext(w, h)
	rd := w / 375
	if rd == 0 {
		rd = 1
	}

	fontFile := "/Library/Fonts/Arial Unicode.ttf"
	if err := dc.LoadFontFace(fontFile, float64(rd*24)); err != nil {
		fmt.Print(err)
		return false
	}

	dc.DrawImage(im, 0, 0)
	// 文字宽 高
	sw, sh := dc.MeasureString(waterDesc)

	// 文字切分
	chs := strings.Split(waterDesc, "")
	dc.SetRGBA(99, 99, 99, 0.6)

	// 取第 1 2 3 4 个文字
	for idx := 0; idx < len(chs); idx++ {

		// 绘制的每个字的位置 图片高度/文字高度*文字数+文字位置*文字高度+60
		dh := float64((h))/float64(sh)*float64(len(chs)) + float64(idx)*sh + 60
		dw := float64((w))/float64(sw)*float64(len(chs)) + float64(idx)*sw + 60
		// if idx == 0 {
		// 	dh += 10
		// 	dw += 10
		// }
		// for hIdx := 0; hIdx < h; hIdx += (int(sh + float64(30*rd))) {
		// for wIdx := 0; wIdx < w; wIdx += (int(sw + float64(24*rd))) {
		dc.Push()
		// for idx := 0; idx < len(chs); idx++ {
		// dc.DrawString(chs[0], float64(i), float64(j))
		dc.DrawString(chs[idx], float64(dw), float64(dh))

		// }
		dc.Pop()
		// }
		// }
	}
	err = dc.SavePNG(outPath)
	fmt.Printf("err:%v", err)
	return true
}

func main() {
	MakeWaterMarker("./22.jpg", "测试水印", "./abc.jpg")
}
