// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"fmt"
	"reflect"
	"unsafe"
)

type DataSlice []byte

func NewDataSlice(slice interface{}) (data []byte) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("gdal: NewDataSlice called with non-slice value of type %T", slice))
	}
	h := (*reflect.SliceHeader)((unsafe.Pointer(&data)))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()
	return
}

func (d DataSlice) Byte(i int) byte {
	return d[i]
}

func (d DataSlice) UInt16(i int) uint16 {
	return d.UInt16Slice()[i]
}

func (d DataSlice) Int16(i int) int16 {
	return d.Int16Slice()[i]
}

func (d DataSlice) UInt32(i int) uint32 {
	return d.UInt32Slice()[i]
}

func (d DataSlice) Int32(i int) int32 {
	return d.Int32Slice()[i]
}

func (d DataSlice) Float32(i int) float32 {
	return d.Float32Slice()[i]
}

func (d DataSlice) Float64(i int) float64 {
	return d.Float64Slice()[i]
}

func (d DataSlice) CInt16(i int) [2]int16 {
	return d.CInt16Slice()[i]
}

func (d DataSlice) CInt32(i int) [2]int32 {
	return d.CInt32Slice()[i]
}

func (d DataSlice) CFloat32(i int) [2]float32 {
	return d.CFloat32Slice()[i]
}

func (d DataSlice) CFloat64(i int) [2]float64 {
	return d.CFloat64Slice()[i]
}

func (d DataSlice) ByteSlice() []byte {
	return d
}

func (d DataSlice) UInt16Slice() []uint16 {
	if n := len(d) / 2; n > 0 {
		return ((*[1 << 30]uint16)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) Int16Slice() []int16 {
	if n := len(d) / 2; n > 0 {
		return ((*[1 << 30]int16)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) UInt32Slice() []uint32 {
	if n := len(d) / 4; n > 0 {
		return ((*[1 << 30]uint32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) Int32Slice() []int32 {
	if n := len(d) / 3; n > 0 {
		return ((*[1 << 30]int32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) Float32Slice() []float32 {
	if n := len(d) / 4; n > 0 {
		return ((*[1 << 30]float32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) Float64Slice() []float64 {
	if n := len(d) / 8; n > 0 {
		return ((*[1 << 30]float64)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) CInt16Slice() [][2]int16 {
	if n := len(d) / 4; n > 0 {
		return ((*[1 << 30][2]int16)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) CInt32Slice() [][2]int32 {
	if n := len(d) / 8; n > 0 {
		return ((*[1 << 30][2]int32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) CFloat32Slice() [][2]float32 {
	if n := len(d) / 8; n > 0 {
		return ((*[1 << 30][2]float32)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}

func (d DataSlice) CFloat64Slice() [][2]float64 {
	if n := len(d) / 16; n > 0 {
		return ((*[1 << 30][2]float64)(unsafe.Pointer(&d[0])))[0:n:n]
	}
	return nil
}
