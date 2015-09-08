// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <gdal.h>
//#include <stdint.h>
//#include <stdlib.h>
import "C"
import (
	"fmt"
	"image"
	"log"
	"reflect"
	"strings"
	"unsafe"
)

type Access int

const (
	GA_ReadOnly Access = iota
	GA_Update
)

type ResampleType int

const (
	ResampleType_Nil            ResampleType = iota // "NONE"
	ResampleType_Nearest                            // "NEAREST"
	ResampleType_Gauss                              // "GAUSS"
	ResampleType_Cubic                              // "CUBIC"
	ResampleType_Average                            // "AVERAGE"
	ResampleType_Mode                               // "MODE"
	ResampleType_AverageMagpase                     // "AVERAGE_MAGPHASE"
)

func NewResampleType(name string) ResampleType {
	switch strings.ToUpper(name) {
	case "NONE":
		return ResampleType_Nil
	case "NEAREST":
		return ResampleType_Nearest
	case "GAUSS":
		return ResampleType_Gauss
	case "CUBIC":
		return ResampleType_Cubic
	case "AVERAGE":
		return ResampleType_Average
	case "MODE":
		return ResampleType_Mode
	case "AVERAGE_MAGPHASE":
		return ResampleType_AverageMagpase
	}
	return ResampleType_Nil
}

func (p ResampleType) Name() string {
	switch p {
	case ResampleType_Nil:
		return "NONE"
	case ResampleType_Nearest:
		return "NEAREST"
	case ResampleType_Gauss:
		return "GAUSS"
	case ResampleType_Cubic:
		return "CUBIC"
	case ResampleType_Average:
		return "AVERAGE"
	case ResampleType_Mode:
		return "MODE"
	case ResampleType_AverageMagpase:
		return "AVERAGE_MAGPHASE"
	}
	return "NONE"
}

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

func OpenDataset(filename string, flag Access) (p *Dataset, err error) {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	p = new(Dataset)
	p.Opt = new(Options)

	switch flag {
	case GA_ReadOnly:
		p.poDataset = C.GDALOpen(cname, C.GA_ReadOnly)
	case GA_Update:
		p.poDataset = C.GDALOpen(cname, C.GA_Update)
	default:
		err = fmt.Errorf("gdal: OpenImage(%q), unknown flag(%d).", filename, int(flag))
		return
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

func OpenDatasetWithOverviews(filename string, resampleType ResampleType) (p *Dataset, err error) {
	p, err = OpenDataset(filename, GA_ReadOnly)
	if err != nil {
		return nil, err
	}
	p.BuildOverviewsIfNotExists(resampleType)
	return p, nil
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
		optsList = append(optsList, k+"="+v)
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
		log.Printf("gdal: GDALSetProjection(%q, %s) failed!\n", filename, cProjName)
	}
	if C.GDALSetGeoTransform(p.poDataset, &padfTransform[0]) != C.CE_None {
		log.Printf("gdal: GDALSetGeoTransform(%q, %v) failed!\n", filename, padfTransform)
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
	return p.readWithSize(r, r.Dx(), r.Dy(), data, stride)
}

func (p *Dataset) ReadImage(r image.Rectangle) (m *MemPImage, err error) {
	cbuf := NewCBuffer(r.Dx() * r.Dy() * p.Channels * SizeofKind(p.DataType))
	defer cbuf.Close()

	if err = p.readWithSize(r, r.Dx(), r.Dy(), cbuf.CData(), 0); err != nil {
		return nil, err
	}
	m = &MemPImage{
		XMemPMagic: MemPMagic,
		XRect:      r,
		XStride:    r.Dx() * p.Channels * SizeofKind(p.DataType),
		XChannels:  p.Channels,
		XDataType:  p.DataType,
		XPix:       append([]byte{}, cbuf.CData()...),
	}
	return
}

func (p *Dataset) ReadImageWithSize(r image.Rectangle, size image.Point) (m *MemPImage, err error) {
	cbuf := NewCBuffer(size.X * size.Y * p.Channels * SizeofKind(p.DataType))
	defer cbuf.Close()

	if err = p.readWithSize(r, size.X, size.Y, cbuf.CData(), 0); err != nil {
		return nil, err
	}
	m = &MemPImage{
		XMemPMagic: MemPMagic,
		XRect:      image.Rect(0, 0, size.X, size.Y),
		XStride:    size.X * p.Channels * SizeofKind(p.DataType),
		XChannels:  p.Channels,
		XDataType:  p.DataType,
		XPix:       append([]byte{}, cbuf.CData()...),
	}
	return
}

func (p *Dataset) readWithSize(r image.Rectangle, nBufXSize, nBufYSize int, data []byte, stride int) error {
	pixelSize := SizeofPixel(p.Channels, p.DataType)

	if stride == 0 {
		stride = nBufXSize * pixelSize
	}
	if n := nBufXSize * pixelSize; stride < n {
		return fmt.Errorf("gdal: Dataset(%q).read, bad stride: %d", p.Filename, stride)
	}

	if n := stride * nBufYSize; p.cBufLen < n {
		p.cBufLen = n
		if p.cBuf != nil {
			C.free(unsafe.Pointer(p.cBuf))
			p.cBuf = nil
		}
	}
	if p.cBuf == nil {
		p.cBuf = (*C.uint8_t)(C.malloc(C.size_t(p.cBufLen)))
	}

	data = data[:nBufYSize*stride]
	cBuf := ((*[1 << 30]byte)(unsafe.Pointer(p.cBuf)))[0:len(data):len(data)]

	for nBandId := 0; nBandId < p.Channels; nBandId++ {
		pBand := C.GDALGetRasterBand(p.poDataset, C.int(nBandId+1))
		cErr := C.GDALRasterIO(pBand, C.GF_Read,
			C.int(r.Min.X), C.int(r.Min.Y), C.int(r.Dx()), C.int(r.Dy()),
			unsafe.Pointer(&cBuf[nBandId*SizeofKind(p.DataType)]), C.int(nBufXSize), C.int(nBufYSize),
			gdalDataType(p.DataType), C.int(pixelSize),
			C.int(stride),
		)
		if cErr != C.CE_None {
			return fmt.Errorf("gdal: Dataset(%q).read failed.", p.Filename)
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
		return fmt.Errorf("gdal: Dataset(%q).ReadToCBuf, bad stride: %d", p.Filename, stride)
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
			return fmt.Errorf("gdal: Dataset(%q).Read failed.", p.Filename)
		}
	}
	return nil
}

func (p *Dataset) Write(r image.Rectangle, data []byte, stride int) error {
	return p.write(r, data, stride)
}

func (p *Dataset) WriteImage(r image.Rectangle, src image.Image) error {
	m, ok := AsMemPImage(src)
	if !ok {
		m = NewMemPImageFrom(src)
	}
	r = r.Intersect(m.Bounds())
	return p.write(r, m.XPix, m.XStride)
}

func (p *Dataset) write(r image.Rectangle, data []byte, stride int) error {
	pixelSize := SizeofPixel(p.Channels, p.DataType)

	if stride == 0 {
		stride = r.Dx() * pixelSize
	}
	if n := r.Dx() * pixelSize; stride < n {
		return fmt.Errorf("gdal: Dataset(%q).writeLevel, bad stride: %d", p.Filename, stride)
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
			return fmt.Errorf("gdal: Dataset(%q).writeLevel failed.", p.Filename)
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
		return fmt.Errorf("gdal: Dataset(%q).WriteFromCBuf, bad stride: %d", p.Filename, stride)
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
			return fmt.Errorf("gdal: Dataset(%q).WriteFromCBuf failed.", p.Filename)
		}
	}
	return nil
}

func (p *Dataset) BuildOverviewsIfNotExists(resampleType ResampleType) error {
	if p.Width <= 256 && p.Height <= 256 {
		return nil
	}
	pBand := C.GDALGetRasterBand(p.poDataset, 1)
	if C.GDALGetOverviewCount(pBand) > 0 {
		return nil
	}
	if overviewList := p.getOverviewList(); len(overviewList) > 0 {
		if err := p.buildOverviews(resampleType, overviewList); err != nil {
			return err
		} else {
			return nil
		}
	}
	return nil
}

func (p *Dataset) BuildOverviews(resampleType ResampleType) error {
	if p.Width <= 256 && p.Height <= 256 {
		return nil
	}
	if overviewList := p.getOverviewList(); len(overviewList) > 0 {
		if err := p.buildOverviews(resampleType, overviewList); err != nil {
			return err
		} else {
			return nil
		}
	}
	return nil
}

func (p *Dataset) buildOverviews(resampleType ResampleType, overviewList []int) error {
	if len(overviewList) == 0 {
		return nil
	}

	pszResampling := C.CString(resampleType.Name())
	defer C.free(unsafe.Pointer(pszResampling))

	cptr := C.malloc(C.size_t(len(overviewList) * 4))
	defer C.free(cptr)

	nOverviews := len(overviewList)
	panOverviewList := (*[1 << 30]C.int)(cptr)[:nOverviews:nOverviews]
	for i := 0; i < len(panOverviewList); i++ {
		panOverviewList[i] = C.int(overviewList[i])
	}

	cErr := C.GDALBuildOverviews(p.poDataset, pszResampling,
		C.int(nOverviews), &panOverviewList[0],
		0, nil,
		nil, nil,
	)
	if cErr != C.CE_None {
		return fmt.Errorf("gdal: Dataset(%q).buildOverviews failed.", p.Filename)
	}
	return nil
}

// []int{2, 4, 8, ...}
func (p *Dataset) getOverviewList() []int {
	const tileSize = 256

	maxImageSize := p.Width
	if maxImageSize < p.Height {
		maxImageSize = p.Height
	}
	if maxImageSize <= tileSize {
		return nil
	}

	anOverviewList := make([]int, 30)
	for i := 0; i < len(anOverviewList); i++ {
		anOverviewList[i] = 1 << uint8(i+1)
		if x := (tileSize << uint8(i+1)); x >= maxImageSize {
			anOverviewList = anOverviewList[:i+1]
			break
		}
	}
	return anOverviewList
}

func (p *Dataset) Flush() error {
	C.GDALFlushCache(p.poDataset)
	return nil
}
