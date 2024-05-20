package models

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserCreationDTO struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserResponseDTO struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
}