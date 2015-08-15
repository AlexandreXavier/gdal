// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"image"
	"log"
	"os"
	"reflect"
)

func ExampleSave() {
	tmpname := "z_test_ExampleSave.tiff"
	defer os.Remove(tmpname)

	gray := NewMemPImage(image.Rect(0, 0, 400, 300), 1, reflect.Uint8)
	if err := Save(tmpname, gray, nil); err != nil {
		log.Fatal(err)
	}
}

func ExampleSave_cbuf() {
	tmpname := "z_test_ExampleSave.tiff"
	defer os.Remove(tmpname)

	b := image.Rect(0, 0, 400, 300)
	cbuf := NewCBuffer(b.Dx() * b.Dy() * 4)
	defer cbuf.Close()

	rgba := &image.RGBA{
		Pix:    cbuf.CData(),
		Stride: b.Dx() * 4,
		Rect:   b,
	}
	if err := Save(tmpname, rgba, nil, cbuf); err != nil {
		log.Fatal(err)
	}

	gray := &image.Gray{
		Pix:    cbuf.CData(),
		Stride: b.Dx() * 1,
		Rect:   b,
	}
	if err := Save(tmpname, gray, nil, cbuf); err != nil {
		log.Fatal(err)
	}
}
