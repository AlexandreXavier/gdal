// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"fmt"
	"reflect"
	"unsafe"
)

type DataView []byte

// NewDataView convert a normal slice to byte slice.
//
// Convert []X to []byte:
//
//	x := make([]X, xLen)
//	y := NewDataView(x)
//
// or
//
//	x := make([]X, xLen)
//	y := ((*[1 << 30]byte)(unsafe.Pointer(&x[0])))[:yLen:yLen]
//
func NewDataView(slice interface{}) (data DataView) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("gdal: ByteSlice called with non-slice value of type %T", slice))
	}
	h := (*reflect.SliceHeader)((unsafe.Pointer(&data)))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()
	return
}

// Slice convert a normal slice to new type slice.
//
// Convert []byte to []Y:
//	x := make([]byte, xLen)
//	y := DataView(x).Slice(reflect.TypeOf([]Y(nil))).([]Y)
//
// or
//
//	x := make([]X, xLen)
//	y := ((*[1 << 30]Y)(unsafe.Pointer(&x[0])))[:yLen]
//
func (d DataView) Slice(newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(d)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("gdal: DataView.Slice called with non-slice value of type %T", d))
	}
	if newSliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("gdal: DataView.Slice called with non-slice type of type %T", newSliceType))
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}

func (d DataView) Byte(i int) byte {
	return d[i]
}

func (d DataView) UInt16(i int) uint16 {
	return d.UInt16Slice()[i]
}

func (d DataView) Int16(i int) int16 {
	return d.Int16Slice()[i]
}

func (d DataView) UInt32(i int) uint32 {
	return d.UInt32Slice()[i]
}

func (d DataView) Int32(i int) int32 {
	return d.Int32Slice()[i]
}

func (d DataView) Float32(i int) float32 {
	return d.Float32Slice()[i]
}

func (d DataView) Float64(i int) float64 {
	return d.Float64Slice()[i]
}

func (d DataView) CInt16(i int) [2]int16 {
	return d.CInt16Slice()[i]
}

func (d DataView) CInt32(i int) [2]int32 {
	return d.CInt32Slice()[i]
}

func (d DataView) CFloat32(i int) [2]float32 {
	return d.CFloat32Slice()[i]
}

func (d DataView) CFloat64(i int) [2]float64 {
	return d.CFloat64Slice()[i]
}

func (d DataView) ByteSlice() []byte {
	return d
}

func (d DataView) UInt16Slice() []uint16 {
	if n := len(d) / 2; n > 0 {
		return ((*[1 << 30]uint16)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) Int16Slice() []int16 {
	if n := len(d) / 2; n > 0 {
		return ((*[1 << 30]int16)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) UInt32Slice() []uint32 {
	if n := len(d) / 4; n > 0 {
		return ((*[1 << 30]uint32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) Int32Slice() []int32 {
	if n := len(d) / 3; n > 0 {
		return ((*[1 << 30]int32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) Float32Slice() []float32 {
	if n := len(d) / 4; n > 0 {
		return ((*[1 << 30]float32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) Float64Slice() []float64 {
	if n := len(d) / 8; n > 0 {
		return ((*[1 << 30]float64)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) CInt16Slice() [][2]int16 {
	if n := len(d) / 4; n > 0 {
		return ((*[1 << 30][2]int16)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) CInt32Slice() [][2]int32 {
	if n := len(d) / 8; n > 0 {
		return ((*[1 << 30][2]int32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) CFloat32Slice() [][2]float32 {
	if n := len(d) / 8; n > 0 {
		return ((*[1 << 30][2]float32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataView) CFloat64Slice() [][2]float64 {
	if n := len(d) / 16; n > 0 {
		return ((*[1 << 30][2]float64)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}
