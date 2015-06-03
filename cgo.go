// Copyright 2015 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

/*
#cgo windows,amd64 CFLAGS: -I./internal/build-windows_amd64/include
#cgo windows,386 CFLAGS: -I./internal/build-windows_386/include

#cgo windows,amd64 LDFLAGS: -L${SRCDIR} -l"gdal-cgo-amd64"
#cgo windows,386 LDFLAGS: -L${SRCDIR} -l"gdal-cgo-386"

#cgo linux pkg-config: gdal

#include "cgo_gdal.h"
*/
import "C"
