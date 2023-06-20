package navigation

import (
	"context"
	"github.com/goadify/goadify/modules/navigation/hydrator"
	"github.com/goadify/goadify/modules/navigation/models"
	"github.com/goadify/openapi/navigation/go/gen"
)

type httpHandler struct {
	devModeEnabled bool
	links          []*models.Link
	groups         []*models.Group
	groupLinks     map[int64][]*models.Link

	cachedNavigators gen.Navigators
}

func (h *httpHandler) GetNavigators(_ context.Context) (gen.Navigators, error) {
	if h.cachedNavigators != nil {
		return h.cachedNavigators, nil
	}

	hydr := hydrator.NewNavigatorHydrator(h.links, h.groups, h.groupLinks)
	h.cachedNavigators = hydr.Hydrate()

	return h.cachedNavigators, nil
}

func (h *httpHandler) NewError(_ context.Context, err error) *gen.ErrorStatusCode {
	return hydrator.Error(err, hydrator.ErrorDisplayInternalMessages(h.devModeEnabled))
}

func newHttpHandler(
	devModeEnabled bool,

	links []*models.Link,
	groups []*models.Group,
	groupLinks map[int64][]*models.Link,
) *httpHandler {
	return &httpHandler{
		devModeEnabled: devModeEnabled,

		links:      links,
		groups:     groups,
		groupLinks: groupLinks,
	}
}
