// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"testing"
)

func TestDataType(t *testing.T) {
	var tests = []struct{ Expect, Got DataType }{
		{GDT_Unknown, _GDT_Unknown},
	}
	for i, tt := range tests {
		if tt.Expect != tt.Got {
			t.Fatalf("%d: expect = %v, got = %v", i, tt.Expect, tt.Got)
		}
	}
}
