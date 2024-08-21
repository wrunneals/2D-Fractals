package main

import(
	"fmt"
	"image/png"
	"os"
	"utils/renderer"
)

var center complex128 = -0.1048976440 + 0.9278883362i

func main(){
	frames := 1
	scale := 1.5625e-06
	center := -0.1048976440 + 0.9278883362i
	for i := 0; i < frames; i++{
		imgOut := renderer.RenderImage(scale, center)
		f, _ := os.Create(fmt.Sprintf("../output%d.png", i))
		png.Encode(f, imgOut)
		fmt.Println("Successfully rendered output", i, ":\n    Center: ", center, "\n    Scale: ", scale)
		scale = scale / 2.0
	}
}

