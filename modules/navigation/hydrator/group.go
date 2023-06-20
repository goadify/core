package hydrator

import (
	"github.com/goadify/goadify/modules/navigation/models"
	"github.com/goadify/openapi/navigation/go/gen"
)

func Group(g *models.Group, links []gen.Link) gen.Group {
	return gen.Group{
		Title: g.Title,
		Links: links,
	}
}
