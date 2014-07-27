package image

import (
	"image"
	"image/draw"
	"runtime"
	"sync"

	"github.com/pierrre/mandelbrot"
)

func Render(im draw.Image, proj Projection, maxIter int, colorizer Colorizer) {
	render(im, im.Bounds(), proj, maxIter, colorizer)
}

func RenderWorker(im draw.Image, proj Projection, maxIter int, colorizer Colorizer, workerCount int) {
	size := im.Bounds().Size()
	width := size.X
	height := size.Y

	wg := new(sync.WaitGroup)
	wg.Add(workerCount)

	for w := 0; w < workerCount; w++ {
		minY := int(float64(height) * (float64(w) / float64(workerCount)))
		maxY := int(float64(height) * (float64(w+1) / float64(workerCount)))
		wBounds := image.Rect(0, minY, width, maxY)

		go func(wBounds image.Rectangle) {
			render(im, wBounds, proj, maxIter, colorizer)
			wg.Done()
		}(wBounds)
	}

	wg.Wait()
}

func RenderWorkerAuto(im draw.Image, proj Projection, maxIter int, colorizer Colorizer) {
	RenderWorker(im, proj, maxIter, colorizer, runtime.GOMAXPROCS(0)*4)
}

func render(im draw.Image, bounds image.Rectangle, proj Projection, maxIter int, colorizer Colorizer) {
	minY := bounds.Min.Y
	maxY := bounds.Max.Y
	minX := bounds.Min.X
	maxX := bounds.Max.X
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			c := proj.Project(x, y)
			res := mandelbrot.Mandelbrot(c, maxIter)
			col := colorizer.Colorize(res)
			im.Set(x, y, col)
		}
	}
}
