package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
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

//CarregarTelaPrincipal vai rendezer a página principal com as publicações
func CarregarTelaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s//publicacoes", config.APIUrl)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, erro)

	respostas.TratarStatusCodeDeErro(w, response)
	utils.ExecutarTemplate(w, "home.html", nil)
}
