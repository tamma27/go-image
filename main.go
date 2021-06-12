package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"

	"golang.org/x/image/webp"
)

func main() {
	Png2jpg("./image/image.png", "./image.png.jpg")
	Watermar("./image.png.jpg")

	Webp2jpg("./image/image.webp", "./image.webp.jpg")
	Watermar("./image.webp.jpg")

	Watermar("./image/image.jpg")
}

func Webp2jpg(orignImage, newImage string) (err error) {

	webpImgFile, err := os.Open(orignImage)
	if err != nil {
		return
	}
	defer webpImgFile.Close()

	imgSrc, err := webp.Decode(webpImgFile)
	if err != nil {
		return
	}

	return convert2jpg(imgSrc, newImage)
}

func Png2jpg(orignImage, newImage string) (err error) {

	pngImgFile, err := os.Open(orignImage)
	if err != nil {
		return
	}

	defer pngImgFile.Close()

	imgSrc, err := png.Decode(pngImgFile)
	if err != nil {
		return
	}

	return convert2jpg(imgSrc, newImage)
}

func convert2jpg(imgSrc image.Image, pathImage string) (err error) {

	// create a new Image with the same dimension of PNG image
	newImg := image.NewRGBA(imgSrc.Bounds())

	// we will use white background to replace PNG's transparent background
	// you can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// paste image OVER to newImage
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	// create new out JPEG file
	jpgImgFile, err := os.Create(pathImage)
	if err != nil {
		return
	}

	defer jpgImgFile.Close()

	var opt jpeg.Options
	opt.Quality = 80

	// convert newImage to JPEG encoded byte and save to jpgImgFile
	// with quality = 80
	err = jpeg.Encode(jpgImgFile, newImg, &opt)

	//err = jpeg.Encode(jpgImgFile, newImg, nil) -- use nil if ignore quality options

	if err != nil {
		return
	}

	return
}

func Watermar(orignImage string) {
	img, err := mergi.Import(impexp.NewFileImporter(orignImage))
	if err != nil {
		return
	}

	w := img.Bounds().Max.X / 4

	watermarkImage, _ := mergi.Import(impexp.NewFileImporter("./image/logo.png"))
	newWatermarkImage, _ := mergi.Resize(watermarkImage, uint(w), uint(w))
	opecWatermark, _ := mergi.Opacity(newWatermarkImage, 0.6)

	res, err := mergi.Watermark(opecWatermark, img, image.Pt(20, 20))
	if err != nil {
		return
	}

	mergi.Export(impexp.NewFileExporter(res, orignImage))
}
