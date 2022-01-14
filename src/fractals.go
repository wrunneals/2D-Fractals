package main

import(
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const resX int = 1920
const resY int = 1080
const numWorkers = 10

var aspectRatio float64 = float64(resX) / float64(resY)
var scale float64 = 2.0
var center complex128 = -0.5 + 0i

type PixelJob struct{
	x int
	y int
}

type PixelResult struct{
	x int
	y int
	value float64
}

func abs(n complex128) float64{
	/* Returns the absolute value of a complex number */
	return math.Sqrt(real(n) * real(n) + imag(n) * imag(n))
}


func iteratePixel(x, y int) float64{
	/* 	Iterates over z = z^2 + c and returns the escape value given a pixel point (x, y) */
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
	/* Maps a pixel to a point on the complex plane given a center and view */
	r := (real(center) - scale * aspectRatio / 2.0) + float64(x) / float64(resX) * scale * aspectRatio
	i := (imag(center) - scale / 2.0) + float64(y) / float64(resY) * scale
	return complex(r, i)
}


func worker(jobs <-chan PixelJob, results chan <- PixelResult){
	/* Worker function for creating a worker pool to render indvidual pixels */
	for job := range jobs{
		value := iteratePixel(job.x, job.y)
		results <- PixelResult{job.x, job.y, value}
	}
}

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
		esc := uint8(result.value * 255.0)
		c := color.RGBA{esc, esc, esc, 255.0}
		imgOut.Set(result.x, result.y, c)
	}

	f, _ := os.Create("../output.png")
	png.Encode(f, imgOut)
}