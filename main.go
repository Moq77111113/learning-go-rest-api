package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome !")
	fmt.Println("Endpoint Hit homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/all", returnAllArticles)
	router.HandleFunc("/article/{id}", returnArticleById).Methods("GET")
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))

}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "This is an Article", Content: "Some beautiful text"},
		{Id: "2", Title: "olleH", Desc: "Another Article", Content: "Blablablabla"},
	}
	handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for _, article := range Articles {
		if article.Id == vars["id"] {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(body, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)

	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upd")
	vars := mux.Vars(r)
	body, _ := ioutil.ReadAll(r.Body)
	id := vars["id"]
	var newVersion Article
	json.Unmarshal(body, &newVersion)
	for index, article := range Articles {
		if article.Id == id {
			Articles[index] = newVersion
		}
	}
	json.NewEncoder(w).Encode(newVersion)
}
