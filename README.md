## Install

Linux

1. install Go1 and GDAL
2. go get github.com/chai2010/gdal
3. go run hello.go

Windows

1. install Go1 and MinGW
2. build gdal-1.9.1 with MinGW
   cd gdal-1.9.1 dir
   ./configure
   make
   make install
3. copy gdal *.h/*.a/*.dll to MinGW dir 
   MinGW\msys\1.0\local\bin\*.dll -> ${MinGWRoot}\*.dll
   MinGW\msys\1.0\local\include   -> ${MinGWRoot}\include
   MinGW\msys\1.0\local\lib\*lib  -> ${MinGWRoot}\lib\*lib
4. go get github.com/chai2010/gdal
5. go run hello.go
