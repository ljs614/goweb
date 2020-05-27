package main

import (
	"net/http"

	"github.com/ljs614/goweb/myapp"
)

func main() {
	http.ListenAndServe(":3030", myapp.NewHTTPHandler())
}
