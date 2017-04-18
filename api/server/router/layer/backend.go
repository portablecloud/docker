package layers

import (
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/layer"
)

type Backend interface {
	layerBackend
	exportBackend
}

type layerBackend interface {
	LookupLayer(name string) (layer.Layer, error)
}

type exportBackend interface {
	ExportLayer(name string, outStream io.Writer) error
}
