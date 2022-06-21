package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escreve informções da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}

}

// Autenticar verifica a existência de cookies
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		next(w, r)
	}

}
