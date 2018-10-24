package users

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	users := GetUsers()

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	json.NewEncoder(writer).Encode(users)
}

func Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	user := GetUser(id)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	json.NewEncoder(writer).Encode(user)
}

func Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var user User
	json.NewDecoder(request.Body).Decode(&user)

	user = CreateUser(user)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(201)

	json.NewEncoder(writer).Encode(user)
}

func Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	var user User
	json.NewDecoder(request.Body).Decode(&user)

	user = UpdateUser(id, user)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)

	json.NewEncoder(writer).Encode(user)
}

func Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	DeleteUser(id)
	writer.WriteHeader(204)
}
