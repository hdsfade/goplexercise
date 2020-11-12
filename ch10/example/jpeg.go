//@author: hdsfade
//@date: 2020-11-11-18:49
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
