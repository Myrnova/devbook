package security

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerficiarSenha compara uma senha e um hash e retorna se elas são iguais
func VerificarSenha(senhaComHash []byte, senha string) error {
	return bcrypt.CompareHashAndPassword(senhaComHash, []byte(senha))
}
