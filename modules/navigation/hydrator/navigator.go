package hydrator

import (
	"github.com/goadify/goadify/modules/navigation/models"
	"github.com/goadify/openapi/navigation/go/gen"
)

type NavigatorHydrator struct {
	links      []*models.Link
	groups     []*models.Group
	groupLinks map[int64][]*models.Link
}

func (h *NavigatorHydrator) Hydrate() (res gen.Navigators) {
	gls := make([]oneOfGroupLink, len(h.links)+len(h.groups))

	for i, link := range h.links {
		gls[i] = oneOfGroupLink{Link: link}
	}

	for i, group := range h.groups {
		gls[i+len(h.links)] = oneOfGroupLink{Group: group}
	}

	for _, gl := range gls {
		item := gl.hydrate(h.groupLinks)
		if item == nil {
			continue
		}

		res = append(res, *item)
	}

	return
}

func NewNavigatorHydrator(links []*models.Link, groups []*models.Group, groupLinks map[int64][]*models.Link) *NavigatorHydrator {
	return &NavigatorHydrator{
		links:      links,
		groups:     groups,
		groupLinks: groupLinks,
	}
}

type oneOfGroupLink struct {
	Group *models.Group
	Link  *models.Link
}

func (gl *oneOfGroupLink) priority() int64 {
	if gl.Link != nil {
		return gl.Link.Priority
	} else if gl.Group != nil {
		return gl.Group.Priority
	}

	return -1
}

func (gl *oneOfGroupLink) hydrate(groupLinks map[int64][]*models.Link) *gen.NavigatorsItem {
	if gl.Link != nil {
		return &gen.NavigatorsItem{
			Type: gen.LinkNavigatorsItem,
			Link: Link(gl.Link),
		}
	} else if gl.Group != nil {
		links, ok := groupLinks[gl.Group.Identifier]
		if !ok || len(links) == 0 {
			return nil
		}

		return &gen.NavigatorsItem{
			Type:  gen.GroupNavigatorsItem,
			Group: Group(gl.Group, Links(links)),
		}
	}

	return nil
}
