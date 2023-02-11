package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nerdcademy/restapi/db"
	model "github.com/nerdcademy/restapi/models"
)

func GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Get Tasks"))

	var articles []model.Article
	db.DB.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	params := mux.Vars(r)
	db.DB.First(&article, params["id"])
	if article.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Article not found"))
		return
	}

	json.NewEncoder(w).Encode(article)

}

func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	json.NewDecoder(r.Body).Decode(&article)
	createdUser := db.DB.Create(&article)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&article)

}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Delete"))

	var article model.Article
	params := mux.Vars(r)
	db.DB.First(&article, params["id"])
	if article.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Article not found"))
		return
	}
	//db.DB.Delete(&user)              //No lo borra de la bd, sólo llena la propiedad deleted_at
	db.DB.Unscoped().Delete(&article) //Para borrarlo efectivamente de la bd y no sólo cambiar la propiedad deleted_at
	w.WriteHeader(http.StatusOK)
}

// func UpdatePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	decoder := json.NewDecoder(r.Body)
// 	var post model.Article
// 	err := decoder.Decode(&post)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	}

// 	err = model.UpdatePost(post)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		return
// 	} else {
// 		w.WriteHeader(http.StatusOK)
// 	}
// }
