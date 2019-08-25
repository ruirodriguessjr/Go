package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Crio uma função onde instancio a variável que terá acesso as funções do meu banco de dados chamada db
// Passo a referência do meu sql onde ele retorna um resultado sql
func exec(db *sql.DB, sql string) sql.Result {
	// Atribuo a função na minha variável result
	// caso dê algum tipo de erro configuro o panic para informar o erro
	//ele para minha conexão
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	// caso de tudo certo, ele retorna minha conexão sql.result
	return result
}

func main() {

	// Abrindo minha conexão com o Banco
	db, err := sql.Open("mysql", "root:0000")
	if err != nil {
		panic(err)
	}

	// Após ele executar toda a minha instrução,
	//essa função fecha a conexão com o banco de dados
	defer db.Close()

	// Executando minha função para manuseio dos comandos do banco de dados
	exec(db, "create database if not exists cursogo")
	//exec(db, "use cursogo")
	//exec(db, "drop table if exists usuarios")
	//exec(db, `create table usuarios (
	//	id integer auto_increment,
	//	nome varchar(80),
	//	PRIMARY KEY (id)
	//)`)
}
