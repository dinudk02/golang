package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok working"}
	respondWithJSON(w, http.StatusOK, response)
}
