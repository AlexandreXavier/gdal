:: Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

@echo off

cd %~dp0
setlocal

:: ----------------------------------------------------------------------------
:: generate libgdal-cgo-xxx.a for dll

copy .\internal\build-windows_amd64\gdal111.dll gdal-cgo-win64.dll
dlltool -dllname gdal-cgo-win64.dll --def gdal-cgo-win64.def --output-lib libgdal-cgo-amd64.a

:: ----------------------------------------------------------------------------
:: PAUSE
