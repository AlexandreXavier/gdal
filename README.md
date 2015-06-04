# Go bindings for GDAL

## Install

1. go get github.com/chai2010/gdal
2. go run hello.go

## Example

```Go
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

	// save data
	err = gdal.Save("output.tiff", m, nil)
	if err != nil {
		log.Fatal("gdal.Save:", err)
	}

	fmt.Println("Done.")
}
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
