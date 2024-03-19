package main

import (
	"net/http"

	"github.com/lucasbyte/go-clipse/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
