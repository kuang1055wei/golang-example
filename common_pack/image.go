package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	//var text = "我是文字"
	fName := "./test.jpg"
	var f *os.File
	res, _ := os.Stat(fName)
	if res == nil {
		f, _ = os.Create(fName)
	} else {
		f, _ = os.OpenFile(fName, os.O_WRONLY, os.ModePerm)
	}
	defer f.Close()
	r := image.Rect(0, 0, 1000, 1000)
	rgbImg := image.NewRGBA(r)
	colorM := color.RGBA{R: 0, G: 100, B: 100, A: 255}
	//rgbImg整个图片涂颜色，rgbImg.Bounds()获取矩形的界限
	draw.Draw(rgbImg, rgbImg.Bounds(), &image.Uniform{colorM}, image.Point{
		X: 0,
		Y: 0,
	}, draw.Src)
	//在图中间画图
	draw.Draw(rgbImg, image.Rect(400, 400, 500, 500),
		&image.Uniform{color.RGBA{
			R: 100,
			G: 100,
			B: 100,
			A: 255,
		}}, image.Point{
			X: 0,
			Y: 0,
		}, draw.Src)

	png.Encode(f, rgbImg)
}
