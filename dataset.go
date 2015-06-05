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

type Dataset struct {
	Filename string
	Width    int
	Height   int
	Channels int
	DataType DataType
	Opt      *Options

	poDataset C.GDALDatasetH
	cBuf      *C.uint8_t
	cBufLen   int
}

func OpenImage(filename string, readOnly bool) (m *Dataset, err error) {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	m = new(Dataset)
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

	m.Filename = filename
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

func CreateImage(filename string, width, height, channels int, dataType DataType, opt *Options) (m *Dataset, err error) {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	m = &Dataset{
		Filename: filename,
		Width:    width,
		Height:   height,
		Channels: channels,
		DataType: dataType,
		Opt:      new(Options),
	}

	if opt != nil {
		*m.Opt = *opt
		m.Opt.ExtOptions = make(map[string]string)
		if len(opt.ExtOptions) != 0 {
			for k, v := range opt.ExtOptions {
				m.Opt.ExtOptions[k] = v
			}
		}
	}
	if m.Opt.DriverName == "" {
		m.Opt.DriverName = getDriverName(filename)
	}

	cDriverName := C.CString(m.Opt.DriverName)
	defer C.free(unsafe.Pointer(cDriverName))

	poDriver := C.GDALGetDriverByName(cDriverName)
	if poDriver == nil {
		err = fmt.Errorf("gdal: CreateImage(%q) failed.", filename)
		return
	}

	// TODO: support ExtOpt
	m.poDataset = C.GDALCreate(poDriver, cname,
		C.int(width), C.int(height), C.int(channels),
		C.GDALDataType(dataType), nil,
	)
	if m.poDataset == nil {
		err = fmt.Errorf("gdal: CreateImage(%q) failed.", filename)
		return
	}

	return
}

func (p *Dataset) Close() error {
	if p.poDataset != nil {
		C.GDALClose(p.poDataset)
		p.poDataset = nil
	}
	if p.cBuf != nil {
		C.free(unsafe.Pointer(p.cBuf))
		p.cBuf = nil
	}
	*p = Dataset{}
	return nil
}

func (p *Dataset) Read(r image.Rectangle, data []byte, stride int) error {
	pixelSize := p.Channels * p.DataType.Depth() / 8
	if stride <= 0 {
		stride = r.Dx() * pixelSize
	}
	data = data[:r.Dy()*stride]

	if p.cBufLen < len(data) {
		if p.cBuf != nil {
			C.free(unsafe.Pointer(p.cBuf))
			p.cBuf = nil
		}
		p.cBuf = (*C.uint8_t)(C.malloc(C.size_t(p.cBufLen)))
		p.cBufLen = len(data)
	}

	for nBandId := 0; nBandId < p.Channels; nBandId++ {
		pBand := C.GDALGetRasterBand(p.poDataset, C.int(nBandId+1))
		cErr := C.GDALRasterIO(pBand, C.GF_Read,
			C.int(r.Min.X), C.int(r.Min.Y), C.int(r.Dx()), C.int(r.Dy()),
			unsafe.Pointer(p.cBuf), C.int(r.Dx()), C.int(r.Dy()),
			C.GDALDataType(p.DataType), C.int(pixelSize),
			C.int(stride),
		)
		if cErr != C.CE_None {
			return fmt.Errorf("gdal: Dataset.Read(%q) failed.", p.Filename)
		}
	}

	copy(data, ((*[1 << 30]byte)(unsafe.Pointer(p.cBuf)))[0:len(data):len(data)])
	nativeToBigEndian(data, p.DataType.Depth())

	return nil
}

func (p *Dataset) Write(r image.Rectangle, data []byte, stride int) error {
	pixelSize := p.Channels * p.DataType.Depth() / 8
	if stride <= 0 {
		stride = r.Dx() * pixelSize
	}
	data = data[:r.Dy()*stride]

	if p.cBufLen < len(data) {
		if p.cBuf != nil {
			C.free(unsafe.Pointer(p.cBuf))
			p.cBuf = nil
		}
		p.cBuf = (*C.uint8_t)(C.malloc(C.size_t(p.cBufLen)))
		p.cBufLen = len(data)
	}

	for nBandId := 0; nBandId < p.Channels; nBandId++ {
		pBand := C.GDALGetRasterBand(p.poDataset, C.int(nBandId+1))
		cErr := C.GDALRasterIO(pBand, C.GF_Write,
			C.int(r.Min.X), C.int(r.Min.Y), C.int(r.Dx()), C.int(r.Dy()),
			unsafe.Pointer(p.cBuf), C.int(r.Dx()), C.int(r.Dy()),
			C.GDALDataType(p.DataType), C.int(pixelSize),
			C.int(stride),
		)
		if cErr != C.CE_None {
			return fmt.Errorf("gdal: Dataset.Read(%q) failed.", p.Filename)
		}
	}

	copy(data, ((*[1 << 30]byte)(unsafe.Pointer(p.cBuf)))[0:len(data):len(data)])
	nativeToBigEndian(data, p.DataType.Depth())

	return nil
}
