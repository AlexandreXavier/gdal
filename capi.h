// Copyright 2015 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef GO_GDAL_H
#define GO_GDAL_H

#include <gdal.h>
#include <cpl_conv.h>

#ifdef __cplusplus
extern "C" {
#endif

// transform GDALProgressFunc to go func
GDALProgressFunc goGDALProgressFuncProxyB();

#ifdef __cplusplus
}
#endif
#endif	// GO_GDAL_H

