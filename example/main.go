package main

import (
	"github.com/goadify/goadify"
	"github.com/goadify/goadify/example/repository"
	"github.com/goadify/goadify/modules/crud"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	g := goadify.New(
		goadify.WithConfig(
			goadify.Config{IsDevMode: true},
		),
		goadify.WithModule(
			crud.NewModule(
				crud.WithEntity(
					"user",
					new(repository.UserRepository),
				),
			),
		),
		goadify.WithLogger(logrus.New()),
	)

	handler, err := g.Handler()
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe("localhost:8080", handler); err != nil {
		panic(err)
	}

}
