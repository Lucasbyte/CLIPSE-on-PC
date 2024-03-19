package routes

import (
	"net/http"

	"github.com/lucasbyte/go-clipse/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/drop", controllers.Drop)
	http.HandleFunc("/file", controllers.File)
	http.HandleFunc("/push", controllers.Push)
}
