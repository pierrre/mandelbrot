package _examples

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"os"
)

func Save(im image.Image, file string) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, im)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(file, buf.Bytes(), os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}
