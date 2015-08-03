// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <gdal.h>
//#include <stdint.h>
import "C"
import (
	"fmt"
	"image"
	"os"
	"reflect"
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
	DataType reflect.Kind
	Opt      *Options

	poDataset C.GDALDatasetH
	cBuf      *C.uint8_t
	cBufLen   int
}

func OpenDataset(filename string, flag int) (p *Dataset, err error) {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	p = new(Dataset)
	p.Opt = new(Options)

	if flag == os.O_RDONLY {
		p.poDataset = C.GDALOpen(cname, C.GA_ReadOnly)
	} else {
		p.poDataset = C.GDALOpen(cname, C.GA_Update)
	}
	if p.poDataset == nil {
		err = fmt.Errorf("gdal: OpenImage(%q) failed.", filename)
		return
	}

	p.Filename = filename
	p.Width = int(C.GDALGetRasterXSize(p.poDataset))
	p.Height = int(C.GDALGetRasterYSize(p.poDataset))
	p.Channels = int(C.GDALGetRasterCount(p.poDataset))
	p.DataType = goDataType(C.GDALGetRasterDataType(C.GDALGetRasterBand(p.poDataset, 1)))

	p.Opt.DriverName = C.GoString(C.GDALGetDriverShortName(C.GDALGetDatasetDriver(p.poDataset)))
	p.Opt.Projection = C.GoString(C.GDALGetProjectionRef(p.poDataset))
	p.Opt.ExtOptions = make(map[string]string)

	var padfTransform [6]C.double
	if C.GDALGetGeoTransform(p.poDataset, &padfTransform[0]) == C.CE_None {
		for i := 0; i < len(padfTransform); i++ {
			p.Opt.Transform[i] = float64(padfTransform[i])
		}
	}

	return
}

func CreateDataset(filename string, width, height, channels int, dataType reflect.Kind, opt *Options) (p *Dataset, err error) {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	p = &Dataset{
		Filename: filename,
		Width:    width,
		Height:   height,
		Channels: channels,
		DataType: dataType,
		Opt:      new(Options),
	}

	if opt != nil {
		*p.Opt = *opt
		p.Opt.ExtOptions = make(map[string]string)
		if len(opt.ExtOptions) != 0 {
			for k, v := range opt.ExtOptions {
				p.Opt.ExtOptions[k] = v
			}
		}
	}
	if p.Opt.DriverName == "" {
		p.Opt.DriverName = getDefaultDriverNameByFilenameExt(filename)
	}

	cDriverName := C.CString(p.Opt.DriverName)
	defer C.free(unsafe.Pointer(cDriverName))

	cProjName := C.CString(p.Opt.Projection)
	defer C.free(unsafe.Pointer(cProjName))

	opts := make([]*C.char, len(p.Opt.ExtOptions)+1)
	optsList := make([]string, 0, len(p.Opt.ExtOptions))

	for k, v := range p.Opt.ExtOptions {
		optsList = append(optsList, k+":"+v)
	}
	for i := 0; i < len(optsList); i++ {
		opts[i] = C.CString(optsList[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}

	poDriver := C.GDALGetDriverByName(cDriverName)
	if poDriver == nil {
		err = fmt.Errorf("gdal: CreateImage(%q) failed.", filename)
		return
	}
	p.poDataset = C.GDALCreate(poDriver, cname,
		C.int(width), C.int(height), C.int(channels),
		gdalDataType(p.DataType), (**C.char)(unsafe.Pointer(&opts[0])),
	)
	if p.poDataset == nil {
		err = fmt.Errorf("gdal: CreateImage(%q) failed.", filename)
		return
	}

	var padfTransform [6]C.double
	for i := 0; i < len(padfTransform); i++ {
		padfTransform[i] = C.double(p.Opt.Transform[i])
	}
	if C.GDALSetProjection(p.poDataset, cProjName) != C.CE_None {
		// log warning
	}
	if C.GDALSetGeoTransform(p.poDataset, &padfTransform[0]) != C.CE_None {
		// log warning
	}

	return
}

func CreateDatasetCopy(filename string, src *Dataset, opt *Options) (p *Dataset, err error) {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	p = &Dataset{
		Filename: filename,
		Width:    src.Width,
		Height:   src.Height,
		Channels: src.Channels,
		DataType: src.DataType,
		Opt:      new(Options),
	}

	if opt != nil {
		*p.Opt = *opt
		p.Opt.ExtOptions = make(map[string]string)
		if len(src.Opt.ExtOptions) != 0 {
			for k, v := range src.Opt.ExtOptions {
				p.Opt.ExtOptions[k] = v
			}
		}
		if len(opt.ExtOptions) != 0 {
			for k, v := range opt.ExtOptions {
				p.Opt.ExtOptions[k] = v
			}
		}
	}
	if p.Opt.DriverName == "" {
		p.Opt.DriverName = getDefaultDriverNameByFilenameExt(filename)
	}

	opts := make([]*C.char, len(p.Opt.ExtOptions)+1)
	optsList := make([]string, 0, len(p.Opt.ExtOptions))

	for k, v := range p.Opt.ExtOptions {
		optsList = append(optsList, k+":"+v)
	}
	for i := 0; i < len(optsList); i++ {
		opts[i] = C.CString(optsList[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}

	cDriverName := C.CString(p.Opt.DriverName)
	defer C.free(unsafe.Pointer(cDriverName))

	poDriver := C.GDALGetDriverByName(cDriverName)
	if poDriver == nil {
		err = fmt.Errorf("gdal: CreateImage(%q) failed.", filename)
		return
	}

	p.poDataset = C.GDALCreateCopy(
		poDriver, cname, src.poDataset, C.FALSE,
		(**C.char)(unsafe.Pointer(&opts[0])),
		nil, nil,
	)
	if p.poDataset == nil {
		err = fmt.Errorf("gdal: CreateDatasetCopy(%q) failed.", filename)
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
	pixelSize := SizeofPixel(p.Channels, p.DataType)
	if stride == 0 {
		stride = r.Dx() * pixelSize
	}
	if n := r.Dx() * pixelSize; stride < n {
		return fmt.Errorf("gdal: Read, bad stride: %d", stride)
	}

	if n := stride * r.Dy(); p.cBufLen < n {
		p.cBufLen = n
		if p.cBuf != nil {
			C.free(unsafe.Pointer(p.cBuf))
			p.cBuf = nil
		}
	}
	if p.cBuf == nil {
		p.cBuf = (*C.uint8_t)(C.malloc(C.size_t(p.cBufLen)))
	}

	data = data[:r.Dy()*stride]
	cBuf := ((*[1 << 30]byte)(unsafe.Pointer(p.cBuf)))[0:len(data):len(data)]

	for nBandId := 0; nBandId < p.Channels; nBandId++ {
		pBand := C.GDALGetRasterBand(p.poDataset, C.int(nBandId+1))
		cErr := C.GDALRasterIO(pBand, C.GF_Read,
			C.int(r.Min.X), C.int(r.Min.Y), C.int(r.Dx()), C.int(r.Dy()),
			unsafe.Pointer(&cBuf[nBandId*SizeofKind(p.DataType)]), C.int(r.Dx()), C.int(r.Dy()),
			gdalDataType(p.DataType), C.int(pixelSize),
			C.int(stride),
		)
		if cErr != C.CE_None {
			return fmt.Errorf("gdal: Dataset.Read(%q) failed.", p.Filename)
		}
	}

	copy(data, cBuf)
	return nil
}

func (p *Dataset) ReadToCBuf(r image.Rectangle, cBuf []byte, stride int) error {
	pixelSize := SizeofPixel(p.Channels, p.DataType)

	if stride == 0 {
		stride = r.Dx() * pixelSize
	}
	if n := r.Dx() * pixelSize; stride < n {
		return fmt.Errorf("gdal: ReadToCBuf, bad stride: %d", stride)
	}

	for nBandId := 0; nBandId < p.Channels; nBandId++ {
		pBand := C.GDALGetRasterBand(p.poDataset, C.int(nBandId+1))
		cErr := C.GDALRasterIO(pBand, C.GF_Read,
			C.int(r.Min.X), C.int(r.Min.Y), C.int(r.Dx()), C.int(r.Dy()),
			unsafe.Pointer(&cBuf[nBandId*SizeofKind(p.DataType)]), C.int(r.Dx()), C.int(r.Dy()),
			gdalDataType(p.DataType), C.int(pixelSize),
			C.int(stride),
		)
		if cErr != C.CE_None {
			return fmt.Errorf("gdal: Dataset.Read(%q) failed.", p.Filename)
		}
	}
	return nil
}

func (p *Dataset) Write(r image.Rectangle, data []byte, stride int) error {
	pixelSize := SizeofPixel(p.Channels, p.DataType)

	if stride == 0 {
		stride = r.Dx() * pixelSize
	}
	if n := r.Dx() * pixelSize; stride < n {
		return fmt.Errorf("gdal: Write, bad stride: %d", stride)
	}

	if n := stride * r.Dy(); p.cBufLen < n {
		p.cBufLen = n
		if p.cBuf != nil {
			C.free(unsafe.Pointer(p.cBuf))
			p.cBuf = nil
		}
	}
	if p.cBuf == nil {
		p.cBuf = (*C.uint8_t)(C.malloc(C.size_t(p.cBufLen)))
	}

	data = data[:r.Dy()*stride]
	cBuf := ((*[1 << 30]byte)(unsafe.Pointer(p.cBuf)))[0:len(data):len(data)]
	copy(cBuf, data)

	for nBandId := 0; nBandId < p.Channels; nBandId++ {
		pBand := C.GDALGetRasterBand(p.poDataset, C.int(nBandId+1))
		cErr := C.GDALRasterIO(pBand, C.GF_Write,
			C.int(r.Min.X), C.int(r.Min.Y), C.int(r.Dx()), C.int(r.Dy()),
			unsafe.Pointer(&cBuf[nBandId*SizeofKind(p.DataType)]), C.int(r.Dx()), C.int(r.Dy()),
			gdalDataType(p.DataType), C.int(pixelSize),
			C.int(stride),
		)
		if cErr != C.CE_None {
			return fmt.Errorf("gdal: Dataset.Write(%q) failed.", p.Filename)
		}
	}

	return nil
}

func (p *Dataset) WriteFromCBuf(r image.Rectangle, cBuf []byte, stride int) error {
	pixelSize := SizeofPixel(p.Channels, p.DataType)

	if stride == 0 {
		stride = r.Dx() * pixelSize
	}
	if n := r.Dx() * pixelSize; stride < n {
		return fmt.Errorf("gdal: WriteFromCBuf, bad stride: %d", stride)
	}

	for nBandId := 0; nBandId < p.Channels; nBandId++ {
		pBand := C.GDALGetRasterBand(p.poDataset, C.int(nBandId+1))
		cErr := C.GDALRasterIO(pBand, C.GF_Write,
			C.int(r.Min.X), C.int(r.Min.Y), C.int(r.Dx()), C.int(r.Dy()),
			unsafe.Pointer(&cBuf[nBandId*SizeofKind(p.DataType)]), C.int(r.Dx()), C.int(r.Dy()),
			gdalDataType(p.DataType), C.int(pixelSize),
			C.int(stride),
		)
		if cErr != C.CE_None {
			return fmt.Errorf("gdal: Dataset.Write(%q) failed.", p.Filename)
		}
	}
	return nil
}
