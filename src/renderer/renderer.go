package renderer

import(
	"image"
	"math/cmplx"
	"utils/palette"
)

const resX int32 = 400
const resY int32 = 400
var scale float64 = 2.0
var center complex128 = 0 - 0i
const numWorkers int32 = 50
var aspectRatio float64 = float64(resX) / float64(resY)

type PixelJob struct{
	x int32
	y int32
}

type PixelResult struct{
	x int32
	y int32
	value float64
}

//Main function to be called from main to generate image
func RenderImage() *image.RGBA{
	const numJobs = resX * resY
	imgOut := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{int(resX), int(resY)}})
	jobsChan := make(chan PixelJob, numJobs)
	resultsChan := make(chan PixelResult, numJobs)
	for i := int32(0); i < numWorkers; i ++{
		go worker(jobsChan, resultsChan)
	}
	for x := int32(0); x < resX; x ++{
		for y := int32(0); y < resY; y ++{
			jobsChan <- PixelJob{x, y}
		}
	}
	close(jobsChan)
	for i := int32(0); i < numJobs; i ++{
		result := <- resultsChan
		c := palette.GetPaletteColor(result.value)
		imgOut.Set(int(result.x), int(result.y), c)
	}
	return imgOut
}

// Worker function for creating a worker pool to render indvidual pixels
func worker(jobs <-chan PixelJob, results chan <- PixelResult){
	for job := range jobs{
		value := iteratePixel(job.x, job.y)
		results <- PixelResult{job.x, job.y, value}
	}
}

// Maps a pixel to a point on the complex plane given a center and view
func mapPixel(x, y int32) complex128{
	r := (real(center) - scale * aspectRatio / 2.0) + float64(x) / float64(resX) * scale * aspectRatio
	i := (imag(center) - scale / 2.0) + float64(y) / float64(resY) * scale
	return complex(r, i)
}

//Iterates over z = z^2 + c and returns the escape value given a pixel point (x, y)
func iteratePixel(x, y int32) float64{
	maxIter := 256
	z := 0 + 0i
	c := mapPixel(x, y)
	i := 0
	for ; cmplx.Abs(z) < 2.0 && i < maxIter; i ++{
		z = z * z + c
	}
	return float64(i) / float64(maxIter)
}