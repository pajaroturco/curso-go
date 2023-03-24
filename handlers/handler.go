package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"pajaro.com/curso-go/db"
	"pajaro.com/curso-go/models"
)

func getUserById(r *http.Request) models.User {
	user := models.User{}
	vars := mux.Vars(r)
	id , _:= strconv.Atoi(vars["id"]) 
	db.Database.First(&user, id)

	return user
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	users := models.Users{}

	db.Database.Find(&users)

	sendData(w, users, http.StatusOK)

}


func GetUser(w http.ResponseWriter, r *http.Request){
	user := getUserById(r)

	if user.Id == 0 {
		sendError(w, errors.New("Not Found"), http.StatusNotFound)
		return
	}

	sendData(w, user, http.StatusOK)
	
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	user := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	db.Database.Create(&user)

	sendData(w, user, http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	user := getUserById(r)

	if user.Id == 0 {
		sendError(w, errors.New("Not Found"), http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	db.Database.Save(&user)

	sendData(w, user, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	user := getUserById(r)

	if user.Id == 0 {
		sendError(w, errors.New("Not Found"), http.StatusNotFound)
		return
	}

	db.Database.Delete(&user)

	sendData(w, user, http.StatusOK)
}
