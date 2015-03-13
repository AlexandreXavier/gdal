// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

/*
#include "go_gdal.h"

#cgo linux  pkg-config: gdal
#cgo darwin pkg-config: gdal
#cgo windows LDFLAGS: -lgdal.dll
*/
import "C"
import (
	_ "runtime"
	"unsafe"
	_ "fmt"
)

func init() {
	gdalAllRegister()
}

/* -------------------------------------------------------------------- */
/*      Significant constants.                                          */
/* -------------------------------------------------------------------- */

// Pixel data types
type GDALDataType int

const (
	// Unknown or unspecified type
	GDT_Unknown = GDALDataType(C.GDT_Unknown)
	// Eight bit unsigned integer
	GDT_Byte = GDALDataType(C.GDT_Byte)
	// Sixteen bit unsigned integer
	GDT_UInt16 = GDALDataType(C.GDT_UInt16)
	// Sixteen bit signed integer
	GDT_Int16 = GDALDataType(C.GDT_Int16)
	// Thirty two bit unsigned integer
	GDT_UInt32 = GDALDataType(C.GDT_UInt32)
	// Thirty two bit signed integer
	GDT_Int32 = GDALDataType(C.GDT_Int32)
	// Thirty two bit floating point
	GDT_Float32 = GDALDataType(C.GDT_Float32)
	// Sixty four bit floating point
	GDT_Float64 = GDALDataType(C.GDT_Float64)
	// Complex Int16
	GDT_CInt16 = GDALDataType(C.GDT_CInt16)
	// Complex Int32
	GDT_CInt32 = GDALDataType(C.GDT_CInt32)
	// Complex Float32
	GDT_CFloat32 = GDALDataType(C.GDT_CFloat32)
	// Complex Float64
	GDT_CFloat64 = GDALDataType(C.GDT_CFloat64)
	// maximum type # + 1
	GDT_TypeCount = GDALDataType(C.GDT_TypeCount)
)

// Get data type size in bits.
//
// Returns the size of a a GDT_* type in bits, not bytes!
func GDALGetDataTypeSize(dataType GDALDataType) int {
	return int(C.GDALGetDataTypeSize(C.GDALDataType(dataType)))
}
func GDALDataTypeIsComplex(dataType GDALDataType) int {
	return int(C.GDALDataTypeIsComplex(C.GDALDataType(dataType)))
}
func GDALGetDataTypeName(dataType GDALDataType) string {
	return C.GoString(C.GDALGetDataTypeName(C.GDALDataType(dataType)))
}
func GDALGetDataTypeByName(dataTypeName string) GDALDataType {
	name := C.CString(dataTypeName)
	defer C.free(unsafe.Pointer(name))
	return GDALDataType(C.GDALGetDataTypeByName(name))
}
func GDALDataTypeUnion(dataTypeA, dataTypeB GDALDataType) GDALDataType {
	return GDALDataType(
		C.GDALDataTypeUnion(C.GDALDataType(dataTypeA), C.GDALDataType(dataTypeB)),
	)
}


// status of the asynchronous stream
type GDALAsyncStatusType int

const (
	GARIO_PENDING = GDALAsyncStatusType(C.GARIO_PENDING)
	GARIO_UPDATE = GDALAsyncStatusType(C.GARIO_UPDATE)
	GARIO_ERROR = GDALAsyncStatusType(C.GARIO_ERROR)
	GARIO_COMPLETE = GDALAsyncStatusType(C.GARIO_COMPLETE)
	GARIO_TypeCount = GDALAsyncStatusType(C.GARIO_TypeCount)
)

func GDALGetAsyncStatusTypeName(statusType GDALAsyncStatusType) string {
	return C.GoString(C.GDALGetAsyncStatusTypeName(C.GDALAsyncStatusType(statusType)))
}
func GDALGetAsyncStatusTypeByName(statusTypeName string) GDALAsyncStatusType {
	name := C.CString(statusTypeName)
	defer C.free(unsafe.Pointer(name))
	return GDALAsyncStatusType(C.GDALGetAsyncStatusTypeByName(name))
}

// Flag indicating read/write, or read-only access to data.
type GDALAccess int

const (
	// Read only (no update) access
	GA_ReadOnly = GDALAccess(C.GA_ReadOnly)
	// Read/write access.
	GA_Update = GDALAccess(C.GA_Update)
)


// Read/Write flag for RasterIO() method
type GDALRWFlag int

const (
	// Read data
	GF_Read = GDALRWFlag(C.GF_Read)
	// Write data
	GF_Write = GDALRWFlag(C.GF_Write)
)


// Types of color interpretation for raster bands.
type GDALColorInterp int

const (
	GCI_Undefined = GDALColorInterp(C.GCI_Undefined)

	// Greyscale
	GCI_GrayIndex = GDALColorInterp(C.GCI_GrayIndex)
	// Paletted (see associated color table)
	GCI_PaletteIndex = GDALColorInterp(C.GCI_PaletteIndex)
	// Red band of RGBA image
	GCI_RedBand = GDALColorInterp(C.GCI_RedBand)
	// Green band of RGBA image
	GCI_GreenBand = GDALColorInterp(C.GCI_GreenBand)
	// Blue band of RGBA image
	GCI_BlueBand = GDALColorInterp(C.GCI_BlueBand)
	// Alpha (0=transparent, 255=opaque)
	GCI_AlphaBand = GDALColorInterp(C.GCI_AlphaBand)
	// Hue band of HLS image
	GCI_HueBand = GDALColorInterp(C.GCI_HueBand)
	// Saturation band of HLS image
	GCI_SaturationBand = GDALColorInterp(C.GCI_SaturationBand)
	// Lightness band of HLS image
	GCI_LightnessBand = GDALColorInterp(C.GCI_LightnessBand)
	// Cyan band of CMYK image
	GCI_CyanBand = GDALColorInterp(C.GCI_CyanBand)
	// Magenta band of CMYK image
	GCI_MagentaBand = GDALColorInterp(C.GCI_MagentaBand)
	// Yellow band of CMYK image
	GCI_YellowBand = GDALColorInterp(C.GCI_YellowBand)
	// Black band of CMLY image
	GCI_BlackBand = GDALColorInterp(C.GCI_BlackBand)
	// Y Luminance
	GCI_YCbCr_YBand = GDALColorInterp(C.GCI_YCbCr_YBand)
	// Cb Chroma
	GCI_YCbCr_CbBand = GDALColorInterp(C.GCI_YCbCr_CbBand)
	// Cr Chroma
	GCI_YCbCr_CrBand = GDALColorInterp(C.GCI_YCbCr_CrBand)
	// Max current value
	GCI_Max = GDALColorInterp(C.GCI_Max)
)

func GDALGetColorInterpretationName(colorInterp GDALColorInterp) string {
	return C.GoString(C.GDALGetColorInterpretationName(C.GDALColorInterp(colorInterp)))
}
func GDALGetColorInterpretationByName(pszName string) GDALColorInterp {
	name := C.CString(pszName)
	defer C.free(unsafe.Pointer(name))
	return GDALColorInterp(C.GDALGetColorInterpretationByName(name))
}


// Types of color interpretations for a GDALColorTable.
type GDALPaletteInterp int

const (
	// Grayscale (in GDALColorEntry.c1)
	GPI_Gray = GDALPaletteInterp(C.GPI_Gray)
	// Red, Green, Blue and Alpha in (in c1, c2, c3 and c4)
	GPI_RGB = GDALPaletteInterp(C.GPI_RGB)
	// Cyan, Magenta, Yellow and Black (in c1, c2, c3 and c4)
	GPI_CMYK = GDALPaletteInterp(C.GPI_CMYK)
	// Hue, Lightness and Saturation (in c1, c2, and c3)
	GPI_HLS = GDALPaletteInterp(C.GPI_HLS)
)

func GDALGetPaletteInterpretationName(paletteInterp GDALPaletteInterp) string {
	return C.GoString(
		C.GDALGetPaletteInterpretationName(C.GDALPaletteInterp(paletteInterp)),
	)
}


// "well known" metadata items.
const (
	GDALMD_AREA_OR_POINT = string(C.GDALMD_AREA_OR_POINT)
	GDALMD_AOP_AREA = string(C.GDALMD_AOP_AREA)
	GDALMD_AOP_POINT = string(C.GDALMD_AOP_POINT)
)

/* -------------------------------------------------------------------- */
/*      GDAL Specific error codes.                                      */
/*                                                                      */
/*      error codes 100 to 299 reserved for GDAL.                       */
/* -------------------------------------------------------------------- */
const CPLE_WrongFormat = int(C.CPLE_WrongFormat)

/* -------------------------------------------------------------------- */
/*      Define handle types related to various internal classes.        */
/* -------------------------------------------------------------------- */

// Opaque type used for the C bindings of the C++ GDALMajorObject class
type GDALMajorObjectH C.GDALMajorObjectH

// Opaque type used for the C bindings of the C++ GDALDataset class
type GDALDatasetH C.GDALDatasetH

// Opaque type used for the C bindings of the C++ GDALRasterBand class
type GDALRasterBandH C.GDALRasterBandH

// Opaque type used for the C bindings of the C++ GDALDriver class
type GDALDriverH C.GDALDriverH

// Opaque type used for the C bindings of the C++ GDALColorTable class
type GDALColorTableH C.GDALColorTableH

// Opaque type used for the C bindings of the C++ GDALRasterAttributeTable class
type GDALRasterAttributeTableH C.GDALRasterAttributeTableH

// Opaque type used for the C bindings of the C++ GDALAsyncReader class
type GDALAsyncReaderH C.GDALAsyncReaderH

/* -------------------------------------------------------------------- */
/*      Callback "progress" function.                                   */
/* -------------------------------------------------------------------- */

type GDALProgressFunc func(dfComplete float64, pszMessage string, pProgressArg interface{}) int

func GDALDummyProgress(dfComplete float64, pszMessage string, pData interface{}) int {
	msg := C.CString(pszMessage)
	defer C.free(unsafe.Pointer(msg))

	rv := C.GDALDummyProgress(C.double(dfComplete), msg, unsafe.Pointer(nil))
	return int(rv)
}
func GDALTermProgress(dfComplete float64, pszMessage string, pData interface{}) int {
	msg := C.CString(pszMessage)
	defer C.free(unsafe.Pointer(msg))

	rv := C.GDALTermProgress(C.double(dfComplete), msg, unsafe.Pointer(nil))
	return int(rv)
}
func GDALScaledProgress(dfComplete float64, pszMessage string, pData interface{}) int {
	msg := C.CString(pszMessage)
	defer C.free(unsafe.Pointer(msg))

	rv := C.GDALScaledProgress(C.double(dfComplete), msg, unsafe.Pointer(nil))
	return int(rv)
}

func GDALCreateScaledProgress(dfMin, dfMax float64, pfnProgress GDALProgressFunc, pData unsafe.Pointer) unsafe.Pointer {
	panic("not impl")
	return nil
}

func GDALDestroyScaledProgress(pData unsafe.Pointer) {
	C.GDALDestroyScaledProgress(pData)
}

// -----------------------------------------------------------------------

type goGDALProgressFuncProxyArgs struct {
	progresssFunc	GDALProgressFunc
	pData		interface{}
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
	GDAL_DMD_LONGNAME = string(C.GDAL_DMD_LONGNAME)
	GDAL_DMD_HELPTOPIC = string(C.GDAL_DMD_HELPTOPIC)
	GDAL_DMD_MIMETYPE = string(C.GDAL_DMD_MIMETYPE)
	GDAL_DMD_EXTENSION = string(C.GDAL_DMD_EXTENSION)
	GDAL_DMD_CREATIONOPTIONLIST = string(C.GDAL_DMD_CREATIONOPTIONLIST)
	GDAL_DMD_CREATIONDATATYPES = string(C.GDAL_DMD_CREATIONDATATYPES)

	GDAL_DCAP_CREATE = string(C.GDAL_DCAP_CREATE)
	GDAL_DCAP_CREATECOPY = string(C.GDAL_DCAP_CREATECOPY)
	GDAL_DCAP_VIRTUALIO = string(C.GDAL_DCAP_VIRTUALIO)
)

func gdalAllRegister() {
	C.GDALAllRegister()
}

// Create a new dataset with this driver.
func GDALCreate(hDriver GDALDriverH,
	pszFilename string,
	nXSize, nYSize, nBands int,
	dataType GDALDataType,
	papszOptions []string,
) GDALDatasetH {
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
	return GDALDatasetH(h)
}

// Create a copy of a dataset.
func GDALCreateCopy(
	hDriver GDALDriverH, pszFilename string,
	hSrcDS GDALDatasetH,
	bStrict int, papszOptions []string,
	pfnProgress GDALProgressFunc, pProgressData interface{},
) GDALDatasetH {
	name := C.CString(pszFilename)
	defer C.free(unsafe.Pointer(name))

	opts := make([]*C.char, len(papszOptions)+1)
	for i := 0; i < len(papszOptions); i++ {
		opts[i] = C.CString(papszOptions[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[len(opts)-1] = (*C.char)(unsafe.Pointer(nil))

	arg := &goGDALProgressFuncProxyArgs {
		pfnProgress, pProgressData,
	}

	h := C.GDALCreateCopy(
		C.GDALDriverH(hDriver), name,
		C.GDALDatasetH(hSrcDS),
		C.int(bStrict), (**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return GDALDatasetH(h)
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



