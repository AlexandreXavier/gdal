// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <gdal.h>
//#include <stdlib.h>
import "C"
import (
	"fmt"
	"sync"
	"unsafe"
)

var (
	vsi_file_mutex sync.Mutex
	vsi_file_id    int32
	vsi_file_map   = make(map[string]*_VsiFileInfo)
)

type _VsiFileInfo struct {
	cbuf CBuffer
}

func VSITempName() (filename string) {
	vsi_file_mutex.Lock()
	defer vsi_file_mutex.Unlock()
	return _VSITempName()
}

func _VSITempName() (filename string) {
	filename = fmt.Sprintf("/vsimem/gdal.VSITempName.%08d.tmp", vsi_file_id)
	vsi_file_id++
	return
}

func VSIFileFromMemBuffer(cbuf CBuffer) (filename string, err error) {
	vsi_file_mutex.Lock()
	defer vsi_file_mutex.Unlock()

	if cbuf.CData() == nil {
		return "", fmt.Errorf("gdal: VSIFileFromMemBuffer, cbuf is zero.")
	}

	filename = _VSITempName()
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	cdata := cbuf.CData()
	f := C.VSIFileFromMemBuffer(cname,
		(*C.GByte)(unsafe.Pointer(&cdata[0])), C.vsi_l_offset(len(cdata)),
		C.FALSE,
	)
	C.VSIFClose(f)

	vsi_file_map[filename] = &_VsiFileInfo{
		cbuf: cbuf,
	}
	return filename, nil
}

func VSIGetMemFileBuffer(filename string) (CBuffer, error) {
	vsi_file_mutex.Lock()
	defer vsi_file_mutex.Unlock()

	if info, ok := vsi_file_map[filename]; ok {
		if info.cbuf != nil {
			p := unsafe.Pointer(&info.cbuf.CData()[0])
			size := len(info.cbuf.CData())

			cbuf := newCBufferFrom(p, size, true)
			cbuf.innerCBuffer.free = func(p unsafe.Pointer) {}

			return cbuf, nil
		}
	}

	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	var length C.vsi_l_offset
	p := unsafe.Pointer(C.VSIGetMemFileBuffer(cname, &length, C.FALSE))
	cbuf := newCBufferFrom(p, int(length), true)
	cbuf.innerCBuffer.free = func(p unsafe.Pointer) {}
	return cbuf, nil
}

func VSIUnlink(filename string) error {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	if C.VSIUnlink(cname) != 0 {
		return fmt.Errorf("gdal: VSIUnlink %s failed!", filename)
	}
	delete(vsi_file_map, filename)
	return nil
}
