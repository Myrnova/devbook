package controllers

import (
	"fmt"
	"net/http"
)

//CriarUsuario inseri um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Criar usuario")
	return

}

//BuscarUsuarios busca todos usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscar usuarios")
	return
}

//BuscarUsuario busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Buscar usuario")
	return

}

//AtualizarUsuario atualiza um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Atualizar usuario")

	return
}

//DeletarUsuario exclui um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deletar usuario")

	return

}
