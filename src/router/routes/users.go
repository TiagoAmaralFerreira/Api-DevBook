package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Routes{
	{
		Uri:                "/users",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/users",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetUsers,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/users/{userId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetUser,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/users/{userId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditUser,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/users/{userId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: false,
	},
}
