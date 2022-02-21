package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sptGabriel/starwars/app/gateway/api/middlewares"
	"github.com/sptGabriel/starwars/app/gateway/api/responses"
	httpSwagger "github.com/swaggo/http-swagger"
)

type planetHandler interface {
	CreatePlanet(r *http.Request) responses.Response
	DeletePlanet(r *http.Request) responses.Response
	GetPlanetByID(r *http.Request) responses.Response
	GetPlanetByName(r *http.Request) responses.Response
	ListPlanets(r *http.Request) responses.Response
}

func NewRouter(planetHandler planetHandler) http.Handler {
	router := chi.NewRouter()

	router.Route("/docs/v1/swagger", func(r chi.Router) {
		r.Get("/", httpSwagger.WrapHandler)
	})

	router.Route("/api/starwars/planets", func(r chi.Router) {
		r.Post("/", middlewares.Handle(planetHandler.CreatePlanet))
		r.Delete("/{id}", middlewares.Handle(planetHandler.DeletePlanet))
		r.Get("/", middlewares.Handle(planetHandler.ListPlanets))
		r.Get("/{id}", middlewares.Handle(planetHandler.GetPlanetByID))
		r.Get("/name/{name}", middlewares.Handle(planetHandler.GetPlanetByName))
	})

	return router
}
