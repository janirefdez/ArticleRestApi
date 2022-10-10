package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/janirefdez/ArticleRestApi/pkg/mocks"
	"github.com/janirefdez/ArticleRestApi/pkg/models"
)

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedArticle models.Article
	json.Unmarshal(body, &updatedArticle)

	for index, article := range mocks.Articles {
		if article.Id == id {
			article.Title = updatedArticle.Title
			article.Desc = updatedArticle.Desc
			article.Content = updatedArticle.Content

			mocks.Articles[index] = article

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
