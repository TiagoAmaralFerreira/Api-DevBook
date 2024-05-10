package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes representa todas as rotas da API
type Routes struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar todas as rotas dentro do Router
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := routesUsers

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Funcao).Methods(route.Metodo)
	}

	return r
}
