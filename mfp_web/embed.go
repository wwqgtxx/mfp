package mfp_web

import (
	"embed"
	"io/fs"
)

//go:embed dist
var dist embed.FS

var Dist fs.FS

func init() {
	var err error
	Dist, err = fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
}
