//go:build embed_frontend

package main

// This file is provided for conditional compilation. When building without
// web frontend assets, omit this file by not passing -tags embed_frontend.
// The variables it sets are declared in main.go as zero-value defaults.

import "embed"

//go:embed web/default/dist
var _buildFS embed.FS

//go:embed web/default/dist/index.html
var _indexPage []byte

//go:embed web/classic/dist
var _classicBuildFS embed.FS

//go:embed web/classic/dist/index.html
var _classicIndexPage []byte

func init() {
	buildFS = _buildFS
	indexPage = _indexPage
	classicBuildFS = _classicBuildFS
	classicIndexPage = _classicIndexPage
}
