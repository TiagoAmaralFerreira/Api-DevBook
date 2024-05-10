package routes

import "net/http"

//Routes representa todas as rotas da API
type Routes struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}
