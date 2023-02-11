package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nerdcademy/restapi/db"
	"github.com/nerdcademy/restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Get Users"))
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(user)

}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Delete"))

	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	//db.DB.Delete(&user)              //No lo borra de la bd, sólo llena la propiedad deleted_at
	db.DB.Unscoped().Delete(&user) //Para borrarlo efectivamente de la bd y no sólo cambiar la propiedad deleted_at
	w.WriteHeader(http.StatusOK)
}
