package main

import (
	"net/http"

	"github.com/reallucky84/GolangStudy/web5/myapp"
)

func main(){

	http.ListenAndServe(":3000", myapp.NewHandler())
}