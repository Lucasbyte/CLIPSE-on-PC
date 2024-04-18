package main

import (
	"fmt"
	"net/http"

	"github.com/lucasbyte/go-clipse/routes"
)

func main() {
	fmt.Println()
	routes.CarregaRotas()
	//models.ItensMGV("Teste")
	http.ListenAndServe(":8000", nil)
}
