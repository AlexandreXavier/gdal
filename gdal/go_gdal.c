// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_gdal.h"
#include "_cgo_export.h"

#include <cpl_conv.h>

static
int goGDALProgressFuncProxyB_(
	double dfComplete, const char *pszMessage, void *pProgressArg
) {
	GoInterface* args = (GoInterface*)pProgressArg;
	GoInt rv = goGDALProgressFuncProxyA(dfComplete, (char*)pszMessage, args);
	return (int)rv;
}

GDALProgressFunc goGDALProgressFuncProxyB() {
	return goGDALProgressFuncProxyB_;
}


