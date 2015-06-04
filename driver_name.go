// Copyright 2015 ChaiShushan <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"path/filepath"
	"strings"
)

func getDriverName(filename string) string {
	filename = strings.ToLower(filename)
	if strings.HasSuffix(filename, "toc.xml") {
		return "ECRGTOC"
	}
	ext := filepath.Ext(filename)
	s, _ := defaultDriverNameMap[ext]
	return s
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
