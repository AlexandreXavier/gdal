// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include "cgo_gdal.h"
import "C"
import (
	_ "fmt"
	_ "runtime"
	"unsafe"
)

/* -------------------------------------------------------------------- */
/*      Significant constants.                                          */
/* -------------------------------------------------------------------- */

// Pixel data types
type _GDALDataType int

const (
	_GDT_Unknown   = _GDALDataType(C.GDT_Unknown)   // Unknown or unspecified type
	_GDT_Byte      = _GDALDataType(C.GDT_Byte)      // Eight bit unsigned integer
	_GDT_UInt16    = _GDALDataType(C.GDT_UInt16)    // Sixteen bit unsigned integer
	_GDT_Int16     = _GDALDataType(C.GDT_Int16)     // Sixteen bit signed integer
	_GDT_UInt32    = _GDALDataType(C.GDT_UInt32)    // Thirty two bit unsigned integer
	_GDT_Int32     = _GDALDataType(C.GDT_Int32)     // Thirty two bit signed integer
	_GDT_Float32   = _GDALDataType(C.GDT_Float32)   // Thirty two bit floating point
	_GDT_Float64   = _GDALDataType(C.GDT_Float64)   // Sixty four bit floating point
	_GDT_CInt16    = _GDALDataType(C.GDT_CInt16)    // Complex Int16
	_GDT_CInt32    = _GDALDataType(C.GDT_CInt32)    // Complex Int32
	_GDT_CFloat32  = _GDALDataType(C.GDT_CFloat32)  // Complex Float32
	_GDT_CFloat64  = _GDALDataType(C.GDT_CFloat64)  // Complex Float64
	_GDT_TypeCount = _GDALDataType(C.GDT_TypeCount) // maximum type # + 1
)

// Get data type size in bits.
//
// Returns the size of a a GDT_* type in bits, not bytes!
func _GDALGetDataTypeSize(dataType _GDALDataType) int {
	return int(C.GDALGetDataTypeSize(C.GDALDataType(dataType)))
}
func _GDALDataTypeIsComplex(dataType _GDALDataType) int {
	return int(C.GDALDataTypeIsComplex(C.GDALDataType(dataType)))
}
func _GDALGetDataTypeName(dataType _GDALDataType) string {
	return C.GoString(C.GDALGetDataTypeName(C.GDALDataType(dataType)))
}
func _GDALGetDataTypeByName(dataTypeName string) _GDALDataType {
	name := C.CString(dataTypeName)
	defer C.free(unsafe.Pointer(name))
	return _GDALDataType(C.GDALGetDataTypeByName(name))
}
func _GDALDataTypeUnion(dataTypeA, dataTypeB _GDALDataType) _GDALDataType {
	return _GDALDataType(
		C.GDALDataTypeUnion(C.GDALDataType(dataTypeA), C.GDALDataType(dataTypeB)),
	)
}

// status of the asynchronous stream
type _GDALAsyncStatusType int

const (
	_GARIO_PENDING   = _GDALAsyncStatusType(C.GARIO_PENDING)
	_GARIO_UPDATE    = _GDALAsyncStatusType(C.GARIO_UPDATE)
	_GARIO_ERROR     = _GDALAsyncStatusType(C.GARIO_ERROR)
	_GARIO_COMPLETE  = _GDALAsyncStatusType(C.GARIO_COMPLETE)
	_GARIO_TypeCount = _GDALAsyncStatusType(C.GARIO_TypeCount)
)

func _GDALGetAsyncStatusTypeName(statusType _GDALAsyncStatusType) string {
	return C.GoString(C.GDALGetAsyncStatusTypeName(C.GDALAsyncStatusType(statusType)))
}
func _GDALGetAsyncStatusTypeByName(statusTypeName string) _GDALAsyncStatusType {
	name := C.CString(statusTypeName)
	defer C.free(unsafe.Pointer(name))
	return _GDALAsyncStatusType(C.GDALGetAsyncStatusTypeByName(name))
}

// Flag indicating read/write, or read-only access to data.
type _GDALAccess int

const (
	_GA_ReadOnly = _GDALAccess(C.GA_ReadOnly) // Read only (no update) access
	_GA_Update   = _GDALAccess(C.GA_Update)   // Read/write access.
)

// Read/Write flag for RasterIO() method
type _GDALRWFlag int

const (
	_GF_Read  = _GDALRWFlag(C.GF_Read)  // Read data
	_GF_Write = _GDALRWFlag(C.GF_Write) // Write data
)

// Types of color interpretation for raster bands.
type _GDALColorInterp int

const (
	_GCI_Undefined      = _GDALColorInterp(C.GCI_Undefined)      // Undefined
	_GCI_GrayIndex      = _GDALColorInterp(C.GCI_GrayIndex)      // Greyscale
	_GCI_PaletteIndex   = _GDALColorInterp(C.GCI_PaletteIndex)   // Paletted (see associated color table)
	_GCI_RedBand        = _GDALColorInterp(C.GCI_RedBand)        // Red band of RGBA image
	_GCI_GreenBand      = _GDALColorInterp(C.GCI_GreenBand)      // Green band of RGBA image
	_GCI_BlueBand       = _GDALColorInterp(C.GCI_BlueBand)       // Blue band of RGBA image
	_GCI_AlphaBand      = _GDALColorInterp(C.GCI_AlphaBand)      // Alpha (0=transparent, 255=opaque)
	_GCI_HueBand        = _GDALColorInterp(C.GCI_HueBand)        // Hue band of HLS image
	_GCI_SaturationBand = _GDALColorInterp(C.GCI_SaturationBand) // Saturation band of HLS image
	_GCI_LightnessBand  = _GDALColorInterp(C.GCI_LightnessBand)  // Lightness band of HLS image
	_GCI_CyanBand       = _GDALColorInterp(C.GCI_CyanBand)       // Cyan band of CMYK image
	_GCI_MagentaBand    = _GDALColorInterp(C.GCI_MagentaBand)    // Magenta band of CMYK image
	_GCI_YellowBand     = _GDALColorInterp(C.GCI_YellowBand)     // Yellow band of CMYK image
	_GCI_BlackBand      = _GDALColorInterp(C.GCI_BlackBand)      // Black band of CMLY image
	_GCI_YCbCr_YBand    = _GDALColorInterp(C.GCI_YCbCr_YBand)    // Y Luminance
	_GCI_YCbCr_CbBand   = _GDALColorInterp(C.GCI_YCbCr_CbBand)   // Cb Chroma
	_GCI_YCbCr_CrBand   = _GDALColorInterp(C.GCI_YCbCr_CrBand)   // Cr Chroma
	_GCI_Max            = _GDALColorInterp(C.GCI_Max)            // Max current value
)

func _GDALGetColorInterpretationName(colorInterp _GDALColorInterp) string {
	return C.GoString(C.GDALGetColorInterpretationName(C.GDALColorInterp(colorInterp)))
}
func _GDALGetColorInterpretationByName(pszName string) _GDALColorInterp {
	name := C.CString(pszName)
	defer C.free(unsafe.Pointer(name))
	return _GDALColorInterp(C.GDALGetColorInterpretationByName(name))
}

// Types of color interpretations for a GDALColorTable.
type _GDALPaletteInterp int

const (
	_GPI_Gray = _GDALPaletteInterp(C.GPI_Gray) // Grayscale (in GDALColorEntry.c1)
	_GPI_RGB  = _GDALPaletteInterp(C.GPI_RGB)  // Red, Green, Blue and Alpha in (in c1, c2, c3 and c4)
	_GPI_CMYK = _GDALPaletteInterp(C.GPI_CMYK) // Cyan, Magenta, Yellow and Black (in c1, c2, c3 and c4)
	_GPI_HLS  = _GDALPaletteInterp(C.GPI_HLS)  // Hue, Lightness and Saturation (in c1, c2, and c3)
)

func _GDALGetPaletteInterpretationName(paletteInterp _GDALPaletteInterp) string {
	return C.GoString(
		C.GDALGetPaletteInterpretationName(C.GDALPaletteInterp(paletteInterp)),
	)
}

// "well known" metadata items.
const (
	_GDALMD_AREA_OR_POINT = string(C.GDALMD_AREA_OR_POINT)
	_GDALMD_AOP_AREA      = string(C.GDALMD_AOP_AREA)
	_GDALMD_AOP_POINT     = string(C.GDALMD_AOP_POINT)
)

/* -------------------------------------------------------------------- */
/*      GDAL Specific error codes.                                      */
/*                                                                      */
/*      error codes 100 to 299 reserved for GDAL.                       */
/* -------------------------------------------------------------------- */
const _CPLE_WrongFormat = int(C.CPLE_WrongFormat)

/* -------------------------------------------------------------------- */
/*      Define handle types related to various internal classes.        */
/* -------------------------------------------------------------------- */

// Opaque type used for the C bindings of the C++ GDALMajorObject class
type _GDALMajorObjectH C.GDALMajorObjectH

// Opaque type used for the C bindings of the C++ GDALDataset class
type _GDALDatasetH C.GDALDatasetH

// Opaque type used for the C bindings of the C++ GDALRasterBand class
type _GDALRasterBandH C.GDALRasterBandH

// Opaque type used for the C bindings of the C++ GDALDriver class
type _GDALDriverH C.GDALDriverH

// Opaque type used for the C bindings of the C++ GDALColorTable class
type _GDALColorTableH C.GDALColorTableH

// Opaque type used for the C bindings of the C++ GDALRasterAttributeTable class
type _GDALRasterAttributeTableH C.GDALRasterAttributeTableH

// Opaque type used for the C bindings of the C++ GDALAsyncReader class
type _GDALAsyncReaderH C.GDALAsyncReaderH

/* -------------------------------------------------------------------- */
/*      Callback "progress" function.                                   */
/* -------------------------------------------------------------------- */

type _GDALProgressFunc func(dfComplete float64, pszMessage string, pProgressArg interface{}) int

func _GDALDummyProgress(dfComplete float64, pszMessage string, pData interface{}) int {
	msg := C.CString(pszMessage)
	defer C.free(unsafe.Pointer(msg))

	rv := C.GDALDummyProgress(C.double(dfComplete), msg, unsafe.Pointer(nil))
	return int(rv)
}
func _GDALTermProgress(dfComplete float64, pszMessage string, pData interface{}) int {
	msg := C.CString(pszMessage)
	defer C.free(unsafe.Pointer(msg))

	rv := C.GDALTermProgress(C.double(dfComplete), msg, unsafe.Pointer(nil))
	return int(rv)
}
func _GDALScaledProgress(dfComplete float64, pszMessage string, pData interface{}) int {
	msg := C.CString(pszMessage)
	defer C.free(unsafe.Pointer(msg))

	rv := C.GDALScaledProgress(C.double(dfComplete), msg, unsafe.Pointer(nil))
	return int(rv)
}

func _GDALCreateScaledProgress(dfMin, dfMax float64, pfnProgress _GDALProgressFunc, pData unsafe.Pointer) unsafe.Pointer {
	panic("not impl")
	return nil
}

func _GDALDestroyScaledProgress(pData unsafe.Pointer) {
	C.GDALDestroyScaledProgress(pData)
}

// -----------------------------------------------------------------------

type goGDALProgressFuncProxyArgs struct {
	progresssFunc _GDALProgressFunc
	pData         interface{}
}

//export goGDALProgressFuncProxyA
func goGDALProgressFuncProxyA(dfComplete C.double, pszMessage *C.char, pData *interface{}) int {
	if arg, ok := (*pData).(goGDALProgressFuncProxyArgs); ok {
		return arg.progresssFunc(
			float64(dfComplete), C.GoString(pszMessage), arg.pData,
		)
	}
	return 0
}

/* ==================================================================== */
/*      Registration/driver related.                                    */
/* ==================================================================== */

const (
	_GDAL_DMD_LONGNAME           = string(C.GDAL_DMD_LONGNAME)
	_GDAL_DMD_HELPTOPIC          = string(C.GDAL_DMD_HELPTOPIC)
	_GDAL_DMD_MIMETYPE           = string(C.GDAL_DMD_MIMETYPE)
	_GDAL_DMD_EXTENSION          = string(C.GDAL_DMD_EXTENSION)
	_GDAL_DMD_CREATIONOPTIONLIST = string(C.GDAL_DMD_CREATIONOPTIONLIST)
	_GDAL_DMD_CREATIONDATATYPES  = string(C.GDAL_DMD_CREATIONDATATYPES)

	_GDAL_DCAP_CREATE     = string(C.GDAL_DCAP_CREATE)
	_GDAL_DCAP_CREATECOPY = string(C.GDAL_DCAP_CREATECOPY)
	_GDAL_DCAP_VIRTUALIO  = string(C.GDAL_DCAP_VIRTUALIO)
)

func gdalAllRegister() {
	C.GDALAllRegister()
}

// Create a new dataset with this driver.
func _GDALCreate(hDriver _GDALDriverH,
	pszFilename string,
	nXSize, nYSize, nBands int,
	dataType _GDALDataType,
	papszOptions []string,
) _GDALDatasetH {
	name := C.CString(pszFilename)
	defer C.free(unsafe.Pointer(name))

	opts := make([]*C.char, len(papszOptions)+1)
	for i := 0; i < len(papszOptions); i++ {
		opts[i] = C.CString(papszOptions[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[len(opts)-1] = (*C.char)(unsafe.Pointer(nil))

	h := C.GDALCreate(
		C.GDALDriverH(hDriver),
		name,
		C.int(nXSize), C.int(nYSize), C.int(nBands),
		C.GDALDataType(dataType),
		(**C.char)(unsafe.Pointer(&opts[0])),
	)
	return _GDALDatasetH(h)
}

// Create a copy of a dataset.
func _GDALCreateCopy(
	hDriver _GDALDriverH, pszFilename string,
	hSrcDS _GDALDatasetH,
	bStrict int, papszOptions []string,
	pfnProgress _GDALProgressFunc, pProgressData interface{},
) _GDALDatasetH {
	name := C.CString(pszFilename)
	defer C.free(unsafe.Pointer(name))

	opts := make([]*C.char, len(papszOptions)+1)
	for i := 0; i < len(papszOptions); i++ {
		opts[i] = C.CString(papszOptions[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[len(opts)-1] = (*C.char)(unsafe.Pointer(nil))

	arg := &goGDALProgressFuncProxyArgs{
		pfnProgress, pProgressData,
	}

	h := C.GDALCreateCopy(
		C.GDALDriverH(hDriver), name,
		C.GDALDatasetH(hSrcDS),
		C.int(bStrict), (**C.char)(unsafe.Pointer(&opts[0])),
		C.cgoGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return _GDALDatasetH(h)
}

/* ==================================================================== */
/*      GDAL_GCP                                                        */
/* ==================================================================== */

/* ==================================================================== */
/*      major objects (dataset, and, driver, drivermanager).            */
/* ==================================================================== */

/* ==================================================================== */
/*      GDALDataset class ... normally this represents one file.        */
/* ==================================================================== */

/* ==================================================================== */
/*      GDALRasterBand ... one band/channel in a dataset.               */
/* ==================================================================== */

/* ==================================================================== */
/*     GDALAsyncReader                                                  */
/* ==================================================================== */

/* ==================================================================== */
/*      Color tables.                                                   */
/* ==================================================================== */

/* ==================================================================== */
/*      Raster Attribute Table						*/
/* ==================================================================== */

/* ==================================================================== */
/*      GDAL Cache Management                                           */
/* ==================================================================== */
