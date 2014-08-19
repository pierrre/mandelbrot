package image

import (
	"image"
	"image/draw"
	"runtime"
	"sync"

	"github.com/pierrre/mandelbrot"
)

func Render(m mandelbrot.Mandelbroter, im draw.Image, trans Transformation, colorizer Colorizer) {
	render(m, im, im.Bounds(), trans, colorizer)
}

func RenderWorker(m mandelbrot.Mandelbroter, im draw.Image, trans Transformation, colorizer Colorizer, workerCount int) {
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
			render(m, im, wBounds, trans, colorizer)
			wg.Done()
		}(wBounds)
	}

	wg.Wait()
}

func RenderWorkerAuto(m mandelbrot.Mandelbroter, im draw.Image, trans Transformation, colorizer Colorizer) {
	RenderWorker(m, im, trans, colorizer, runtime.GOMAXPROCS(0)*4)
}

func render(m mandelbrot.Mandelbroter, im draw.Image, bounds image.Rectangle, trans Transformation, colorizer Colorizer) {
	minY := bounds.Min.Y
	maxY := bounds.Max.Y
	minX := bounds.Min.X
	maxX := bounds.Max.X
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			c := complex(float64(x), float64(y))
			c = trans.Transform(c)
			res := m.Mandelbrot(c)
			col := colorizer.Colorize(c, res)
			im.Set(x, y, col)
		}
	}
}
