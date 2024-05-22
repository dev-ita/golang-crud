package handler

import (
	"crud/db"
	"crud/entitie"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// @Summary Atualiza um usuário
// @Description Atualiza um usuário através do Id correspondente
// @Tags users
// @Accept  json
// @Produce  json
// @Param userId path int true "ID of the user to be updated"
// @Success 200 {object} entitie.Usuario
// @Router /usuarios/{id} [put]
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("Id inválido"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario entitie.Usuario
	if err := json.Unmarshal(requestBody, &usuario); err != nil {
		w.Write([]byte("Erro ao converter usuário para struct"))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("erro ao criar statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Email, userId); err != nil {
		w.Write([]byte("Erro ao atualizar usuário"))
		return
	}

	usuario.Id = uint32(userId)

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao converter usuário para json"))
		return
	}

}
