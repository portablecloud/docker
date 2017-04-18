package layer

import (
	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/server/router"
)

type layerRouter struct {
	backend Backend
	decoder httputils.ContainerDecoder
	routes  []router.Route
}

func NewRouter(backend Backend, decoder httputils.ContainerDecoder) router.Router {
	r := &layerRouter{
		backend: backend,
		decoder: decoder,
	}
	r.initRoutes()
	return r
}

func (r *layerRouter) Routes() []router.Route {
	return r.routes
}

func (r *imageRouter) initRoutes() {
	r.routes = []router.Route{
		// GET
		router.NewGetRoute("/layers/{name:.*}/get", r.getLayersGet),
		router.NewGetRoute("/layers/{name:.*}/json", r.getLayersByName),
	}
}
