package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CriarUsuario chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	nome := r.FormValue("nome")
	nick := r.FormValue("nick")
	email := r.FormValue("email")
	senha := r.FormValue("password")

	usuario, erro := json.Marshal(map[string]string{
		"nome":  nome,
		"nick":  nick,
		"email": email,
		"senha": senha,
	})

	if erro != nil {
		log.Fatal(erro)
	}
	//bytes.NewBuffer(usuario) transforma pra json
	response, erro := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))

	if erro != nil {
		log.Fatal(erro)
	}

	defer response.Body.Close() //reponse.Body precisa ser fechado

	fmt.Println(response.Body)
}
