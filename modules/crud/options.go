package crud

type Option func(*Module)

func WithEntity(entityName string, repository Repository) Option {
	return func(module *Module) {
		module.repositories = append(module.repositories, repository)

		if module.entityRepositories == nil {
			module.entityRepositories = make(map[string]Repository)
		}
		module.entityRepositories[entityName] = repository
	}
}
