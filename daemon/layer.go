package daemon

import (
	"io"
	"time"

	"github.com/docker/docker/layer"
	"github.com/pkg/errors"
)

func (daemon *Daemon) LookupLayer(digest string) (layer.Layer, error) {
	return daemon.layerStore.Get(layer.ChainID(digest))
}

func (daemon *Daemon) ExportLayer(digest string, outStream io.Writer) error {
	layer, err := daemon.layerStore.Get(layer.ChainID(digest))
	if err != nil {
		return errors.Wrap(err, "no such layer: %s", digest)
	}

	stream, err := layer.TarStream()
	if err != nil {
		return errors.Wrap(err, "unable to create a new tar stream for layer %s", digest)
	}

	if _, err := io.Copy(outStream, stream); err != nil {
		return errors.Wrap(err, "unable to copy tar stream")
	}

	return nil
}
