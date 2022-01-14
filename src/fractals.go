package main

import(
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var resX int = 400
var resY int = 400

var viewWidth float64 = 2.0
var viewHeight float64 = 2.0

var center complex128 = -0.5 + 0i

func abs(n complex128) float64{
	return math.Sqrt(real(n) * real(n) + imag(n) * imag(n))
}

func iteratePixel(x, y int) float64{
	maxIter := 100
	z := 0 + 0i
	c := mapPixel(x, y)
	i := 0
	for ; abs(z) < 2 && i < maxIter; i ++{
		z = z * z + c
	}
	return float64(i) / float64(maxIter)
}

func mapPixel(x, y int) complex128{
	r := (real(center) - viewWidth / 2.0) + float64(x) / float64(resX) * viewWidth
	i := (imag(center) - viewHeight / 2.0) + float64(y) / float64(resY) * viewHeight
	return complex(r, i)
}

func main(){
	imgOut := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{resX, resY}})
	
	for x := 0; x < resX; x ++{
		for y := 0; y < resY; y ++{
			esc := uint8(iteratePixel(x, y) * 255)
			c := color.RGBA{esc, esc, esc, 255}
			imgOut.Set(x, y, c)
		}
	}

	f, _ := os.Create("../output.png")
	png.Encode(f, imgOut)
}