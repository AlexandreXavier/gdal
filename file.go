// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

type File struct {
	private int
}

func Create(name string) (f *File, err error) {
	return
}

func Open(name string) (f *File, err error) {
	return
}

func OpenFile(name string, flag int) (f *File, err error) {
	return
}

func (f *File) Close() error {
	return nil
}
