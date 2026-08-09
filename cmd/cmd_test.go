package cmd

import (
	"image"
	"os"
	"path/filepath"
	"testing"
)

func TestSave(t *testing.T) {
	im := image.NewGray(image.Rect(0, 0, 16, 16))
	file := filepath.Join(t.TempDir(), "test.png")
	err := Save(im, file)
	if err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(file)
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() == 0 {
		t.Fatal("file is empty")
	}
}

func TestSaveEncodeError(t *testing.T) {
	im := image.NewGray(image.Rect(0, 0, 0, 0))
	file := filepath.Join(t.TempDir(), "test.png")
	err := Save(im, file)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSaveWriteError(t *testing.T) {
	im := image.NewGray(image.Rect(0, 0, 16, 16))
	file := filepath.Join(t.TempDir(), "nonexistent", "test.png")
	err := Save(im, file)
	if err == nil {
		t.Fatal("expected error")
	}
}
