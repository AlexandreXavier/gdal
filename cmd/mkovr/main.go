// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
// Make GDAL Overviews file is not exists.
//
//	Usage: mkovr filename [ResampleType]
//	       mkovr -h
//
//	Example:
//	  mkovr filename
//	  mkovr filename GAUSS
//	  mkovr filename AVERAGE
//
//	ResampleType: NONE|NEAREST|GAUSS|CUBIC|AVERAGE|MODE|AVERAGE_MAGPHASE.
//
//	Report bugs to <chaishushan{AT}gmail.com>.
//
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chai2010/gdal"
)

const usage = `
Usage: mkovr filename [ResampleType]
       mkovr -h

Example:
  mkovr filename
  mkovr filename GAUSS
  mkovr filename AVERAGE

ResampleType: NONE|NEAREST|GAUSS|CUBIC|AVERAGE|MODE|AVERAGE_MAGPHASE.

Report bugs to <chaishushan{AT}gmail.com>.
`

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}

	filename, resampleTypeName := os.Args[1], "NONE"
	if len(os.Args) > 2 {
		resampleTypeName = os.Args[2]
	}

	resampleType := gdal.NewResampleType(resampleTypeName)
	po, err := gdal.OpenDatasetWithOverviews(filename, resampleType, gdal.GA_ReadOnly)
	if err != nil {
		log.Fatal(filename, resampleType, err)
	}
	defer po.Close()
}
