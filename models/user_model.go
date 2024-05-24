package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserCreationDTO represents the data needed to create a new user.
//	@Description:	Represents the data structure for creating a new user account.
type UserCreationDTO struct {
    // Name of the user.
    //	@Description:	The full name of the user.
    //	@Required:		true
    Name string `json:"name"`

    // Email address of the user.
    //	@Description:	The email address used to identify the user.
    //	@Required:		true
    Email string `json:"email"`

    // Password of the user.
    //	@Description:	The password chosen by the user.
    //	@Required:		true
    Password string `json:"password"`

    // Confirmation of the user's password.
    //	@Description:	The confirmation of the user's password to ensure accuracy.
    //	@Required:		true
    ConfirmPassword string `json:"confirm_password"`
    // Bio             *string `json:"bio,omitempty" description:"Bio of the user."` // optional
}

type UserResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
