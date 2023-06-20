package main

import (
	"github.com/goadify/goadify"
	"github.com/goadify/goadify/modules/crud"
	"github.com/goadify/goadify/modules/navigation"
	"net/http"
)

func main() {
	userGroup := navigation.NewGroup("User section", 0)
	g := goadify.New(
		goadify.WithModules(
			crud.New(
				crud.Entity{
					Slug:       "user",
					Name:       "User",
					Repository: new(UserRepository),
					Link:       navigation.NewLink("Profiles", 0),
					Group:      userGroup,
				},
			),
		),
	)

	srv, err := g.HttpHandler()
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe("localhost:8080", srv)
	if err != nil {
		panic(err)
	}
}
