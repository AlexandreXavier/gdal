// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"

	"github.com/chai2010/gdal"
)

func main() {
	fmt.Printf("GDAL %d.%d.%d\n", gdal.MajorVersion, gdal.MinorVersion, gdal.RevVersion)
}
