package main

import(
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const resX int = 1920 * 4 //4k image
const resY int = 1080 * 4
const numWorkers = 50
var aspectRatio float64 = float64(resX) / float64(resY)
var scale float64 = 0.01
var center complex128 = -0.761574 - 0.0847596i

// =========================================================== Thread Pool ==============================================================

type PixelJob struct{
	x int
	y int
}

type PixelResult struct{
	x int
	y int
	value float64
}

func worker(jobs <-chan PixelJob, results chan <- PixelResult){
	/* Worker function for creating a worker pool to render indvidual pixels */
	for job := range jobs{
		value := iteratePixel(job.x, job.y)
		results <- PixelResult{job.x, job.y, value}
	}
}

// =========================================================== Rendering =================================================================

func abs(n complex128) float64{
	/* Returns the absolute value of a complex number */
	return math.Sqrt(real(n) * real(n) + imag(n) * imag(n))
}

func iteratePixel(x, y int) float64{
	/* 	Iterates over z = z^2 + c and returns the escape value given a pixel point (x, y) */
	maxIter := 2056
	z := 0 + 0i
	c := mapPixel(x, y)
	i := 0
	for ; abs(z) < 2 && i < maxIter; i ++{
		z = z * z + c
	}
	return float64(i) / float64(maxIter)
}

func mapPixel(x, y int) complex128{
	/* Maps a pixel to a point on the complex plane given a center and view */
	r := (real(center) - scale * aspectRatio / 2.0) + float64(x) / float64(resX) * scale * aspectRatio
	i := (imag(center) - scale / 2.0) + float64(y) / float64(resY) * scale
	return complex(r, i)
}

func getPaletteColor(val float64) color.RGBA{
		/* returns a color by lerping through a palette of colors based on a parameter value*/
		if val == 1.0{
			return color.RGBA{0, 0, 0, 255} // set color
		}
		val = math.Sqrt(val) // gamma correction
		palette := [...]color.RGBA{
			color.RGBA{255, 0 , 0, 255}, // bright red
			color.RGBA{0, 255, 0, 255}, // bright green
			color.RGBA{0, 0, 255, 255}, // bright blue
		}
		index := int(val * float64(len(palette) - 1))
		frac := val * float64(len(palette) - 1) - float64(index)
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
		return color.RGBA{uint8(red * 255.0), uint8(green * 255.0), uint8(blue * 255.0), 255}
	}

// =========================================================== Main =====================================================================

func main(){
	imgOut := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{resX, resY}})
	const numJobs = resX * resY
	jobsChan := make(chan PixelJob, numJobs)
	resultsChan := make(chan PixelResult, numJobs)
	for i := 0; i < numWorkers; i ++{
		go worker(jobsChan, resultsChan)
	}
	for x := 0; x < resX; x ++{
		for y := 0; y < resY; y ++{
			jobsChan <- PixelJob{x, y}
		}
	}
	close(jobsChan)
	for i := 0; i < numJobs; i ++{
		result := <- resultsChan
		c := getPaletteColor(result.value)
		imgOut.Set(result.x, result.y, c)
	}
	f, _ := os.Create("../output.png")
	png.Encode(f, imgOut)
}
