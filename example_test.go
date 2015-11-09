// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal_test

import (
	"fmt"
	"log"

	"github.com/chai2010/gdal"
)

func ExampleDataset_setTFW() {
	po, err := gdal.OpenDataset("some.tiff", gdal.GA_Update)
	if err != nil {
		log.Fatal(err)
	}
	defer po.Close()

	if err = po.SetGeoTransformX0Y0DxDy(0, 0, 1, 1); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
