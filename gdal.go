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

func genVsiMemFilename() (filename string) {
	vsi_file_mutex.Lock()
	defer vsi_file_mutex.Unlock()
	filename = fmt.Sprintf("/vsimem/newVsiMemFilename_%08d.tmp", vsi_file_id)
	vsi_file_id++
	return
}

func newVsiMemFile(cbuf CBuffer) (filename string, err error) {
	vsi_file_mutex.Lock()
	defer vsi_file_mutex.Unlock()

	if cbuf.CData() == nil {
		return "", fmt.Errorf("gdal: newVsiMemFile, cbuf is zero.")
	}

	filename = genVsiMemFilename()
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

func getVsiMemFileData(filename string) (CBuffer, error) {
	vsi_file_mutex.Lock()
	defer vsi_file_mutex.Unlock()

	info, ok := vsi_file_map[filename]
	if !ok {
		return nil, fmt.Errorf("gdal: getVsiMemFileData %s, file not found!", filename)
	}
	if info.cbuf != nil {
		return info.cbuf, nil
	}

	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))

	var length C.vsi_l_offset
	p := unsafe.Pointer(C.VSIGetMemFileBuffer(cname, &length, C.FALSE))
	cbuf := newCBufferFrom(p, int(length), true)
	cbuf.innerCBuffer.free = func(p unsafe.Pointer) {}
	return cbuf, nil
}

func removeVsiMemFile(filename string) error {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	if C.VSIUnlink(cname) != 0 {
		return fmt.Errorf("gdal: removeVsiMemFile %s failed!", filename)
	}
	delete(vsi_file_map, filename)
	return nil
}
