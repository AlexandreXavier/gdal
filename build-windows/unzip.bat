:: Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.
@echo off

setlocal

go run needunzip.go gdal192-win32\gdal-1.9.2\nmake.opt
if not %errorlevel% == 0 (
	.\7za.exe x -y gdal192.7z -ogdal192-win32 >7z-stdout.txt
	copy /Y nmake-mt.opt gdal192-win32\gdal-1.9.2\nmake.opt
)

go run needunzip.go gdal192-win64\gdal-1.9.2\nmake.opt
if not %errorlevel% == 0 (
	.\7za.exe  x -y gdal192.7z -ogdal192-win64 >7z-stdout.txt
	copy /Y nmake-mt.opt gdal192-win64\gdal-1.9.2\nmake.opt
)

