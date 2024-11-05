// pacote db
package db

//chama os importes e as bibliotecas do SQL
import (

	//habilita para usar os comando sql
	"database/sql"

	//importa o driver do postgresql
	_ "github.com/lib/pq"
)

// criar a função acessar e habilita o uso do database postgresql
func Acesse() *sql.DB {

	//criar uma variavel chamada(conexao) e realiza a verificação se os dados
	//De conexao estão corretos
	var conexao = "user=postgres dbname=sgco host=localhost password=1234 sslmode=disable"

	// criar uma variavel chamada(db e err) e abrir uma conexao com o
	//banco de dados postgresql
	db, err := sql.Open("postgres", conexao)
	//Se caso acontecer um erro nos dados informados acima os
	//comandos mostraram os erros de conexao
	if err != nil {
		panic(err)
	}
	//Retorna o valor verdadeiro da conexao com o banco de dados postgresql
	return db
}
