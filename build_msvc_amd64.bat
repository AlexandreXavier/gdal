:: Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

@echo off

cd %~dp0
setlocal

:: ----------------------------------------------------------------------------
:::: Setup MSVC
::
:::: VS2010
::if not "x%VS100COMNTOOLS%" == "x" (
::	echo Setup VS2010 Win64 ...
::	call "%VS100COMNTOOLS%\..\..\VC\vcvarsall.bat" x86_amd64
::	goto build
::)
::
:::: VS2012
::if not "x%VS110COMNTOOLS%" == "x" (
::	echo Setup VS2012 Win64 ...
::	call "%VS110COMNTOOLS%\..\..\VC\vcvarsall.bat" x86_amd64
::	goto build
::)
::
:::: VS2013
::if not "x%VS120COMNTOOLS%" == "x" (
::	echo Setup VS2013 Win64 ...
::	call "%VS120COMNTOOLS%\..\..\VC\vcvarsall.bat" x86_amd64
::	goto build
::)
::
:::build

:: ----------------------------------------------------------------------------

::mkdir zz_msvc_amd64
::cd zz_msvc_amd64
::
::cmake ..^
::  -G "NMake Makefiles"^
::  -DCMAKE_BUILD_TYPE=release^
::  -DCMAKE_INSTALL_PREFIX=..^
::  ^
::  -DCMAKE_C_FLAGS_DEBUG="/MTd /Zi /Od /Ob0 /RTC1"^
::  -DCMAKE_CXX_FLAGS_DEBUG="/MTd /Zi /Od /Ob0 /RTC1"^
::  ^
::  -DCMAKE_C_FLAGS_RELEASE="/MT /O2 /Ob2 /DNDEBUG"^
::  -DCMAKE_CXX_FLAGS_RELEASE="/MT /O2 /Ob2 /DNDEBUG"^
::  ^
::  -DCMAKE_MODULE_LINKER_FLAGS="/MANIFEST:NO /machine:x64"^
::  -DCMAKE_SHARED_LINKER_FLAGS="/MANIFEST:NO /machine:x64"^
::  -DCMAKE_STATIC_LINKER_FLAGS="/MANIFEST:NO /machine:x64"^
::  -DCMAKE_EXE_LINKER_FLAGS="/MANIFEST:NO /MACHINE:x64"
::
::nmake install
::cd ..

copy .\internal\build-windows_amd64\gdal111.dll gdal-cgo-win64.dll
dlltool -dllname gdal-cgo-win64.dll --def gdal-cgo-win64.def --output-lib libgdal-cgo-amd64.a

:: ----------------------------------------------------------------------------
:: PAUSE
