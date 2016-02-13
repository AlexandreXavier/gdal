// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"testing"
)

func TestLoad_gdal(t *testing.T) {
	if _, err := LoadImage("./testdata/video-001.tiff"); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkLoad_empty_8000x6000_png_gdal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := LoadImage("./testdata/empty8000x6000.png"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLoad_video_001_tiff_gdal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := LoadImage("./testdata/video-001.tiff"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLoad_video_001_png(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := LoadImage("./testdata/video-001.png"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLoad_video_001_jpeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := LoadImage("./testdata/video-001.jpeg"); err != nil {
			b.Fatal(err)
		}
	}
}
