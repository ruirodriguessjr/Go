package main

func main(){
	fs := http.FileServer(http.Dir("public"))
	http.Handler("/", fs)

	log.Println("Executando")
	log.Fatal(http.ListenAndServer(":3000", nil))
}