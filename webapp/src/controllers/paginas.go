package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//CarregarTelaDeLogin vai renderizar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

//CarregarTelaDeCadastroDeUsuario vai rendezer a tela de cadastro de usuario
func CarregarTelaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro-usuario.html", nil)
}
