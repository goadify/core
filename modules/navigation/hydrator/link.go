package hydrator

import (
	"github.com/goadify/goadify/modules/navigation/models"
	"github.com/goadify/openapi/navigation/go/gen"
)

func Link(link *models.Link) gen.Link {
	return gen.Link{
		Title:      link.Title,
		ModuleName: link.ModuleName,
		Identifier: link.Identifier,
	}
}

func Links(links []*models.Link) []gen.Link {
	res := make([]gen.Link, len(links))
	for i, link := range links {
		res[i] = Link(link)
	}

	return res
}
