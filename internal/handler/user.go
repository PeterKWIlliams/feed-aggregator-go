package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
	"github.com/PeterKWIlliams/feed-aggregator-go/internal/models"
)

func (c *Config) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Couldn't decode request body")
		return
	}

	user, err := c.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "User creation failed")
		return
	}
	RespondWithJSON(w, http.StatusCreated, models.DatabaseUsertoUser(user))
}

func (c *Config) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	RespondWithJSON(w, http.StatusAccepted, models.DatabaseUsertoUser(user))
}
