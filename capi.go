// Copyright 2015 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

/*
#cgo windows CFLAGS: -I./. -fno-stack-check -fno-stack-protector -mno-stack-arg-probe
#cgo windows,amd64 LDFLAGS: -L${SRCDIR} -l"gdal-cgo-amd64"
#cgo windows,386 LDFLAGS: -L${SRCDIR} -l"gdal-cgo-386"

#cgo linux pkg-config: gdal

#include "capi.h"
*/
import "C"
