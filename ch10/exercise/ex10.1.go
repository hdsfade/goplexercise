//@author: hdsfade
//@date: 2020-11-11-19:00
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

var usage = "usage:please input format [png],[jpg],[gif]]"
var f = flag.String("f", "png", "usage")

func main() {
	flag.Parse()
	format := strings.ToLower(*f)
	err := toformat(os.Stdout, os.Stdin, format)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ex10.1: %s\n", err)
		os.Exit(1)
	}
}

func toformat(w io.Writer, r io.Reader, format string) error {
	img, kind, err := image.Decode(r)
	if err != nil {
		return nil
	}
	fmt.Fprintf(os.Stdout, "input format: %s\n", kind)
	err = Encode(w, img, format)
	return err
}

func Encode(out io.Writer, img image.Image, format string) (err error) {
	switch format {
	case "png":
		err = png.Encode(out, img)
	case "jpg", "jpeg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		err = gif.Encode(out, img, nil)
	}
	return err
}
