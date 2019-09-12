package main

import (
	"encoding/json"
	"fmt"
)

// ID = Público
// id = Privado
type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {

	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	// Marshal semelhante ao NewEncoder onde ele retorna dois valores,
	// sendo o texto em JSON e o Erro, percorrendo o seu todo para saber se é nulo ou não,
	// assim decidindo se retornará erro ou não de acordo com o conteúdo de dentro
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString :=
		`{"id": 2, "nome": "Caneta", "preco": 8.90, "tags" :["Papelaria","Importado"]}`
	// O mesmo de Marshal serve para o unmarshal
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[0], p2.Tags[1])

}
