package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dinudk02/golang/internal/database" //  database link

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}
	id := uuid.New()
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        id,
		Username:  params.Name,
		CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		Email:     params.Email,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("coundt create user %v", err))
		return
	}

	respondWithJSON(w, 200, user)
}
