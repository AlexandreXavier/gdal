// Copyright 2015 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if a, b := MajorVersion, _GDAL_VERSION_MAJOR; a != b {
		t.Fatalf("expect = %v, got = %v", b, a)
	}
	if a, b := MinorVersion, _GDAL_VERSION_MINOR; a != b {
		t.Fatalf("expect = %v, got = %v", b, a)
	}
	if a, b := RevVersion, _GDAL_VERSION_REV; a != b {
		t.Fatalf("expect = %v, got = %v", b, a)
	}
	if a, b := BuildVersion, _GDAL_VERSION_BUILD; a != b {
		t.Fatalf("expect = %v, got = %v", b, a)
	}

	if a, b := ReleaseDate, _GDAL_RELEASE_DATE; a != b {
		t.Fatalf("expect = %v, got = %v", b, a)
	}
	if a, b := ReleaseName, _GDAL_RELEASE_NAME; a != b {
		t.Fatalf("expect = %v, got = %v", b, a)
	}
}
