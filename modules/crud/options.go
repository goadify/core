package crud

type Option func(*Module)

func WithEntity(entity Entity, repository Repository) Option {
	return func(module *Module) {
		module.entities = append(module.entities, entity)
		module.repositories = append(module.repositories, repository)

		module.entityRepositories[entity.EntityName()] = repository
	}
}
