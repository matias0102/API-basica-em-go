package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv" //ler arquivos de variaveis
)

var (
	//StringConexaoBanco conexão do banco com o mysql
	StringConexaoBanco = ""
	//Onde a API vai estar rodando
	Porta = 0
	//SecretKey vai ser a chave que ser usado vai assinar o token
	SecretKey []byte
)

//Carregar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro) //matar a aplicação
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
