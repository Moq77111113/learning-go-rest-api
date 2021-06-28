package app

import (
	"fmt"
	"log"
	"net/http"

	"moq.com/test/cmd/models"
	"moq.com/test/cmd/helpers"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the api")
	}
}

func (a *App) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := helpers.Parse(w, r, &req)
		if err != nil {
			log.Printf("Unable to parse body, err=%v \n", err)
			helpers.SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		p := &models.Post{
			ID: 0,
			Title: req.Title,
			Content: req.Content,
			Author: req.Author,
		}

		err = a.DB.Create(p)
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			helpers.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := helpers.MapToJson(p)
		helpers.SendResponse(w, r, resp, http.StatusCreated)
	}
}

func (a *App) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := a.DB.Get()
		if err != nil {
			log.Printf("Unable to get posts, err=%v, \n", err)
			helpers.SendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, len(posts))
		for idx, post := range posts {
			resp[idx] = helpers.MapToJson(post)
		}
		helpers.SendResponse(w, r, resp, http.StatusOK)
	}
}
