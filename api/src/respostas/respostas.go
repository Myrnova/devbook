package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON retorna uma resposta em json no writer da requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil { //in case we want to use the nocontent status we can't send a body even if the data passed is nil
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

//Erro retorna um erro na função w da requisição
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	}) //criando um struc generico para passar o erro
}
