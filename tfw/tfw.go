// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tfw provides tfw reader.
package tfw

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type TFW struct {
	OriginX float64
	OriginY float64
	StepX   float64
	StepY   float64
}

func Load(filename string) (p *TFW, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return read(f)
}

func LoadString(s string) (p *TFW, err error) {
	return read(strings.NewReader(s))
}

func Read(r io.Reader) (p *TFW, err error) {
	return read(r)
}

func read(r io.Reader) (p *TFW, err error) {
	p = new(TFW)

	var t0, t1 float64
	n, err := fmt.Fscan(r,
		&p.StepX, &t0, &t1, &p.StepY,
		&p.OriginX, &p.OriginY,
	)
	if err != nil {
		return nil, err
	}
	if n != 6 {
		panic("unreach able!")
	}
	return
}
