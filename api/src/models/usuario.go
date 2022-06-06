package models

import "time"

type Usuario struct {
	ID       uint64    `json:"id,omitempty"` //omitempty se o campo estiver vazio ao converter para json ele vai retirar o campo ao inves de passar o valor padrao
	Nome     string    `json:"nome,omitempty"`
	Email    string    `json:"email,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}
