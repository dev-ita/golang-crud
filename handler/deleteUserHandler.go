package handler

import (
	"crud/db"
	"net/http"
	"strconv"
)

// @Summary Deleta um usuário
// @Description Deleta um usuário correspondente ao ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param userId path int true "ID of the user to be deleted"
// @Success 204 "No Content"
// @Router /usuarios/{id} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("Id inválido"))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()
	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao preparar o statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(userId); err != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
