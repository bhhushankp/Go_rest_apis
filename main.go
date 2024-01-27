package main

import (
	"net/http"
	"user_api/routers"
)

func main() {

	router := routers.InitializeRoute()
	http.ListenAndServe(":8181", router)

}
