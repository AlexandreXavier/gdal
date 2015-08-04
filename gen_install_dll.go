// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Copy gdal-cgo-win64.dll to $(GOPATH)/bin
package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	src := "./gdal-cgo-win64.dll"
	dst := filepath.Join(gopathRoot(), "bin/gdal-cgo-win64.dll")
	cpFile(dst, src)
}

func gopathRoot() string {
	listOut, err := exec.Command(`go`, `list`, `-f`, `{{.Root}}`).Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(listOut))
}

func cpFile(dst, src string) {
	err := os.MkdirAll(filepath.Dir(dst), 0666)
	if err != nil && !os.IsExist(err) {
		log.Fatal("cpFile: ", err)
	}
	fsrc, err := os.Open(src)
	if err != nil {
		log.Fatal("cpFile: ", err)
	}
	defer fsrc.Close()

	fdst, err := os.Create(dst)
	if err != nil {
		log.Fatal("cpFile: ", err)
	}
	defer fdst.Close()
	if _, err = io.Copy(fdst, fsrc); err != nil {
		log.Fatal("cpFile: ", err)
	}
}
