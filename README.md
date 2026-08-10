# Mandelbrot

Mandelbrot set library for Go (Golang).

[![Go Reference](https://pkg.go.dev/badge/github.com/pierrre/mandelbrot.svg)](https://pkg.go.dev/github.com/pierrre/mandelbrot)

## Features

- [Compute](https://pkg.go.dev/github.com/pierrre/mandelbrot) the Mandelbrot set for a point
- [Render](https://pkg.go.dev/github.com/pierrre/mandelbrot/image) to an image (parallel)
- [Transformations](https://pkg.go.dev/github.com/pierrre/mandelbrot/image#Transformation): scale, rotate, translate
- [Colorizers](https://pkg.go.dev/github.com/pierrre/mandelbrot/image#Colorizer): black & white, custom colors
- [Rainbow colorizer](https://pkg.go.dev/github.com/pierrre/mandelbrot/image/colorizer/rainbow)
- Powers from 2 to 20 (and arbitrary powers)
- Example programs: simple render, colored render, exploration, HTTP server

## Usage

```bash
# Module install
go get github.com/pierrre/mandelbrot@latest

# Local build
make build
./build/<program>
```

### Programs

- `simple` - renders a simple black & white image to `simple.png`
- `color` - renders a colored image to `color.png`
- `explore` - explores the set by zooming in, outputting `explore_XXXX.png` files
- `httpserver` - HTTP server that serves Mandelbrot tiles (OpenLayers)

## Example

```go
package main

import (
	"image"

	"github.com/pierrre/mandelbrot"
	mandelbrot_cmd "github.com/pierrre/mandelbrot/cmd"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
)

func main() {
	size := image.Pt(1024, 1024)
	rotate := 0.0
	scale := 1.6
	translate := complex(-0.75, 0)

	im := image.NewGray(image.Rect(0, 0, size.X, size.Y))

	scale *= mandelbrot_image.Scale(size)
	tsf := mandelbrot_image.BaseTransformation(im, rotate, scale, translate)
	maxIter := mandelbrot_image.MaxIter(scale)
	f := mandelbrot.New(maxIter)
	clr := mandelbrot_image.BWColorizer(false)
	mandelbrot_image.RenderParallel(im, tsf, f, clr)

	mandelbrot_cmd.Save(im, "simple.png")
}
```
