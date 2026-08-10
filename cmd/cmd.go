// Package cmd contains common utilities for commands.
package cmd

import (
	"bytes"
	"image"
	"image/png"
	"os"

	"github.com/pierrre/errors"
)

// Save saves an image to a file.
func Save(im image.Image, file string) error {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, im)
	if err != nil {
		return errors.Wrap(err, "encode")
	}
	err = os.WriteFile(file, buf.Bytes(), os.FileMode(0o644))
	if err != nil {
		return errors.Wrap(err, "write")
	}
	return nil
}
