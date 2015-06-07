:: Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.
@echo off

setlocal

cd %~dp0

call unzip.bat

setlocal
	echo build gdal192 win32 release ...
	call .\register-vc-386.bat
	cd gdal192-win32\gdal-1.9.2
	nmake /f makefile.vc 1>build-log-1.txt 2>build-log-2.txt
endlocal

setlocal
	echo build gdal192 win64 release ...
	call .\register-vc-x64.bat
	cd gdal192-win64\gdal-1.9.2
	nmake /f makefile.vc 1>build-log-1.txt 2>build-log-2.txt
endlocal


