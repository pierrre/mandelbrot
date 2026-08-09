// Package cmd contains common utilities for commands.
package cmd

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
)

// Save saves an image to a file.
func Save(im image.Image, file string) error {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, im)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	err = os.WriteFile(file, buf.Bytes(), os.FileMode(0o644))
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}
