// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <stdint.h>
//#include <stdlib.h>
import "C"
import "unsafe"

type _CBuffer struct {
	cptr      *C.uint8_t
	cptrSlice []byte
}

func newCBuffer(capacity int) *_CBuffer {
	p := new(_CBuffer)
	p.cptr = (*C.uint8_t)(C.malloc(C.size_t(capacity)))
	p.cptrSlice = ((*[1 << 30]byte)(unsafe.Pointer(p.cptr)))[0:capacity:capacity]
	return nil
}

func (p *_CBuffer) Delete() {
	if p != nil && p.cptr != nil {
		C.free(unsafe.Pointer(p.cptr))
	}
	if p != nil {
		*p = _CBuffer{}
	}
}

func (p *_CBuffer) Buffer() []byte {
	return p.cptrSlice
}
