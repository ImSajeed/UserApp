package main

import (
	config "Auth/Config"
	Controllers "Auth/Controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()
	config.Load()
	r.HandleFunc("/api/Register", Controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/Login", Controllers.LoginUser).Methods("POST")
	r.HandleFunc("/api/Users", Controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/Users/me", Controllers.UserME).Methods("GET")
	r.HandleFunc("/api/Users/{userid}", Controllers.GetUserById).Methods("GET")
	r.HandleFunc("/api/Users/{userid}", Controllers.UpdateUser).Methods("PUT")

	nMiddleware := negroni.New()
	nMiddleware.UseHandler(r)
	nMiddleware.Run(":9999")

}
