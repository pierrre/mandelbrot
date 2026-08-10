package cmd

import (
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/pierrre/assert"
)

func TestSave(t *testing.T) {
	im := image.NewGray(image.Rect(0, 0, 16, 16))
	file := filepath.Join(t.TempDir(), "test.png")
	assert.NoError(t, Save(im, file))
	info, err := os.Stat(file)
	assert.NoError(t, err)
	assert.NotZero(t, info.Size())
}

func TestSaveEncodeError(t *testing.T) {
	im := image.NewGray(image.Rect(0, 0, 0, 0))
	file := filepath.Join(t.TempDir(), "test.png")
	assert.Error(t, Save(im, file))
}

func TestSaveWriteError(t *testing.T) {
	im := image.NewGray(image.Rect(0, 0, 16, 16))
	file := filepath.Join(t.TempDir(), "nonexistent", "test.png")
	assert.Error(t, Save(im, file))
}
