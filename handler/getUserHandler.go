package handler

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all users"))
}