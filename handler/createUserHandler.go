package handler

import (
	"crud/db"
	"crud/entitie"
	"encoding/json"
	"io"
	"net/http"
)

// @Summary Cria um usuário
// @Description Cria um usuário no banco de dados
// @Tags users
// @Accept  json
// @Produce  json
// @Param usuario body entitie.Usuario true "Cria um usuário"
// @Success 200 {object} entitie.Usuario
// @Router /usuarios [post]
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}
	var usuario entitie.Usuario
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

	usuario.Id = uint32(idInserted)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usuario)
	if err != nil {
		w.Write([]byte("Erro ao criar usuário"))
		return
	}
}
