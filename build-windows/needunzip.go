// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

// return 0: do not need unzip
// return 1: need unzip
import (
	"flag"
	"io/ioutil"
	"os"
)

var flagName = flag.String("name", "", "file name")

var golden = func() string {
	data, _ := ioutil.ReadFile("nmake-mt.opt")
	return string(data)
}()

func main() {
	flag.Parse()

	if *flagName != "" {
		if needReplce(*flagName) {
			os.Exit(1)
		}
	} else {
		if needReplce(`gdal192-win32\gdal-1.9.2\nmake.opt`) {
			os.Exit(1)
		}
		if needReplce(`gdal192-win64\gdal-1.9.2\nmake.opt`) {
			os.Exit(1)
		}
	}
}

func needReplce(name string) bool {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return true
	}
	if string(data) != golden {
		return true
	}
	return false
}
