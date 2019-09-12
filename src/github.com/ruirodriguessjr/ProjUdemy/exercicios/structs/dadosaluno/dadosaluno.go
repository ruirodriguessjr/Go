package main

import "fmt"

type aluno struct {
	nome      string
	matricula string
	curso     string
}

type dado struct {
	dados []aluno
}

func (a *dado) addAluno(novo aluno) {

	a.dados = append(a.dados, novo)

}

func main() {

	aluno1 := aluno{nome: "Rui Rodrigues", matricula: "3318-542", curso: "SI"}
	aluno2 := aluno{nome: "Rui Santos", matricula: "3318-456", curso: "SA"}
	aluno3 := aluno{nome: "Rui Junior", matricula: "3318-789", curso: "SE"}
	aluno4 := aluno{nome: "Rui Juliani", matricula: "3318-321", curso: "SO"}
	aluno5 := aluno{nome: "Rui Cardoso", matricula: "3318-654", curso: "SU"}

	a := dado{}
	a.addAluno(aluno1)
	a.addAluno(aluno2)
	a.addAluno(aluno3)
	a.addAluno(aluno4)
	a.addAluno(aluno5)
	//a.addAluno(aluno2)
	//a.addAluno(aluno3)
	//a.addAluno(aluno4)
	//a.addAluno(aluno5)
	fmt.Println(len(a.dados))
}
