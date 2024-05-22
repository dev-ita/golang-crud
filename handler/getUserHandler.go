package handler

import (
	"crud/db"
	"crud/entitie"
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary Retorna os usuários
// @Description  Retorna todos os detalhes dos usuários
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} entitie.Usuario
// @Router /usuarios [get]
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	db, err := db.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer rows.Close()

	var usuarios []entitie.Usuario
	for rows.Next() {
		var usuario entitie.Usuario
		if err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao enviar os usuários em forma de JSON"))
		return
	}
}

// @Summary Retorna um único usuário por ID
// @Description Retorna todos os detalhes do usuário através de um ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param userId path int true "ID of the user"
// @Success 200 {object} entitie.Usuario
// @Router /usuarios/{id} [get]
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("Id inválido"))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao se conectar no banco de dados"))
		return
	}
	defer db.Close()

	row, err := db.Query("select * from usuarios where id = ?", userId)
	if err != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}

	var usuario entitie.Usuario
	if row.Next() {
		if err := row.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.Write([]byte("Erro ao converter usuário para JSON!"))
		return
	}
}
