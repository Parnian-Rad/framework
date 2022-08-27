package fiber

import (
	"git.snapp.ninja/search-and-discovery/framework/pkg/ports"
)

func (fh *FiberHttpServer) SetRouteGroups(groupName string, routes []ports.Route) {
	g := fh.app.Group("/" + groupName)
	for _, route := range routes {
		g.Add(string(route.Method), route.Path, route.Handler)
	}
}
