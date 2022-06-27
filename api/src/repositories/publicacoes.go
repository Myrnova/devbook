package repositories

import (
	"api/src/models"
	"database/sql"
)

//RepositorioPublicacoes representa um repositorio de publicacoes
type RepositorioPublicacoes struct {
	db *sql.DB
}

//NovoRepositorioDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *RepositorioPublicacoes {
	return &RepositorioPublicacoes{db}
}

func (repositorioPublicacoes RepositorioPublicacoes) CriarPublicacao(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repositorioPublicacoes.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

//BuscarPorID traz uma única publicação do banco de dados
func (repositorioPublicacoes RepositorioPublicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := repositorioPublicacoes.db.Query(`
		select p.*, u.nick FROM 
		publicacoes p INNER JOIN usuarios u
		ON u.id = p.autor_id WHERE p.id = ?`,
		publicacaoID,
	)

	if erro != nil {
		return models.Publicacao{}, erro
	}

	defer linha.Close()

	var publicacao models.Publicacao
	if linha.Next() {
		if erro := linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil

}

//BuscarPublicacoes traz as publicaçõs dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorioPublicacoes RepositorioPublicacoes) BuscarPublicacoes(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorioPublicacoes.db.Query(
		`SELECT DISTINCT p.*, u.nick FROM publicacoes p 
		 INNER JOIN usuarios u ON u.id = p.autor_id 
		 INNER JOIN seguidores s ON p.autor_id = s.usuario_id
		 WHERE u.id = ? OR s.seguidor_id = ?
		 ORDER BY 1 DESC
		`, usuarioID, usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao
	for linhas.Next() {
		var publicacao models.Publicacao
		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

//AtualizarPublicacao altera os dados de uma publicação no banco de dados
func (repositorioPublicacoes RepositorioPublicacoes) AtualizarPublicacao(publicacaoID uint64, publicacao models.Publicacao) error {
	statement, erro := repositorioPublicacoes.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	return nil

}

//DeletarPublicacao excluir uma publicacao do banco de dados
func (repositorioPublicacoes RepositorioPublicacoes) DeletarPublicacao(publicacaoID uint64) error {
	statement, erro := repositorioPublicacoes.db.Prepare(
		`DELETE FROM publicacoes WHERE id = ?
		`,
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil

}

//BuscarPublicacoesPorUsuario traz todas as publicações de um usuário especifico
func (repositorioPublicacoes RepositorioPublicacoes) BuscarPublicacoesPorUsuario(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorioPublicacoes.db.Query(`
			SELECT p.*, u.nick from publicacoes p
			JOIN usuarios u ON u.id = p.autor_id
			WHERE p.autor_id = ?		
	`, usuarioID)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao
	for linhas.Next() {
		var publicacao models.Publicacao
		if erro := linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

//CurtirPublicacao adiciona uma curtida a uma publicacao do banco de dados
func (repositorioPublicacoes RepositorioPublicacoes) CurtirPublicacao(publicacaoID uint64) error {
	statement, erro := repositorioPublicacoes.db.Prepare(
		`UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

//DescurtirPublicacao subtrai uma curtida a uma publicacao do banco de dados
func (repositorioPublicacoes RepositorioPublicacoes) DescurtirPublicacao(publicacaoID uint64) error {
	statement, erro := repositorioPublicacoes.db.Prepare(
		`UPDATE publicacoes SET curtidas = 
		 CASE 
				WHEN curtidas > 0 THEN curtidas - 1 
				ELSE 0 
		 END
		 WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
