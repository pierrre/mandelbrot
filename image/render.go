package image

import (
	"image"
	"image/draw"
	"runtime"
	"sync"

	"github.com/pierrre/mandelbrot"
)

type Renderer interface {
	Render(draw.Image, Transformation, mandelbrot.Mandelbroter, Colorizer)
}

type RendererFunc func(draw.Image, Transformation, mandelbrot.Mandelbroter, Colorizer)

func (f RendererFunc) Render(im draw.Image, transf Transformation, mandel mandelbrot.Mandelbroter, colzr Colorizer) {
	f(im, transf, mandel, colzr)
}

func NewRenderer() Renderer {
	return RendererFunc(func(im draw.Image, transf Transformation, mandel mandelbrot.Mandelbroter, colzr Colorizer) {
		render(im, im.Bounds(), transf, mandel, colzr)
	})
}

func NewRendererWorker(workerCount int) Renderer {
	return RendererFunc(func(im draw.Image, transf Transformation, mandel mandelbrot.Mandelbroter, colzr Colorizer) {
		renderWorker(im, transf, mandel, colzr, workerCount)
	})
}

func NewRendererWorkerAuto() Renderer {
	return NewRendererWorker(runtime.GOMAXPROCS(0) * 4)
}

func renderWorker(im draw.Image, transf Transformation, mandel mandelbrot.Mandelbroter, colzr Colorizer, workerCount int) {
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
			render(im, wBounds, transf, mandel, colzr)
			wg.Done()
		}(wBounds)
	}

	wg.Wait()
}

func render(im draw.Image, bounds image.Rectangle, transf Transformation, mandel mandelbrot.Mandelbroter, colzr Colorizer) {
	minY := bounds.Min.Y
	maxY := bounds.Max.Y
	minX := bounds.Min.X
	maxX := bounds.Max.X
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			c := complex(float64(x), float64(y))
			c = transf.Transform(c)
			res := mandel.Mandelbrot(c)
			col := colzr.Colorize(c, res)
			im.Set(x, y, col)
		}
	}
}
