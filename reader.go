// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"fmt"
	"image"
)

// Load reads a GDAL image from file and returns it as an image.Image.
func Load(filename string) (m image.Image, err error) {
	err = fmt.Errorf("gdal: Load, TODO")
	return
}
