package main

import (
	"github.com/goadify/goadify"
	"net/http"
)

func main() {
	g := goadify.New()

	handler, err := g.Handler()
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe("localhost:8080", handler); err != nil {
		panic(err)
	}

}
