package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func getUserInfo89Handler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Fprint(w, "User Id:", vars["id"])
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "web5 Hello world!")
}

func usersHandler(w http.ResponseWriter, r *http.Request){	
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

// NewHandler make new http.Handler
func NewHandler() http.Handler{

	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/users", usersHandler)

	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfo89Handler)
	return mux
}