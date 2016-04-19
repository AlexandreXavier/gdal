// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tfw_test

import (
	"testing"

	tfw "."
)

func TestTFW_Load(t *testing.T) {
	p, err := tfw.Load("./testdata/simple.tfw")
	if err != nil {
		t.Fatal(err)
	}

	var q = &tfw.TFW{
		OriginX: 8021.1234,
		OriginY: 10651.5678,
		StepX:   0.2,
		StepY:   -0.2,
	}
	if !tTfEqual(p, q, 0.001) {
		t.Fatalf("expect = %v, got = %v", q, p)
	}
}

func TestTFW_LoadString(t *testing.T) {
	p, err := tfw.LoadString(`
0.2000000000
-0.0000000000
0.0000000000
-0.2000000000
8021.1234
10651.5678
`)
	if err != nil {
		t.Fatal(err)
	}

	var q = &tfw.TFW{
		OriginX: 8021.1234,
		OriginY: 10651.5678,
		StepX:   0.2,
		StepY:   -0.2,
	}
	if !tTfEqual(p, q, 0.001) {
		t.Fatalf("expect = %v, got = %v", q, p)
	}
}

func tTfEqual(p, q *tfw.TFW, maxDiff float64) bool {
	if tAbsF64(p.OriginX, q.OriginX) > maxDiff {
		return false
	}
	if tAbsF64(p.OriginY, q.OriginY) > maxDiff {
		return false
	}
	if tAbsF64(p.StepX, q.StepX) > maxDiff {
		return false
	}
	if tAbsF64(p.StepY, q.StepY) > maxDiff {
		return false
	}
	return true
}

func tAbsF64(a, b float64) float64 {
	if a >= b {
		return a - b
	}
	return b - a
}
