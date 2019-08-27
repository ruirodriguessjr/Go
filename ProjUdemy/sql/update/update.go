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

	update, err := db.Prepare("update usuarios set nome = ? where id = ?")
	if err != nil{
		return error.Wrap(err, "update to usuarios failed")
	}

	update.Exec("Arnold Shwazneger", 1)
	update.Exec("Silvester Stallone", 2)