package models

type ErrorDTO struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}