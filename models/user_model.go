package models

import (
    "github.com/google/uuid"
)

type User struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
    Password string    `json:"password"`
}

type UserCreationDTO struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserResponseDTO struct {
    ID       uuid.UUID `json:"id"`
    Name     string    `json:"name"`
    Email    string    `json:"email"`
}