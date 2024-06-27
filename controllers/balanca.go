package controllers

import (
	"net/http"

	global "github.com/lucasbyte/go-clipse/Global"
	"github.com/lucasbyte/go-clipse/models"
)

func Import(w http.ResponseWriter, r *http.Request) {
	global.SetStatus(false)
	importData := models.BuscaEventoImport()
	temp.ExecuteTemplate(w, "Import", importData)
}

func Load(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Load", nil)
}

// func Send(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		portaStr := r.FormValue("porta")
// 		porta, _ := strconv.Atoi(portaStr)
// 		tipo := r.FormValue("tipo")
// 		velocidadeStr := r.FormValue("velocidade-select")
// 		velocidade, _ := strconv.Atoi(velocidadeStr)
// 		if velocidade != 115200 {
// 			velocidade = 9600
// 		}

// 		temp.ExecuteTemplate(w, "Load", nil)
// 		plus, err := models.ObterDadosProdutos()
// 		if tipo == "1" {
// 			plus, err := models.ObterCodigosFaltantes()
// 			serial.Delete(porta, velocidade, plus)
// 			if err != nil {
// 				http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 			}
// 		}
// 		progress := 100 / (len(plus))
// 		portaCOM, _ := serial.Porta(porta, velocidade)
// 		defer portaCOM.Close()
// 		for _, plu := range plus {
// 			progress += serial.EnviarDado(portaCOM, plu)
// 			if err != nil {
// 				http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 			}
// 		}
// 	}
// 	http.Redirect(w, r, "/", http.StatusMovedPermanently)
// }
