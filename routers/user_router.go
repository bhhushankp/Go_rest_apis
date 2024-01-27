package routers

import (
	"user_api/handlers"

	"github.com/gorilla/mux"
)

func InitializeRoute() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/getuser/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/getalluser", handlers.GetAllUser).Methods("GET")
	r.HandleFunc("/createuser", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/updateuser/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/deleteuser/{id}", handlers.DeleteUser).Methods("DELETE")
	return r
}
