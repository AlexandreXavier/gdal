// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

//#include <stdlib.h>
import "C"
import (
	"errors"
	"io"
	"reflect"
	"runtime"
	"unsafe"
)

var (
	_ CBuffer = (*cBuffer)(nil)
)

type CBuffer interface {
	CBufMagic() string
	CanResize() bool
	Resize(size int) error
	CData() []byte
	Own(d []byte) bool
	io.Closer
}

type cBuffer struct {
	*innerCBuffer
}

func NewCBuffer(size int, dontResize ...bool) CBuffer {
	return newCBuffer(size, dontResize...)
}

func newCBuffer(size int, dontResize ...bool) *cBuffer {
	p := &cBuffer{
		innerCBuffer: newInnerCBuffer(size, dontResize...),
	}
	runtime.SetFinalizer(p.innerCBuffer, (*innerCBuffer).Close)
	return p
}

func newCBufferFrom(cptr unsafe.Pointer, size int, dontResize ...bool) *cBuffer {
	p := &cBuffer{
		innerCBuffer: newInnerCBufferFrom(cptr, size, dontResize...),
	}
	runtime.SetFinalizer(p.innerCBuffer, (*innerCBuffer).Close)
	return p
}

func (p *cBuffer) Close() error {
	var err error
	if p.innerCBuffer != nil {
		err = p.innerCBuffer.Close()
	}
	*p = cBuffer{}
	return err
}

type innerCBuffer struct {
	dontResize bool
	cptr       unsafe.Pointer
	data       []byte

	malloc func(size C.size_t) unsafe.Pointer
	free   func(p unsafe.Pointer)
}

func newInnerCBufferFrom(cptr unsafe.Pointer, size int, dontResize ...bool) *innerCBuffer {
	p := new(innerCBuffer)
	p.malloc = func(size C.size_t) unsafe.Pointer { return C.malloc(size) }
	p.free = func(p unsafe.Pointer) { C.free(p) }
	if cptr != nil && size > 0 {
		p.cptr = cptr
		p.data = (*[1 << 30]byte)(p.cptr)[0:size:size]
	}
	if len(dontResize) > 0 {
		p.dontResize = dontResize[0]
	}
	return p
}

func newInnerCBuffer(size int, dontResize ...bool) *innerCBuffer {
	p := new(innerCBuffer)
	p.malloc = func(size C.size_t) unsafe.Pointer { return C.malloc(size) }
	p.free = func(p unsafe.Pointer) { C.free(p) }
	if size > 0 {
		p.cptr = p.malloc(C.size_t(size))
		p.data = (*[1 << 30]byte)(p.cptr)[0:size:size]
	}
	if len(dontResize) > 0 {
		p.dontResize = dontResize[0]
	}
	return p
}

func (p *innerCBuffer) CBufMagic() string {
	return "CBufMagic"
}

func (p *innerCBuffer) Close() error {
	if p.cptr != nil {
		p.free(p.cptr)
	}
	*p = innerCBuffer{}

	// no need for a finalizer anymore
	runtime.SetFinalizer(p, nil)
	return nil
}

func (p *innerCBuffer) CanResize() bool {
	return !p.dontResize
}

func (p *innerCBuffer) Resize(size int) error {
	if size < 0 {
		return errors.New("gdal: cBuffer.Resize, bad size!")
	}
	if p.dontResize {
		return errors.New("gdal: cBuffer.Resize, donot resize!")
	}
	if n := len(p.data); n > 0 && n != size {
		p.free(p.cptr)
		p.cptr = nil
		p.data = nil
	}
	if size > 0 {
		p.cptr = p.malloc(C.size_t(size))
		p.data = (*[1 << 30]byte)(p.cptr)[0:size:size]
	}
	return nil
}

func (p *innerCBuffer) CData() []byte {
	return p.data
}

func (p *innerCBuffer) Own(d []byte) bool {
	if cap(d) == 0 || p.cptr == nil {
		return false
	}
	min := uintptr(p.cptr)
	max := uintptr(p.cptr) + uintptr(len(p.data)-1)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	if x := hdr.Data; x < min || x > max {
		return false
	}
	if x := hdr.Data + uintptr(hdr.Cap-1); x < min || x > max {
		return false
	}
	return true
}
