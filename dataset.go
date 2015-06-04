// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include "cgo_gdal.h"
import "C"
import (
	"fmt"
	"image"
	"unsafe"
)

// GDAL Raster Formats
//
// See http://www.gdal.org/formats_list.html
type Options struct {
	DriverName string
	Projection string
	Transform  [6]float64
	ExtOptions map[string]string
}

type Image struct {
	Width    int
	Height   int
	Channels int
	DataType DataType
	Opt      *Options

	poDataset C.GDALDatasetH
}

func OpenImage(filename string, readOnly bool) (m *Image, err error) {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	m = new(Image)
	m.Opt = new(Options)

	if readOnly {
		m.poDataset = C.GDALOpen(cname, C.GA_ReadOnly)
	} else {
		m.poDataset = C.GDALOpen(cname, C.GA_Update)
	}
	if m.poDataset == nil {
		err = fmt.Errorf("gdal: OpenImage(%q) failed.", filename)
		return
	}

	m.Width = int(C.GDALGetRasterXSize(m.poDataset))
	m.Height = int(C.GDALGetRasterYSize(m.poDataset))
	m.Channels = int(C.GDALGetRasterCount(m.poDataset))
	m.DataType = DataType(C.GDALGetRasterDataType(C.GDALGetRasterBand(m.poDataset, 1)))

	m.Opt.DriverName = C.GoString(C.GDALGetDriverShortName(C.GDALGetDatasetDriver(m.poDataset)))
	m.Opt.Projection = C.GoString(C.GDALGetProjectionRef(m.poDataset))
	m.Opt.ExtOptions = make(map[string]string)

	var padfTransform [6]C.double
	if C.GDALGetGeoTransform(m.poDataset, &padfTransform[0]) == C.CE_None {
		for i := 0; i < len(padfTransform); i++ {
			m.Opt.Transform[i] = float64(padfTransform[i])
		}
	}

	return
}

func CreateImage(filename string, width, height, channels int, dataType DataType, opt *Options) (m *Image, err error) {
	err = fmt.Errorf("gdal: CreateImage, TODO")
	return
}

func (p *Image) Close() error {
	if p.poDataset != nil {
		C.GDALClose(p.poDataset)
		p.poDataset = nil
	}
	return nil
}

func (p *Image) Read(r image.Rectangle, data []byte, stride int) error {
	return fmt.Errorf("gdal: Image.Read, TODO")
}

func (p *Image) Write(r image.Rectangle, data []byte, stride int) error {
	return fmt.Errorf("gdal: Image.Write, TODO")
}
