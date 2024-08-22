package main

import(
	"fmt"
	"image/png"
	"os"
	"utils/renderer"
)

func main(){
	frames := 1
	scale := 0.000375
	center := -0.749348806 + 0.038702074i
	for i := 0; i < frames; i++{
		imgOut := renderer.RenderImage(scale, center)
		f, _ := os.Create(fmt.Sprintf("../output%d.png", i))
		png.Encode(f, imgOut)
		fmt.Println("Successfully rendered output", i, ":\n    Center: ", center, "\n    Scale: ", scale)
		scale = scale / 2.0
	}
}
