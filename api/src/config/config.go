package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	SecretKey          []byte
)

//Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT")) //strconv.Atoi = Parse.int

	if erro != nil {
		Porta = 9000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"))
}
