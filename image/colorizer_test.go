package image_test

import (
	"image/color"
	"testing"

	"github.com/pierrre/mandelbrot"
	"github.com/pierrre/mandelbrot/image"
)

func TestColorsIterColorizer(t *testing.T) {
	cols := []color.Color{
		color.RGBA{R: 0, A: 255},
		color.RGBA{R: 1, A: 255},
		color.RGBA{R: 2, A: 255},
	}
	tests := []struct {
		name  string
		shift int
		iter  int
		want  color.Color
	}{
		{
			name:  "zero shift iter 0",
			shift: 0,
			iter:  0,
			want:  cols[0],
		},
		{
			name:  "zero shift iter 1",
			shift: 0,
			iter:  1,
			want:  cols[1],
		},
		{
			name:  "zero shift wrap",
			shift: 0,
			iter:  3,
			want:  cols[0],
		},
		{
			name:  "positive shift",
			shift: 1,
			iter:  0,
			want:  cols[1],
		},
		{
			name:  "negative shift iter 0",
			shift: -1,
			iter:  0,
			want:  cols[2],
		},
		{
			name:  "negative shift iter 1",
			shift: -1,
			iter:  1,
			want:  cols[0],
		},
		{
			name:  "negative shift no remainder",
			shift: -3,
			iter:  0,
			want:  cols[0],
		},
		{
			name:  "large negative shift",
			shift: -10,
			iter:  0,
			want:  cols[2],
		},
		{
			name:  "shift larger than len",
			shift: 5,
			iter:  0,
			want:  cols[2],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clr := image.ColorsIterColorizer(cols, tt.shift)
			res := mandelbrot.Result{Iter: tt.iter}
			got := clr(0, res)
			if !colorsEqual(got, tt.want) {
				r1, g1, b1, a1 := got.RGBA()
				r2, g2, b2, a2 := tt.want.RGBA()
				t.Fatalf("got RGBA(%d,%d,%d,%d), want RGBA(%d,%d,%d,%d)", r1, g1, b1, a1, r2, g2, b2, a2)
			}
		})
	}
}

func TestColorsIterColorizerEmpty(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("expected panic")
			}
		}()
		image.ColorsIterColorizer(nil, 0)
	})
	t.Run("empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("expected panic")
			}
		}()
		image.ColorsIterColorizer([]color.Color{}, 0)
	})
}

func colorsEqual(c1, c2 color.Color) bool {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}
