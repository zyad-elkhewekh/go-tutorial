package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/zyad-elkhewekh/go-tutorial/bubTea/internals/middleware"
)

func Handler(r *chi.Mux) {
	//middleware as the name suggestes is to be used before the main function
	//add a global mw for a functionality needed by all
	r.Use(chimiddle.StripSlashes) //avoid a 404 error if slash at the end of request etc

	r.Route("/choice", func(router chi.Router) {

		router.Use(middleware.Authorize)
		router.Get("/", getChoices)
	})
}
