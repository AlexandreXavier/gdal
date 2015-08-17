// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoad_unclosed(t *testing.T) {
	data := tbLoadData(t, "video-001.tiff")

	tmpfilename := "zz_video-001.tiff"
	defer os.Remove(tmpfilename)

	f, err := os.Create(tmpfilename)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := LoadConfig(tmpfilename)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Width != 150 || cfg.Height != 103 {
		t.Fatalf("cfg: %v", cfg)
	}

	_, err = Load(tmpfilename)
	if err != nil {
		t.Fatal(err)
	}
}

func tbLoadData(tb testing.TB, filename string) []byte {
	data, err := ioutil.ReadFile("./testdata/" + filename)
	if err != nil {
		tb.Fatal(err)
	}
	return data
}
