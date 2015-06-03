// Copyright 2015 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "cgo_gdal.h"
#include "_cgo_export.h"

#include <cpl_conv.h>

static
int cgoGDALProgressFuncProxyB_(
	double dfComplete, const char *pszMessage, void *pProgressArg
) {
	GoInterface* args = (GoInterface*)pProgressArg;
	GoInt rv = goGDALProgressFuncProxyA(dfComplete, (char*)pszMessage, args);
	return (int)rv;
}

GDALProgressFunc cgoGDALProgressFuncProxyB() {
	return cgoGDALProgressFuncProxyB_;
}
