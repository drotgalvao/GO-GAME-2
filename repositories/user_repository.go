// user_repository.go
package repositories

import (
	"database/sql"
	"log"

	"github.com/drotgalvao/GO-GAME-2/models"
	"github.com/google/uuid"
)

func SaveUser(db *sql.DB, user models.UserCreationDTO) (*models.UserResponseDTO, error) {
	row := db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email",
		user.Name, user.Email, user.Password)

	var userID uuid.UUID
	var userName, userEmail string
	err := row.Scan(&userID, &userName, &userEmail)
	if err != nil {
		log.Printf("Error saving user: %v\n", err)
		return nil, err
	}

	return &models.UserResponseDTO{ID: userID, Name: userName, Email: userEmail}, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.UserResponseDTO, error) {
	var user models.UserResponseDTO
	err := db.QueryRow("SELECT id, name, email FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Printf("Error getting user: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func CheckIfUserExistsByEmail(db *sql.DB, email string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		log.Printf("Error checking if user exists: %v\n", err)
		return false, err
	}
	return count > 0, nil
}