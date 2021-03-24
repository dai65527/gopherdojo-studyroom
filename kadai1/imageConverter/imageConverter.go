package imageConverter

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func Jpg2png(filename string) {
	fileext := strings.ToLower(filepath.Ext(filename))
	if fileext != ".jpg" && fileext != ".jpeg" {
		return
	}

	fs, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	img, err := jpeg.Decode(fs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fs.Close()

	fs, err = os.OpenFile(filename+".png", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = png.Encode(fs, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fs.Close()
}

func Png2jpg(filename string) {
	fileext := strings.ToLower(filepath.Ext(filename))
	if fileext != ".png" {
		return
	}

	fs, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	img, err := png.Decode(fs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fs.Close()

	fs, err = os.OpenFile(filename+".jpg", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = jpeg.Encode(fs, img, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fs.Close()
}
