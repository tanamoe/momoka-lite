package docs

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var distDir embed.FS

// DistDirFS contains the embedded dist directory files (without the "dist" prefix)
var DistDirFS, Err = fs.Sub(distDir, "dist")
