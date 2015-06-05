// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"fmt"
	"reflect"
	"unsafe"
)

// If returns trueVal if condition is tue, otherwise return falseVal.
//
//	s := If(true, "true", "false").(string)
//
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// ByteSlice convert a normal slice to byte slice.
//
// Convert []X to []byte:
//
//	x := make([]X, xLen)
//	y := ByteSlice(x)
//
// or
//
//	x := make([]X, xLen)
//	y := ((*[1 << 30]byte)(unsafe.Pointer(&x[0])))[:yLen:yLen]
//
func ByteSlice(slice interface{}) (data []byte) {
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
// Convert []X to []Y:
//	x := make([]X, xLen)
//	y := Slice(x, reflect.TypeOf([]Y(nil))).([]Y)
//
// or
//
//	x := make([]X, xLen)
//	y := ((*[1 << 30]Y)(unsafe.Pointer(&x[0])))[:yLen]
//
func Slice(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("gdal: Slice called with non-slice value of type %T", slice))
	}
	if newSliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("gdal: Slice called with non-slice type of type %T", newSliceType))
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}
