// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

type GCP struct {
	Id       string  // Unique identifier, often numeric
	Info     string  // Informational message or ""
	GCPPixel float64 // Pixel (x) location of GCP on raster
	GCPLine  float64 // Line (y) location of GCP on raster
	GCPX     float64 // X position of GCP in georeferenced space
	GCPY     float64 // Y position of GCP in georeferenced space
	GCPZ     float64 // Elevation of GCP, or zero if not known
}
