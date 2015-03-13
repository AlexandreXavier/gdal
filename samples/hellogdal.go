// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"code.google.com/p/go-gdal/trunk/gdal"
)

func main() {
	fmt.Printf("GO-GDAL Info\n")
	fmt.Printf("Version Num: %d\n", gdal.GDAL_VERSION_NUM)
	fmt.Printf("Release Data: %d\n", gdal.GDAL_RELEASE_DATE)
	fmt.Printf("Release Name: %s\n", gdal.GDAL_RELEASE_NAME)
}


