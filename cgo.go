// Copyright 2015 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

/*
#cgo windows CFLAGS: -I./build-windows/include
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/build-windows/lib -l"gdal-cgo-amd64"
#cgo windows,386 LDFLAGS: -L${SRCDIR}/build-windows/lib -l"gdal-cgo-386"

#cgo linux CFLAGS: -I/usr/include/gdal
#cgo linux LDFLAGS: -L/usr/lib -lgdal

#cgo darwin pkg-config: gdal

#include <gdal.h>
*/
import "C"
