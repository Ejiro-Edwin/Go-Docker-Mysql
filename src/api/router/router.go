package router

import (
	"github.com/ejiro-edwin/Blog/src/api/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router  {
	r := mux.NewRouter().StrictSlash(true)
    return routes.SetupRoutes(r)
}