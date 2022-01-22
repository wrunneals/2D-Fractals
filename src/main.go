package main

import(
	"image/png"
	"os"
	"utils/renderer"
	// "utils/palette"
)

func main(){
	// TESTING imgOut := palette.TestGradientImage()
	imgOut := renderer.RenderImage()
	f, _ := os.Create("../output.png")
	png.Encode(f, imgOut)
}
