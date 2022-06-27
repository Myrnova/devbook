package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
)

type repositorioDeUsuarios struct {
	db *sql.DB
}

// NovoReposotorioDeUsuarios cria um novo repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *repositorioDeUsuarios { //recebe um banco e retorna um ponteiro de usuarios
	return &repositorioDeUsuarios{db} //criamos uma instancia de um struct e passamos o db para ela

}

func (repositorio repositorioDeUsuarios) CriarUsuario(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, email, nick, senha) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Email, usuario.Nick, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// BuscarUsuarios traz todos os usuarios um filtro de nome ou nick
func (repositorio repositorioDeUsuarios) BuscarUsuarios(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOUnick% utiliza dois %% para fazer o escape de % para não ser usada para fazer substituição de caracteres
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criado_em from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

//BuscaPorID traz um usuario especifico do banco de dados
func (repositorio repositorioDeUsuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"select id, nome, nick, email, criado_em from usuarios where id = ?", ID,
	)

	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro := linha.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuario no bando de dados
func (repositorio repositorioDeUsuarios) AtualizarUsuario(ID uint64, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, email = ?, nick = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(
		usuario.Nome,
		usuario.Email,
		usuario.Nick,
		ID,
	); erro != nil {
		return erro
	}
	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio repositorioDeUsuarios) DeletarUsuario(ID uint64) error {

	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

//BuscarPorEmail traz o id e senha do usuário com aquele email
func (repositorio repositorioDeUsuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = ?", email,
	)

	if erro != nil {
		return models.Usuario{}, erro //models.Usuario{} é um usuario vazio
	}

	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	} else {
		return models.Usuario{}, errors.New("nenhum usuário encontrado")
	}
	return usuario, nil
}

//BuscarSenha traz a senha de um usuário pelo ID
func (repositorio repositorioDeUsuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}

//AtualizarSenha atualiza a senha do usuário
func (repositorio repositorioDeUsuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio repositorioDeUsuarios) SeguirUsuario(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)",
	) //ignore não permite que insira duas colunas iguais novamente
	if erro != nil {
		return erro
	}

	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

//PararDeSeguirUsuario permite que um usuário para de seguir o outro
func (repositorio repositorioDeUsuarios) PararDeSeguirUsuario(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? AND seguidor_id = ?",
	) //ignore não permite que insira duas colunas iguais novamente
	if erro != nil {
		return erro
	}

	defer statement.Close()
	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

//BuscarSeguidores traz todos os usuários que seguem determinado usuário
func (repositorio repositorioDeUsuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criado_em
		from usuarios u
			inner join seguidores s on u.id = s.seguidor_id 
			where s.usuario_id = ? 
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

//BuscarSeguindo traz todos os usuários que determinado usuário segue
func (repositorio repositorioDeUsuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criado_em
		from usuarios u
			inner join seguidores s on u.id = s.usuario_id 
			where s.seguidor_id = ? 
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
