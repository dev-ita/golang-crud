package handler

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}
	var usuario usuario
	err = json.Unmarshal(bodyRequest, &usuario)
	if err != nil {
		w.Write([]byte("Falha ao desserializar JSON"))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao inserir usuário"))
	}
	defer statement.Close()
	insert, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserted, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserted)))
}
