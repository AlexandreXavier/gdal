:: Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

:: VS2010
if not "x%VS100COMNTOOLS%" == "x" (
	call "%VS100COMNTOOLS%\..\..\VC\vcvarsall.bat"
)

:: VS2012
if not "x%VS110COMNTOOLS%" == "x" (
	call "%VS110COMNTOOLS%\..\..\VC\vcvarsall.bat"
)

:: VS2013
if not "x%VS120COMNTOOLS%" == "x" (
	call "%VS120COMNTOOLS%\..\..\VC\vcvarsall.bat"
)

