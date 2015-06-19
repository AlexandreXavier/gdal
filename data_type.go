// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <gdal.h>
import "C"
import (
	"reflect"
)

func gdalDataType(dataType reflect.Kind) C.GDALDataType {
	switch dataType {
	case reflect.Int8:
		// invalid
	case reflect.Int16:
		return C.GDT_Int16
	case reflect.Int32:
		return C.GDT_Int32
	case reflect.Int64:
		// invalid
	case reflect.Uint8:
		return C.GDT_Byte
	case reflect.Uint16:
		return C.GDT_UInt16
	case reflect.Uint32:
		return C.GDT_UInt32
	case reflect.Uint64:
		// invalid
	case reflect.Float32:
		return C.GDT_Float32
	case reflect.Float64:
		return C.GDT_Float64
	case reflect.Complex64:
		// invalid
	case reflect.Complex128:
		// invalid
	}
	return C.GDT_Unknown
}

func goDataType(dataType C.GDALDataType) reflect.Kind {
	switch dataType {
	case C.GDT_Unknown:
		// invalid
	case C.GDT_Byte:
		return reflect.Uint8
	case C.GDT_UInt16:
		return reflect.Uint16
	case C.GDT_Int16:
		return reflect.Int16
	case C.GDT_UInt32:
		return reflect.Uint32
	case C.GDT_Int32:
		return reflect.Int32
	case C.GDT_Float32:
		return reflect.Float32
	case C.GDT_Float64:
		return reflect.Float64
	case C.GDT_CInt16:
		// invalid
	case C.GDT_CInt32:
		// invalid
	case C.GDT_CFloat32:
		// invalid
	case C.GDT_CFloat64:
		// invalid
	case C.GDT_TypeCount:
		// invalid
	}
	return reflect.Invalid
}
