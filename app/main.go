package main

//https://youtu.be/B6gQ1B0cn4s
//https://faztweb.com/contenido/docker-postgresql
import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nerdcademy/restapi/db"
	"github.com/nerdcademy/restapi/models"
	"github.com/nerdcademy/restapi/routes"
)

func main() {

	db.Init()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Article{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	//Users routes
	router.HandleFunc("/api/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/api/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/api/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/api/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//Tasks routes
	router.HandleFunc("/api/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/tasks", routes.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	//Articles
	router.HandleFunc("/api/posts", routes.GetArticlesHandler).Methods("GET")
	router.HandleFunc("/api/posts/{id}", routes.GetArticleHandler).Methods("GET")
	router.HandleFunc("/api/posts", routes.PostArticleHandler).Methods("POST")
	//router.HandleFunc("/api/post/update", controller.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", routes.DeleteArticleHandler).Methods("DELETE")

	http.ListenAndServe(":3200", router)

}
