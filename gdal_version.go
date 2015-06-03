// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include "cgo_gdal.h"
import "C"

const (
	MajorVersion = int(C.GDAL_VERSION_MAJOR)
	MinorVersion = int(C.GDAL_VERSION_MINOR)
	RevVersion   = int(C.GDAL_VERSION_REV)
	BuildVersion = int(C.GDAL_VERSION_BUILD)

	ReleaseDate = int(C.GDAL_RELEASE_DATE)
	ReleaseName = string(C.GDAL_RELEASE_NAME)
)
