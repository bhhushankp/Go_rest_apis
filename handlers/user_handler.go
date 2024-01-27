package handlers

import (
	"encoding/json"
	"net/http"
	"user_api/models"
	"user_api/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	// Database connection
	utils.ConnectDB()
	Db = utils.GetDB()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//db := utils.GetDB()

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid User", http.StatusBadRequest)
	}
	Db.Create(&user)
	respondWithJSON(w, http.StatusCreated, user)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	err := Db.First(&user, params["id"]).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, user)

}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	Db.Find(&users)
	respondWithJSON(w, http.StatusOK, users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	if err := Db.First(&user, params["id"]).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	Db.Save(&user)
	respondWithJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var user models.User

	if err := Db.First(&user, params["id"]).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	Db.Delete(&user)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})

}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)

}
