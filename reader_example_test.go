// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"fmt"
	"log"
)

func ExampleLoadConfig() {
	cfg, err := LoadConfig("./testdata/video-001.tiff")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Width = %d\n", cfg.Width)
	fmt.Printf("Height = %d\n", cfg.Height)
	// Output:
	// Width = 150
	// Height = 103
}

func ExampleLoad() {
	m, err := Load("./testdata/video-001.tiff")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bounds = %v\n", m.Bounds())
	// Output:
	// Bounds = (0,0)-(150,103)
}
