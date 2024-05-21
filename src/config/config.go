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
func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	wd, _ := os.Getwd()
	port := fmt.Sprintf(wd + os.Getenv("API_PORT"))
	user := fmt.Sprintf(wd + os.Getenv("DB_USER"))
	pass := fmt.Sprintf(wd + os.Getenv("DB_PASSWORD"))
	name := fmt.Sprintf(wd + os.Getenv("DB_NAME"))

	Porta, erro = strconv.Atoi(port)
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		name,
	)
}
