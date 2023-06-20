package navigation

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/goadify/modules/navigation/models"
	"github.com/goadify/openapi/navigation/go/gen"
	"net/http"
	"sort"
)

type Module struct {
	devModeEnabled bool
	modules        []interfaces.Module

	links      []*models.Link
	groups     []*models.Group
	groupLinks map[int64][]*models.Link
}

type NavigatableModule interface {
	Links() []*Link
	GroupLinks() map[*Group][]*Link
}

func (m *Module) Name() string {
	return "navigation"
}

func (m *Module) HttpPrefix() string {
	return "/navigation/v1"
}

func (m *Module) HttpHandler() (http.Handler, error) {
	hh := newHttpHandler(m.devModeEnabled, m.links, m.groups, m.groupLinks)

	return gen.NewServer(hh)
}

func convertLink(link *Link, moduleName string) *models.Link {
	if link == nil {
		return nil
	}
	return &models.Link{
		Title:      link.Title,
		Identifier: link.Identifier,
		Priority:   link.Priority,
		ModuleName: moduleName,
	}
}

func convertLinks(links []*Link, moduleName string) []*models.Link {
	var result []*models.Link
	for _, link := range links {
		convLink := convertLink(link, moduleName)
		if convLink == nil {
			continue
		}
		result = append(result, convLink)
	}

	return result
}

func convertGroup(group *Group) *models.Group {
	return &models.Group{
		Identifier: group.Identifier,
		Title:      group.Title,
		Priority:   group.Priority,
	}
}

func (m *Module) prepare() {
	m.groupLinks = make(map[int64][]*models.Link)
	for _, module := range m.modules {
		nm, ok := module.(NavigatableModule)
		if !ok {
			continue
		}

		moduleName := module.Name()

		if ls := nm.Links(); len(ls) > 0 {
			m.links = append(m.links, convertLinks(ls, moduleName)...)
		}

		for group, links := range nm.GroupLinks() {
			gID := group.Identifier

			if ls, ok := m.groupLinks[gID]; ok {

				ls = append(ls, convertLinks(links, moduleName)...)

			} else {

				m.groupLinks[gID] = convertLinks(links, moduleName)
				m.groups = append(m.groups, convertGroup(group))

			}
		}
	}

	m.sortGroupLinks()
}

func (m *Module) sortGroupLinks() {
	for _, links := range m.groupLinks {
		sort.Slice(links, func(i, j int) bool {
			return links[i].Priority < links[j].Priority
		})
	}
}

func (m *Module) Init(provider interfaces.ExtendedDependencyProvider) error {
	m.devModeEnabled = provider.DevModeEnabled()
	m.modules = provider.Modules()

	m.prepare()

	return nil
}

func New() *Module {
	return new(Module)
}
