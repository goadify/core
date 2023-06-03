package crud

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/goadify/modules/crud/models"
	"github.com/goadify/goadify/modules/crud/structs"
	"github.com/pkg/errors"
)

type entityMaster struct {
	logger interfaces.Logger

	entities           []Entity
	repositories       []Repository
	entityRepositories map[string]Repository

	entityMappingsMap map[string]models.EntityMapping
	entityMappings    []models.EntityMapping
}

func (em *entityMaster) EntityMappings() []models.EntityMapping {
	if em.entityMappings != nil {
		return em.entityMappings
	}

	em.entityMappings = make([]models.EntityMapping, len(em.entityMappingsMap))

	ind := 0
	for _, entityMapping := range em.entityMappingsMap {
		em.entityMappings[ind] = entityMapping
		ind++
	}

	return em.entityMappings
}

func (em *entityMaster) Repository(entityName string) (Repository, bool) {
	r, ok := em.entityRepositories[entityName]
	return r, ok
}

func (em *entityMaster) buildEntityMappings() {
	em.entityMappingsMap = make(map[string]models.EntityMapping)

	for _, entity := range em.entities {
		entityMapping, errs := structs.EntityToEntityMapping(entity)

		for _, err := range errs {
			em.logger.Warn(errors.Wrapf(err, "error caught while building entity mappings (%s)", entity.EntityName()))
		}

		em.entityMappingsMap[entityMapping.Name] = entityMapping
	}
}

func newEntityMaster(entities []Entity, repositories []Repository, entityRepositories map[string]Repository, logger interfaces.Logger) *entityMaster {
	em := &entityMaster{
		entities:           entities,
		repositories:       repositories,
		entityRepositories: entityRepositories,
		logger:             logger,
	}

	em.buildEntityMappings()

	return em
}
