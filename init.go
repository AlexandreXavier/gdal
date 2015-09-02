// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

/*
#include <gdal.h>
#include <cpl_conv.h>

void initGDAL() {
	GDALAllRegister();
	CPLSetConfigOption("GDAL_TIFF_OVR_BLOCKSIZE", "256");
}
*/
import "C"

func init() {
	C.initGDAL()
}
