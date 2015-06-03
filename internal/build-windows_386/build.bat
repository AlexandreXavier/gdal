:: Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

go run cpdir.go -recurse=false ../gdal-1.11.2/port            include "\.h$"
go run cpdir.go -recurse=false ../gdal-1.11.2/gcore           include "\.h$"
go run cpdir.go -recurse=false ../gdal-1.11.2/alg             include "\.h$"
go run cpdir.go -recurse=false ../gdal-1.11.2/ogr             include "\.h$"
go run cpdir.go -recurse=false ../gdal-1.11.2/ogr/ogrsf_frmts include "\.h$"
go run cpdir.go -recurse=false ../gdal-1.11.2/frmts/mem       include "\.h$"
go run cpdir.go -recurse=false ../gdal-1.11.2/frmts/raw       include "\.h$"

