package goadify

import (
	"github.com/goadify/goadify/interfaces"
	"github.com/goadify/goadify/modules/core"
	"github.com/goadify/goadify/modules/navigation"
	"github.com/pkg/errors"
	"net/http"
)

var (
	ErrModuleAlreadyLoaded        = errors.New("module already loaded")
	ErrPrefixNotStartingWithSlash = errors.New("module's http prefix must starts with slash")
	ErrPrefixEndsWithSlash        = errors.New("module's http prefix must NOT ends with slash")

	autoloadModules = []interfaces.ExtendedModule{
		core.New(),
		navigation.New(),
	}
)

type moduleMaster struct {
	mux *http.ServeMux

	logger         interfaces.Logger
	devModeEnabled bool

	ordinaryModules    []interfaces.OrdinaryModule
	ordinaryModulesMap map[string]interfaces.OrdinaryModule

	extendedModules    []interfaces.ExtendedModule
	extendedModulesMap map[string]interfaces.ExtendedModule
}

type dependencyProvider struct {
	logger         interfaces.Logger
	devModeEnabled bool
}

type extendedDependencyProvider struct {
	dependencyProvider
	modules []interfaces.Module
}

func (d *dependencyProvider) Logger() interfaces.Logger {
	return d.logger
}

func (d *dependencyProvider) DevModeEnabled() bool {
	return d.devModeEnabled
}

func (d *extendedDependencyProvider) Modules() []interfaces.Module {
	return d.modules
}

func (mm *moduleMaster) prepare() error {
	mm.mux = http.NewServeMux()

	mm.extendedModules = append(mm.extendedModules, autoloadModules...)

	mm.ordinaryModulesMap = make(map[string]interfaces.OrdinaryModule)

	for _, module := range mm.ordinaryModules {
		moduleName := module.Name()
		if _, ok := mm.ordinaryModulesMap[moduleName]; ok {

			return errors.Wrap(ErrModuleAlreadyLoaded, moduleName)
		}

		mm.ordinaryModulesMap[moduleName] = module

	}

	mm.extendedModulesMap = make(map[string]interfaces.ExtendedModule)

	for _, module := range mm.extendedModules {
		moduleName := module.Name()
		if _, ok := mm.extendedModulesMap[moduleName]; ok {

			return errors.Wrap(ErrModuleAlreadyLoaded, moduleName)
		}

		mm.extendedModulesMap[moduleName] = module

	}

	return nil
}

func checkHttpModule(m interfaces.Module) error {
	prefix := m.HttpPrefix()

	if prefix[0] != '/' {
		return ErrPrefixNotStartingWithSlash
	}

	if prefix[len(prefix)-1] == '/' {
		return ErrPrefixEndsWithSlash
	}

	return nil
}

func (mm *moduleMaster) registerModuleHttp(module interfaces.Module) error {
	err := checkHttpModule(module)
	if err != nil {
		return err
	}

	prefix := module.HttpPrefix()

	handler, err := module.HttpHandler()
	if err != nil {
		return errors.Wrapf(err, "can not build module %s", module.Name())
	}

	mm.mux.Handle(
		prefix+"/",
		http.StripPrefix(prefix, handler),
	)

	return nil
}

func (mm *moduleMaster) loadModules() error {
	dp := dependencyProvider{
		logger:         mm.logger,
		devModeEnabled: mm.devModeEnabled,
	}

	for _, module := range mm.ordinaryModules {
		err := module.Init(&dp)
		if err != nil {
			return errors.Wrap(err, module.Name())
		}

		if err := mm.registerModuleHttp(module); err != nil {
			return errors.Wrap(err, module.Name())
		}
	}

	var modules []interfaces.Module
	for _, module := range mm.ordinaryModulesMap {
		modules = append(modules, module)
	}

	for _, module := range mm.extendedModules {
		modules = append(modules, module)
	}

	edp := &extendedDependencyProvider{
		dependencyProvider: dp,
		modules:            modules,
	}

	for _, module := range mm.extendedModules {
		err := module.Init(edp)
		if err != nil {
			return errors.Wrap(err, module.Name())
		}

		if err := mm.registerModuleHttp(module); err != nil {
			return errors.Wrap(err, module.Name())
		}
	}

	return nil
}

func (mm *moduleMaster) HttpHandler() http.Handler {
	return mm.mux
}

func newModuleMaster(
	logger interfaces.Logger,
	devModeEnabled bool,
	ordinaryModules []interfaces.OrdinaryModule,
	extendedModules []interfaces.ExtendedModule,
) (*moduleMaster, error) {
	mm := &moduleMaster{
		logger:          logger,
		devModeEnabled:  devModeEnabled,
		ordinaryModules: ordinaryModules,
		extendedModules: extendedModules,
	}

	err := mm.prepare()
	if err != nil {
		return nil, err
	}

	err = mm.loadModules()
	if err != nil {
		return nil, err
	}

	return mm, nil
}
