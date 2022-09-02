// Package cmd contains common utils for commands.
package cmd

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

// Save saves an image to a file.
func Save(im image.Image, file string) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, im)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(file, buf.Bytes(), os.FileMode(0o644))
	if err != nil {
		panic(err)
	}
}
