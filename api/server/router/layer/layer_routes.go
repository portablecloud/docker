package layer

import (
	"context"
	"io/ioutil"

	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/pkg/ioutils"
	"github.com/docker/docker/pkg/streamformatter"
)

func (s *layerRouter) getLayersGet(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if err := httputils.ParseForm(r); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/x-tar")

	output := ioutils.NewWriteFlusher(w)

	if err := s.backend.ExportLayer(vars["name"], output); err != nil {
		if !output.Flushed() {
			return err
		}
		sf := streamformatter.NewJSONStreamFormatter()
		output.Write(sf.FormatError(err))
	}
	return nil
}

func (s *layerRouter) getLayersByName(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	layerInspect, err := s.backend.LookupLayer(vars["name"])
	if err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, layerInspect)
}
