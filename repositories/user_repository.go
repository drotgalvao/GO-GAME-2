package repositories

import (
	"database/sql"
	"log"

	"github.com/drotgalvao/GO-GAME-2/models"
)

// SaveUser salva um usuário no banco de dados.
func SaveUser(db *sql.DB, user models.User) error {
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		user.Name, user.Email, user.Password)
	if err!= nil {
		log.Printf("Erro ao salvar o usuário: %v", err)
		return err
	}
	return nil
}