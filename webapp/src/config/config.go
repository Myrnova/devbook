package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIUrl representa a URL para comunicação com a API
	APIUrl = ""
	// POrta onde a aplicação web esta rodando
	Porta = 0
	// HashKey é utilizada para autenticar o cookie
	HashKey []byte
	// BlockKey é utilizada para criptografar o cookie
	BlockKey []byte
)

// Carregar inicializa as variáveis de ambiente

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))

	if erro != nil {
		log.Fatal(erro)
	}

	APIUrl = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
