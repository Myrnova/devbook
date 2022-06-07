package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"` //omitempty se o campo estiver vazio ao converter para json ele vai retirar o campo ao inves de passar o valor padrao
	Nome     string    `json:"nome,omitempty"`
	Email    string    `json:"email,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("nome é obrigatório")
	}
	if usuario.Email == "" {
		return errors.New("email é obrigatório")
	}
	if usuario.Nick == "" {
		return errors.New("nick é obrigatório")
	}
	if usuario.Senha == "" {
		return errors.New("senha é obrigatório")
	}
	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

}
