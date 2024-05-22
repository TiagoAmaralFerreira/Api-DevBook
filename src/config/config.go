package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
)

// Load vai inicializar as variaveis de ambiente
func Load() error {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	wd, _ := os.Getwd()
	port := fmt.Sprintf(wd + os.Getenv("API_PORT"))
	user := fmt.Sprintf(wd + os.Getenv("DB_USER"))
	pass := fmt.Sprintf(wd + os.Getenv("DB_PASSWORD"))
	name := fmt.Sprintf(wd + os.Getenv("DB_NAME"))
	host := fmt.Sprintf(wd + os.Getenv("DB_HOST"))

	Porta, erro = strconv.Atoi(port)
	if erro != nil {
		Porta = 5432
	}
	// postgres://%s:%s@%s/devbook
	StringConexaoBanco = fmt.Sprintf("postgres://%s:%s@%s/%s",
		user,
		pass,
		host,
		name,
	)

	return nil
}
