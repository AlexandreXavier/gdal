// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package gdal

/*
#include <gdal_version.h>
*/
import "C"

const (
	GDAL_VERSION_MAJOR = int(C.GDAL_VERSION_MAJOR)
	GDAL_VERSION_MINOR = int(C.GDAL_VERSION_MINOR)
	GDAL_VERSION_REV   = int(C.GDAL_VERSION_REV)
	GDAL_VERSION_BUILD = int(C.GDAL_VERSION_BUILD)

	GDAL_VERSION_NUM  = int(C.GDAL_VERSION_NUM)
	GDAL_RELEASE_DATE = int(C.GDAL_RELEASE_DATE)

	GDAL_RELEASE_NAME = string(C.GDAL_RELEASE_NAME)
)
