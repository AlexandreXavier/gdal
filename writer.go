// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"fmt"
	"image"
)

// Encode writes the image m to w in GDAL format.
func Save(filename string, m image.Image, opt *Options) error {
	return fmt.Errorf("gdal: Save, TODO")
}

// See http://www.gdal.org/formats_list.html
var defaultDriverNameMap = map[string]string{
	".blx":    "BLX",
	".xlb":    "BLX",
	".bmp":    "BMP",
	".kap":    "BSB",
	".bt":     "BT",
	".dim":    "DIMAP",
	".dog":    "DOQ1", // "DOQ2",
	".dt0":    "DTED",
	".dt1":    "DTED",
	".dt2":    "DTED",
	"toc.xml": "ECRGTOC", // /path/to/TOC.xml
	".hdr":    "EHdr",    // "ENVI","GENBIN"
	".ers":    "ERS",
	".n1":     "ESAT",
	".gif":    "GIF",
	".grb":    "GRIB",
	".gta":    "GTA",
	".tif":    "GTiff",
	".tiff":   "GTiff",
	".img":    "HFA",
	".mpr":    "ILWIS",
	".mpl":    "ILWIS",
	".jpg":    "JPEG",
	".jpeg":   "JPEG",
	".ntf":    "NITF",
	".nsf":    "NITF",
	".grc":    "NWT_GRC",
	".tab":    "NWT_GRC",
	".png":    "PNG",
	".ppm":    "PNM",
	".pgm":    "PNM",
	".rik":    "RIK",
	".rsw":    "RMF",
	".mtw":    "RMF",
	".ter":    "TERRAGEN",
	".dem":    "USGSDEM",
	".vrt":    "VRT",
	".xpm":    "XPM",
}
