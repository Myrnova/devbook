package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//CarregarTemplates insere os tempaltes html na váriavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/*/*.html"))
}

//ExecutarTemplate renderiza uma página html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
