:: Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

@echo off

cd %~dp0
setlocal

:: ----------------------------------------------------------------------------
:: generate libgdal-cgo-xxx.a for dll

dlltool -dllname leveldb-cgo-win32.dll --def leveldb-cgo-win32.def --output-lib libleveldb-cgo-win32.a
copy leveldb-cgo-win32.dll ..\..\..\bin

:: ----------------------------------------------------------------------------
:: PAUSE
