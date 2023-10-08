package ImageToAscii

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
	"testing"
)

var asciiChar = " MND8OZ$7I?+=~:,.."

func asciiArt(img image.Image, w, h int) []byte {
	table := []byte(asciiChar)
	buffer := new(bytes.Buffer)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			g := color.GrayModel.Convert(img.At(j, i))
			y := reflect.ValueOf(g).FieldByName("Y").Uint()
			pos := int(y * 16 / 255)
			_ = buffer.WriteByte(table[pos])
		}
		_ = buffer.WriteByte('\n')
	}
	return buffer.Bytes()
}
func getHeight(img image.Image, w int) (image.Image, int, int) {
	sz := img.Bounds()
	height := (sz.Max.Y * w * 10) / (sz.Max.X * 16)
	img = resize.Resize(uint(w), uint(height), img, resize.Lanczos3)
	return img, w, height
}

func TestImageToAscii(t *testing.T) {
	imageFile, err := os.Open("koala.png")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	defer imageFile.Close()

	loadedImage, err := png.Decode(imageFile)
	if err != nil {
		fmt.Println("Couldn't decode image,\nError:", err.Error())
		return
	}

	finalASCIIArt := asciiArt(getHeight(loadedImage, 200))
	fmt.Println(string(finalASCIIArt))

}
