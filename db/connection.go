package db

import (
    "database/sql"
    "log"
	"os"

	"fmt"

    _ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")
    sslMode := os.Getenv("DB_SSLMODE")

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        username,
        password,
        host,
        port,
        dbname,
        sslMode)

    dbConn, err := sql.Open("postgres", connStr)
    if err!= nil {
        log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
    }

    err = dbConn.Ping()
    if err!= nil {
        log.Fatalf("Erro ao fazer ping no banco de dados: %v", err)
    }

    log.Println("Conectado ao banco de dados com sucesso!!")

    return dbConn, nil
}