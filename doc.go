// Copyright 2015 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package gdal provides Go bindings for GDAL.

Install `GCC` or `MinGW` (http://tdm-gcc.tdragon.net/download) at first,
and then run these commands:

	1. `go get github.com/chai2010/gdal`
	2. `go run hello.go`

Example:

	// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
	// Use of this source code is governed by a BSD-style
	// license that can be found in the LICENSE file.

	package main

	import (
		"fmt"
		"log"

		"github.com/chai2010/gdal"
	)

	func main() {
		fmt.Printf("GDAL %d.%d.%d\n", gdal.MajorVersion, gdal.MinorVersion, gdal.RevVersion)

		// load data
		m, err := gdal.Load("./testdata/lena512color.png")
		if err != nil {
			log.Fatal("gdal.Load:", err)
		}

		// save bmp
		err = gdal.Save("output.bmp", m, nil)
		if err != nil {
			log.Fatal("gdal.Save:", err)
		}

		// save tiff
		err = gdal.Save("output.tiff", m, nil)
		if err != nil {
			log.Fatal("gdal.Save:", err)
		}

		// save jpeg-tiff data
		err = gdal.Save("output.jpeg.tiff", m, &gdal.Options{
			DriverName: "GTiff",
			ExtOptions: map[string]string{
				"COMPRESS":     "JPEG",
				"JPEG_QUALITY": "75",
			},
		})
		if err != nil {
			log.Fatal("gdal.Save:", err)
		}

		fmt.Println("Done.")
	}

Report bugs to <chaishushan@gmail.com>.

Thanks!
*/
package gdal // import "github.com/chai2010/gdal"
