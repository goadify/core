package crud

type Option func(*Module)

func WithEntity(entity Entity, repository Repository) Option {
	return func(module *Module) {
		module.entities = append(module.entities, entity)
		module.repositories = append(module.repositories, repository)

		if module.entityRepositories == nil {
			module.entityRepositories = make(map[string]Repository)
		}
		module.entityRepositories[entity.EntityName()] = repository
	}
}
