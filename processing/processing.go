package processing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// reads the image file and returns an object
func ReadImage(path string) image.Image{
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// decodes the image using image libraries decode function
	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println(path)
		panic(err)
	}
	return img
}

// Writes the image object to a file
func WriteImage(path string , img image.Image){
	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// encoding the image to the new file, basically using png cause all of my images are in png
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		panic(err)
	}
}

// function to convert the image object to grayscale
func Grayscale(img image.Image) image.Image{
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// converting each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Min.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x ++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

// scales the image to 500x500 
func Resize(img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return resizedImg
}