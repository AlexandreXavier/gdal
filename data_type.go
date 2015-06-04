// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include "cgo_gdal.h"
import "C"

type DataType int

const (
	GDT_Unknown   DataType = iota // Unknown or unspecified type
	GDT_Byte                      // Eight bit unsigned integer
	GDT_UInt16                    // Sixteen bit unsigned integer
	GDT_Int16                     // Sixteen bit signed integer
	GDT_UInt32                    // Thirty two bit unsigned integer
	GDT_Int32                     // Thirty two bit signed integer
	GDT_Float32                   // Thirty two bit floating point
	GDT_Float64                   // Sixty four bit floating point
	GDT_CInt16                    // Complex Int16
	GDT_CInt32                    // Complex Int32
	GDT_CFloat32                  // Complex Float32
	GDT_CFloat64                  // Complex Float64
	GDT_TypeCount                 // maximum type # + 1
)

const (
	_GDT_Unknown   = DataType(C.GDT_Unknown)
	_GDT_Byte      = DataType(C.GDT_Byte)
	_GDT_UInt16    = DataType(C.GDT_UInt16)
	_GDT_Int16     = DataType(C.GDT_Int16)
	_GDT_UInt32    = DataType(C.GDT_UInt32)
	_GDT_Int32     = DataType(C.GDT_Int32)
	_GDT_Float32   = DataType(C.GDT_Float32)
	_GDT_Float64   = DataType(C.GDT_Float64)
	_GDT_CInt16    = DataType(C.GDT_CInt16)
	_GDT_CInt32    = DataType(C.GDT_CInt32)
	_GDT_CFloat32  = DataType(C.GDT_CFloat32)
	_GDT_CFloat64  = DataType(C.GDT_CFloat64)
	_GDT_TypeCount = DataType(C.GDT_TypeCount)
)

func (d DataType) Valid() bool {
	return d > GDT_Unknown && d < GDT_TypeCount
}

func (d DataType) Depth() int {
	switch d {
	case GDT_Unknown:
		return 0
	case GDT_Byte:
		return 1 * 8
	case GDT_UInt16:
		return 2 * 8
	case GDT_Int16:
		return 4 * 8
	case GDT_UInt32:
		return 4 * 8
	case GDT_Int32:
		return 4 * 8
	case GDT_Float32:
		return 4 * 8
	case GDT_Float64:
		return 8 * 8
	case GDT_CInt16:
		return 2 * 8
	case GDT_CInt32:
		return 4 * 8
	case GDT_CFloat32:
		return 4 * 8
	case GDT_CFloat64:
		return 8 * 8
	case GDT_TypeCount:
		return 0
	}
	return 0
}
