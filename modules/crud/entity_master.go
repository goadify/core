package crud

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/goadify/modules/crud/models"
	"github.com/goadify/goadify/modules/crud/structs"
	"github.com/goadify/goadify/modules/navigation"
)

type entityMaster struct {
	logger   interfaces.Logger
	entities []Entity

	entityMappings     []models.EntityMapping
	entityMappingsMap  map[string]models.EntityMapping
	entityRepositories map[string]Repository
}

func (em *entityMaster) prepare() {
	em.buildEntityMappings()
}

func (em *entityMaster) buildEntityMappings() {
	em.entityMappingsMap = make(map[string]models.EntityMapping)

	for _, entity := range em.entities {
		entityMapping, errs := structs.EntityToEntityMapping(entity.Repository.NewModel(), entity.Name)
		if len(errs) > 0 {
			for _, err := range errs {
				em.logger.Error(err)
			}
		}

		em.entityMappingsMap[entityMapping.Name] = entityMapping
	}

	em.entityMappings = make([]models.EntityMapping, len(em.entityMappingsMap))

	ind := 0
	for _, entityMapping := range em.entityMappingsMap {
		em.entityMappings[ind] = entityMapping
		ind++
	}
}

func (em *entityMaster) EntityMappings() []models.EntityMapping {
	return em.entityMappings
}

func (em *entityMaster) Repository(entityName string) (Repository, bool) {
	r, ok := em.entityRepositories[entityName]
	return r, ok
}

func (em *entityMaster) Links() (res []*navigation.Link) {
	for _, entity := range em.entities {
		entity.Link.Identifier = entity.Slug
		if entity.Link != nil && entity.Group == nil {
			res = append(res, entity.Link)
		}
	}

	return
}

func (em *entityMaster) GroupLinks() map[*navigation.Group][]*navigation.Link {
	res := make(map[*navigation.Group][]*navigation.Link)
	for _, entity := range em.entities {
		if entity.Group != nil && entity.Link != nil {
			if links, ok := res[entity.Group]; ok {
				links = append(links, entity.Link)
			} else {
				res[entity.Group] = []*navigation.Link{entity.Link}
			}
		}
	}

	return res
}

func newEntityMaster(
	logger interfaces.Logger,
	entities []Entity,
) *entityMaster {
	em := &entityMaster{
		logger:   logger,
		entities: entities,
	}

	em.prepare()

	return em
}
