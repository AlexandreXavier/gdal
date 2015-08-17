// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <gdal.h>
import "C"

func init() {
	C.GDALAllRegister()
	C.VSIInstallMemFileHandler()
}
