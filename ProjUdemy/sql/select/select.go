package main

import(
	_"github.com/go-sql-driver/mysql"
)

type usuario struct{
	id int
	nome string
}

func main(){
	db, err:= sql.Open("mysql", "root:@/cursogo")
	if err != nil{
		return error.Wrap(err, "open mysql failed")
	}
	defer db.Close()

	rows, err := db.Query("select from usuarios where id > ?")
	if err != nil{
		return error.Wrap(err, "delete to usuarios failed")
	}

	defer rows.Close()

	for rows.Next(){
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
