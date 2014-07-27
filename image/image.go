package image

import (
	"image"
	"image/color"
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
			ok, iter, abs := mandelbrot.Mandelbrot(c, maxIter)
			col := colorizer.Colorize(ok, iter, abs)
			im.Set(x, y, col)
		}
	}
}

type Projection interface {
	Project(x, y int) complex128
}

type ProjectionFunc func(x, y int) complex128

func (pf ProjectionFunc) Project(x, y int) complex128 {
	return pf(x, y)
}

type Colorizer interface {
	Colorize(ok bool, iter int, abs float64) color.Color
}

type ColorizerFunc func(ok bool, iter int, abs float64) color.Color

func (f ColorizerFunc) Colorize(ok bool, iter int, abs float64) color.Color {
	return f(ok, iter, abs)
}

var BWColorizer = ColorizerFunc(func(ok bool, iter int, abs float64) color.Color {
	if !ok {
		return color.White
	} else {
		return color.Black
	}
})
