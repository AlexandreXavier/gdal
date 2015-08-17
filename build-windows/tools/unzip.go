// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

//
// unzip file.
//
// Example:
//	unzip zipfile
//	unzip src gdal1112.zip .
//
// Help:
//	unzip -h
//
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const usage = `
Usage: unzip zipfile [outdir]
       unzip -h

Example:
  unzip src gdal1112.zip
  unzip src gdal1112.zip .
    
Report bugs to <chaishushan{AT}gmail.com>.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}

	var filename = os.Args[1]
	var basename = filepath.Base(filename)

	var destdir = "./out"
	if strings.HasSuffix(basename, ".zip") {
		destdir = basename[:len(basename)-len(".zip")]
	}
	if len(os.Args) > 2 {
		destdir = os.Args[2]
	}

	err := Unzip(filename, destdir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done.\n")
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fmt.Printf("unzip %s\n", f.Name)
		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
