// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include "cgo_gdal.h"
import "C"

const (
	MajorVersion = 1  // GDAL_VERSION_MAJOR
	MinorVersion = 11 // GDAL_VERSION_MINOR
	RevVersion   = 2  // GDAL_VERSION_REV
	BuildVersion = 0  // GDAL_VERSION_BUILD

	ReleaseDate = 20150210 // GDAL_RELEASE_DATE
	ReleaseName = "1.11.2" // GDAL_RELEASE_NAME
)

const (
	_GDAL_VERSION_MAJOR = int(C.GDAL_VERSION_MAJOR)
	_GDAL_VERSION_MINOR = int(C.GDAL_VERSION_MINOR)
	_GDAL_VERSION_REV   = int(C.GDAL_VERSION_REV)
	_GDAL_VERSION_BUILD = int(C.GDAL_VERSION_BUILD)

	_GDAL_VERSION_NUM  = int(C.GDAL_VERSION_NUM)
	_GDAL_RELEASE_DATE = int(C.GDAL_RELEASE_DATE)

	_GDAL_RELEASE_NAME = string(C.GDAL_RELEASE_NAME)
)
