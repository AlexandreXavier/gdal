// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef GO_GDAL_H_
#define GO_GDAL_H_

#include <gdal.h>
#include <cpl_conv.h>

// transform GDALProgressFunc to go func
GDALProgressFunc goGDALProgressFuncProxyB();

#endif // GO_GDAL_H_


