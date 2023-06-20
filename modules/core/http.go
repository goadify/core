package core

import (
	"context"
	"github.com/goadify/goadify/modules/core/hydrator"
	"github.com/goadify/goadify/modules/core/models"
	"github.com/goadify/openapi/core/go/gen"
)

type httpHandler struct {
	modules []models.Module
}

func (h *httpHandler) ModulesGet(_ context.Context) ([]gen.Module, error) {
	return hydrator.Modules(h.modules), nil
}

func (h *httpHandler) NewError(_ context.Context, err error) *gen.ErrorStatusCode {
	return hydrator.Error(err)
}

func newHttpHandler(modules []models.Module) *httpHandler {
	return &httpHandler{
		modules: modules,
	}
}
