// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"fmt"
	"image"
)

// Encode writes the image m to w in GDAL format.
func Save(filename string, m image.Image, opt *Options) error {
	return fmt.Errorf("gdal: Save, TODO")
}
