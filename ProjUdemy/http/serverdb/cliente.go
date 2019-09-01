package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Struct Usuário
type Usuario struct {
	// Mapeando com Json
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	// Trim.Prefix remove tudo da url que estiver antes de /usuarios/
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	// As conversões numéricas mais comuns são Atoi (string para int) e Itoa (int para string).
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
		// Response, Request e o ID
	case r.Method == "GET":
		usuarioTodos(w, r)
		// Response, Request
	default:
		// Status não encontrado
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}
func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var u Usuario
	// QueryRow executa uma consulta que deve retornar no máximo uma linha.
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)
	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	// Imprima formatos usando os formatos padrão para seus operandos e grave em w.
	fmt.Fprint(w, string(json))
}
func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// QueryRow executa uma consulta que deve retornar no máximo uma linha.
	rows, _ := db.Query("select id, nome from usuarios")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
