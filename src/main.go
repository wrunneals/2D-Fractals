package main

import(
	"image/png"
	"os"
	"utils/renderer"
)

func main(){
	imgOut := renderer.RenderImage()
	f, _ := os.Create("../output.png")
	png.Encode(f, imgOut)
}
