package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPaginaPrincipal = Rota{
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarTelaPrincipal,
	RequerAutenticacao: true,
}
