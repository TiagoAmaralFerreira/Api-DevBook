package config

import (
	"fmt"
	"log"
	"os"
	"runtime"
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

	port := fmt.Sprintf(os.Getenv("API_PORT"))
	user := fmt.Sprintf(os.Getenv("DB_USER"))
	pass := fmt.Sprintf(os.Getenv("DB_PASSWORD"))
	name := fmt.Sprintf(os.Getenv("DB_NAME"))
	host := fmt.Sprintf(os.Getenv("DB_HOST"))

	Porta, erro = strconv.Atoi(port)
	if erro != nil {
		Porta = 5000
	}

	var wd string
	if runtime.GOOS != "windows" {
		wd, _ = os.Getwd()
	}

	// Construir a string de conexão corretamente
	StringConexaoBanco = fmt.Sprintf("postgres://%s:%s@%s/%s",
		user,
		pass,
		host,
		name,
	)

	if wd != "" {
		StringConexaoBanco = fmt.Sprintf("%s/%s", wd, StringConexaoBanco)
	}

	return nil

}
