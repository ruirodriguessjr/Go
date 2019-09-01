package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Abertura da conexão com o banco
	db, err := sql.Open("mysql", "root:@/cursogo")
	if err != nil {
		panic(err)
	}

	// Fechamento da conexão com o banco com função defer
	// Ou seja, fechamentoa pós a execução de todo o main
	defer db.Close()

	// inserção de dados para os campos da tabela
	// Coloquei ? podendo adicionar vários parâmetros para os campos sequenciamente
	// Prepare cria uma declaração preparada para consultas ou execuções posteriores.
	insert, err := db.Prepare("insert into usuarios(nome) values(?)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec("Maria")
	insert.Exec("João")

	res, err := insert.Exec("Pedro")
	if err != nil {
		panic(err.Error())
	}

	// LastInserId = mostra o ultimo Id que foi inserido
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id)

	// RowsAffected = Mostra quantas linhas foram afetadas
	linhas, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(linhas)
}
