package renderer

import(
	"image"
	"math/cmplx"
	"utils/palette"
)

//======================================================================================
//                                     CONFIG                                           
//======================================================================================
const numWorkers int = 50
const resX int = 1920
const resY int = 1080
const maxIter int = 5000
var scale float64 = 0.01
var center complex128 = -0.761574 - 0.0847596i
var aspectRatio float64 = float64(resX) / float64(resY)
// =====================================================================================

type PixelJob struct{
	x int
	y int
}

type PixelResult struct{
	x int
	y int
	iter int
}

//Main function to be called from main to generate image
func RenderImage() *image.RGBA{
	hist := [maxIter + 1]int{}
	iterCounts := [resX * resY]int{}
	const numJobs int = resX * resY
	imgOut := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{int(resX), int(resY)}})
	jobsChan := make(chan PixelJob, numJobs)
	resultsChan := make(chan PixelResult, numJobs)
	for i := 0; i < numWorkers; i ++{
		go worker(jobsChan, resultsChan)
	}

	// First pass to get iteration counts
	for x := 0; x < resX; x ++{
		for y := 0; y < resY; y ++{
			jobsChan <- PixelJob{x, y}
		}
	}
	close(jobsChan)
	for i := 0; i < numJobs; i ++{
		result := <- resultsChan
		index := result.x + resX * result.y
		iterCounts[index] = result.iter
		hist[result.iter] ++
	}

	// Second pass to get a normalized histogram
	runningTotal := [len(hist)]int{}
	n := 0
	for i := 0; i < len(hist); i ++{
		n += hist[i]
		runningTotal[i] = n
	}

	//Third pass to build image based on the normalized historgram value
	for x := 0; x < resX; x ++{
		for y := 0; y < resY; y ++{
			i := flatten(x, y)
			/*
			val := 0.0
			if iterCounts[i] == maxIter{
				val = 1.0
			} else{
				val = float64(runningTotal[iterCounts[i]]) / float64(runningTotal[len(runningTotal) - 2])
			}
			*/
			val := float64(iterCounts[i]) / float64(maxIter)
			c := palette.GetPaletteColor(val)
			imgOut.Set(x, y, c)
		}
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

// Flattens a 2D array
func flatten(x, y int) int{
	return x + resX * y
}

// Maps a pixel to a point on the complex plane given a center and view
func mapPixel(x, y int) complex128{
	r := (real(center) - scale * aspectRatio / 2.0) + float64(x) / float64(resX) * scale * aspectRatio
	i := (imag(center) - scale / 2.0) + float64(y) / float64(resY) * scale
	return complex(r, i)
}

//Iterates over z = z^2 + c and returns the escape value given a pixel point (x, y)
func iteratePixel(x, y int) int{
	z := 0 + 0i
	c := mapPixel(x, y)
	i := 0
	for ; cmplx.Abs(z) < 1024.0 && i < maxIter; i ++{
		z = z * z + c
	}
	return i
}
