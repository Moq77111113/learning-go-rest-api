package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"moq.com/test/cmd/models"
)

func Parse (w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func SendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader((status))

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot forma json, err=%v \n", err)
	}
}

func MapToJson(p *models.Post) models.JsonPost {
	return models.JsonPost{
		ID: p.ID,
		Title: p.Title,
		Content: p.Content,
		Author: p.Author,
	}
}