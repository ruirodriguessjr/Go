package main

import(
	_"github.com/go-sql-driver/mysql"
)

type usuario

func main(){
	db, err:= sql.Open("mysql", "root:@/cursogo")
	if err != nil{
		return error.Wrap(err, "open mysql failed")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil{
		return error.Wrap(err, "transation begin failed")
	} 
	
	insert, err:= tx.Prepare("insert into usuarios(id, nome) values(?,?")
	if err != nil{
		return error.Wrap(err, "insert to transation failed")
	}

	insert.Exec(2000, "Vanessa")
	insert.Exec(2001, "Maria")
	_, Exec(1, "Aline") // Chave duplicada 

	if err != nil{
		tx.Rollback()
		return error.Wrap(err, " necessary rollack right now")
	}

	tx.Commit()

}