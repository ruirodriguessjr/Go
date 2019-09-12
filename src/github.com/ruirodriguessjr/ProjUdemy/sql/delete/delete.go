package main

import(
	_"github.com/go-sql-driver/mysql"
)

func main(){
	db, err:= sql.Open("mysql", "root:@/cursogo")
	if err != nil{
		return error.Wrap(err, "open mysql failed")
	}
	defer db.Close()

	delete, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil{
		return error.Wrap(err, "delete to usuarios failed")
	}

	delete.Exec(3)
	delete.Exec(2000)