package main

import (
	"net/http"

	"github.com/reallycky84/GolangStudy/web3/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
