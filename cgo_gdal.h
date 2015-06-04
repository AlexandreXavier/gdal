// Copyright 2015 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef CGO_GDAL_H
#define CGO_GDAL_H

#include <stdint.h>

#include <gdal.h>
#include <gdal_version.h>
#include <cpl_conv.h>

#ifdef __cplusplus
extern "C" {
#endif

// transform GDALProgressFunc to go func
GDALProgressFunc cgoGDALProgressFuncProxyB();

#ifdef __cplusplus
}
#endif
#endif	// CGO_GDAL_H

