package main

import "runtime/debug"

func f3() {
	// O PrintStack imprime a pilha de execução no momento que a função é chamada	
	debug.PrintStack()
}

func f2(){
	f3()
}

func f1(){
	f2()
}

func main(){
	f1()
}