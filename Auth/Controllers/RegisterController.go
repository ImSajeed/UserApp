package Controllers

import (
	utils "Auth/Utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var register utils.Register
	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if register.Username != "" && register.Password != "" {
		inserted := utils.RegisterUserInDB(register)
		if inserted >= 1 {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			outResponse, _ := json.Marshal("Registered Successfully!")
			w.Write(outResponse)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var login utils.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if login.Username != "" && login.Password != "" {
		isexist := utils.IsLoginExists(login)
		if isexist {
			token, err := utils.CreateToken(login.Username)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			resp := utils.LoginResponse{Token: token}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			outResponse, _ := json.Marshal(resp)
			w.Write(outResponse)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			outResponse, _ := json.Marshal("Invalid data")
			w.Write(outResponse)
			return
		}
	}

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	username, err := utils.TokenValid(r)
	fmt.Printf("%v", username)
	if err == nil && utils.IsAdmin(username) {
		data := utils.GetAllUsers()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		outResponse, _ := json.Marshal(data)
		w.Write(outResponse)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		outResponse, _ := json.Marshal("Invalid data")
		w.Write(outResponse)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userid, _ := strconv.Atoi(params["userid"])
	username, err := utils.TokenValid(r)
	fmt.Printf("%v", username)
	if err == nil && utils.IsAdmin(username) {
		data := utils.GetUserById(userid)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		outResponse, _ := json.Marshal(data)
		w.Write(outResponse)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		outResponse, _ := json.Marshal("Invalid data")
		w.Write(outResponse)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userid, _ := strconv.Atoi(params["userid"])
	username, err := utils.TokenValid(r)
	fmt.Printf("%v", username)
	if err == nil && utils.IsAdmin(username) {
		data := utils.UpdateUser(userid)
		fmt.Printf("records updated ", data)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		outResponse, _ := json.Marshal(fmt.Sprintf("update admin access for username %s", username))
		w.Write(outResponse)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		outResponse, _ := json.Marshal("Invalid data")
		w.Write(outResponse)
	}
}

func UserME(w http.ResponseWriter, r *http.Request) {
	username, err := utils.TokenValid(r)
	fmt.Printf("name is %v", username)
	if err == nil && utils.IsUserRole(username) {
		data := utils.GetUserByName(username)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		outResponse, _ := json.Marshal(data)
		w.Write(outResponse)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		outResponse, _ := json.Marshal("Invalid data")
		w.Write(outResponse)
	}
}
