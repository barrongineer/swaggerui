package swaggerui

import (
	"net/http"
	"github.com/elazarl/go-bindata-assetfs"
)

func FileSystem() http.FileSystem {
	return http.FileSystem(assetFS())
}

func AssetFS() *assetfs.AssetFS {
	return assetFS()
}
