package palette

import (
	"image"
	"image/color"
	"math"
)

var setColor color.RGBA = color.RGBA{0, 0, 0, 255}

var palette [16]color.RGBA = [16]color.RGBA{
		color.RGBA{60, 30, 15, 255},
		color.RGBA{25, 7, 26, 255},
		color.RGBA{9, 1, 47, 255},
		color.RGBA{4, 4, 73, 255},
		color.RGBA{0, 7, 100, 255},
		color.RGBA{12, 44, 138, 255},

		color.RGBA{24, 82, 177, 255},
		color.RGBA{57, 125, 209, 255},
		color.RGBA{134, 181, 229, 255},
		color.RGBA{211, 236, 248, 255},
		color.RGBA{241, 233, 191, 255},
		color.RGBA{248, 201, 95, 255},

		color.RGBA{255, 170, 0, 255},
		color.RGBA{204, 128, 0, 255},
		color.RGBA{153, 87, 0, 255},
		color.RGBA{106, 52, 3, 255},
}

// returns a color by lerping through a palette of colors based on a parameter value
func GetPaletteColor(val float64) color.RGBA{
	if val == 1.0{
		return setColor
	}	
	val = math.Pow(val, 0.3)
	index := int(val * float64(len(palette)))
	/*
	frac := (val * float64(len(palette))) - float64(index)
	next := index + 1
	if next >= len(palette){
		next = 0
	}
	c0 := palette[index]
	c1 := palette[next]
	dr := float64(c1.R) / 255.0 - float64(c0.R) / 255.0
	dg := float64(c1.G) / 255.0 - float64(c0.G) / 255.0
	db := float64(c1.B) / 255.0 - float64(c0.B) / 255.0
	red := float64(c0.R) / 255.0 + dr * frac
	green := float64(c0.B) / 255.0 + db * frac
	blue := float64(c0.G) / 255.0 + dg * frac
	return color.RGBA{uint8(math.Round(red * 255.0)), uint8(math.Round(green * 255.0)), uint8(math.Round(blue * 255.0)), 255}
	*/
	return palette[index]
}

// ===============================================================================
//                                  TESTING                                       
// ===============================================================================
func TestGradientImage() *image.RGBA{
	resX := 1000
	resY := 400
	imgOut := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{int(resX), int(resY)}})
	for x := 0; x < resX; x ++{
		for y := 0; y < resY; y ++{
			c := GetPaletteColor(float64(x) / float64(resX))
			imgOut.Set(x, y, c)
		}
	}
	return imgOut
}