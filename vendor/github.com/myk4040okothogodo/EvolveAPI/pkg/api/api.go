package api

import (
    "net/http"
    "github.com/myk4040okothogodo/EvolveAPI/pkg/db"
    "go.uber.org/zap"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    httpSwagger "github.com/swaggo/http-swagger"

    um "github.com/myk4040okothogodo/EvolveAPI/pkg/middleware"

)

var DBClient db.ClientInterface

func SetDBClient(c db.ClientInterface) {
	DBClient = c
	um.SetDBClient(DBClient)
}

// GetRouter configures a chi router and starts the http server
// @title My API
// @description This API is a fullfillment of Evolves hiring test.
// @description It also does this.
// @host example.com
// @BasePath /
func GetRouter(log *zap.Logger, dbClient db.ClientInterface) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	SetDBClient(dbClient)
	if log != nil {
		r.Use(um.SetLogger(log))
	}
	buildTree(r)

	return r
}


func buildTree(r *chi.Mux) {
	r.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/", http.StatusMovedPermanently)
	})
	r.Get("/swagger*", httpSwagger.Handler())

	r.Route("/users", func(r chi.Router) {
		r.With(um.Pagination).Get("/", ListUsers)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(um.User)
			r.Get("/", GetUser)
		})

		r.Put("/", PutUser)
	})

}
