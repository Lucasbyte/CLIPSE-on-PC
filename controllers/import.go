package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	global "github.com/lucasbyte/go-clipse/Global"
	"github.com/lucasbyte/go-clipse/models"
	"github.com/lucasbyte/go-clipse/serial"
)

func Importeste(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" && !global.GetStatus() {
		global.SetStatus(true)
		portaStr := r.FormValue("porta")
		porta, _ := strconv.Atoi(portaStr)
		tipo := r.FormValue("tipo")
		velocidadeStr := r.FormValue("velocidade-select")
		Ehcompleto := r.FormValue("import-select")
		velocidade, _ := strconv.Atoi(velocidadeStr)
		dataUltimoEnvio := models.BuscaEventoImport()

		if velocidade != 115200 {
			velocidade = 9600
		}
		plus, err := models.ObterDadosProdutos()
		if err != nil {
			fmt.Println(err)
		}
		if tipo == "1" {
			plus, err := models.ObterCodigosFaltantes()
			serial.Delete(porta, velocidade, plus)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusMovedPermanently)
			}
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming not supported", http.StatusInternalServerError)
			return
		}

		fmt.Println(plus)
		progress := 0.0
		portaCOM, _ := serial.Porta(porta, velocidade)
		defer portaCOM.Close()
		progressUnit := (100.0 / float64(len(plus)))

		for _, plu_Cod := range plus {
			if !global.GetStatus() {
				break
			}
			plu_e_cod := strings.Split(plu_Cod, "£")
			plu := plu_e_cod[0]
			fmt.Println(plu)
			cod, _ := strconv.Atoi(plu_e_cod[1])
			produto, _ := models.ObterProduto(cod)
			dataUpdateProduto := produto.UpdatedAt
			if produto.Plu == 199 {
				fmt.Println("oi")
			}
			if Ehcompleto != "Total" {
				if dataUltimoEnvio.EventDate.After(dataUpdateProduto) {
					progress += 1 * progressUnit
					fmt.Fprintf(w, "data: {\"progress\": %s}\n\n", fmt.Sprint(progress))
					time.Sleep(time.Microsecond * 20)
					flusher.Flush()
					fmt.Println(progress)
					if !global.GetStatus() {
						break
					}
					continue
				}
			}
			progress += float64(serial.EnviarDado(portaCOM, plu)) * progressUnit
			fmt.Fprintf(w, "data: {\"progress\": %s}\n\n", fmt.Sprint(progress))
			flusher.Flush()
			fmt.Println(progress)
			if !global.GetStatus() {
				break
			}
		}
		fmt.Fprintf(w, "data: {\"complete\": true}\n\n")
		flusher.Flush()
		models.UpdateEvento("import", time.Now())
		global.SetStatus(false)
	}
}

// func SendProducts(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		portaStr := r.FormValue("porta")
// 		porta, _ := strconv.Atoi(portaStr)
// 		tipo := r.FormValue("tipo")
// 		Ehcompleto := r.FormValue("import-select")
// 		velocidadeStr := r.FormValue("velocidade-select")
// 		velocidade, _ := strconv.Atoi(velocidadeStr)

// 		dataUltimoEnvio := models.BuscaEventoImport()

// 		if velocidade != 115200 {
// 			velocidade = 9600
// 		}
// 		plus, _ := models.ObterDadosProdutos()
// 		if tipo == "1" {
// 			plus, err := models.ObterCodigosFaltantes()
// 			serial.Delete(porta, velocidade, plus)
// 			if err != nil {
// 				http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 			}
// 		}
// 		portaCOM, _ := serial.Porta(porta, velocidade)
// 		defer portaCOM.Close()
// 		progress := 100 / (len(plus))
// 		for _, plu_Cod := range plus {
// 			plu_e_cod := strings.Split(plu_Cod, "£")
// 			plu := plu_e_cod[0]
// 			cod, _ := strconv.Atoi(plu_e_cod[1])
// 			produto, err := models.ObterProduto(cod)
// 			dataUpdateProduto := produto.UpdatedAt
// 			if Ehcompleto != "Total" {
// 				if dataUltimoEnvio.EventDate.After(dataUpdateProduto) {
// 					continue
// 				}
// 			}
// 			progress += serial.EnviarDado(portaCOM, plu)
// 			if err != nil {
// 				http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 			}
// 		}
// 	}

// }

// func ToImport(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		flusher, ok := w.(http.Flusher)
// 		if !ok {
// 			http.Error(w, "Streaming not supported", http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "text/event-stream")
// 		w.Header().Set("Cache-Control", "no-cache")
// 		w.Header().Set("Connection", "keep-alive")
// 		var progress int
// 		var balancas_import []models.Balanca
// 		var balancas_import_error []models.Balanca
// 		balancas, err := models.BuscaBalancas()
// 		if err != nil {
// 			fmt.Fprintf(w, "event: error\ndata: %s\n\n", err.Error())
// 			flusher.Flush()
// 			return
// 		}

// 		totalSteps := len(balancas) + 1
// 		currentStep := 0

// 		checkImport := r.FormValue("Importar-checkbox")
// 		if checkImport == "on" {

// 		}

// 		{

// 			time.Sleep(500 * time.Millisecond)
// 			balancas_conect, err := json.Marshal(balancas_import)
// 			if err != nil {
// 				fmt.Fprintf(w, "data: {\"error\": \"%s\"}\n\n", err.Error())
// 				flusher.Flush()
// 				return
// 			}
// 			fmt.Fprintf(w, "data: {\"conect\": %s}\n\n", balancas_conect)
// 			flusher.Flush()
// 			time.Sleep(time.Millisecond * 500)

// 			progress = 25
// 			fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 			flusher.Flush()

// 			var balancas_import_checked []models.Balanca
// 			formatos := file.LerTipoJson()
// 			if formatos == "TXITENS" {
// 				arquivos := txitens.ReadTxitensJson()
// 				itensFile := arquivos.Caminhos.Itens
// 				err := models.Txitens(itensFile, balancas_import)
// 				if err == nil {
// 					progress += 75
// 					fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 					flusher.Flush()
// 				}
// 			} else if formatos == "Cad" {
// 				arquivos := Cad.ReadCadJson()
// 				itensFile := arquivos.Caminhos_Cad.Itens_Cad
// 				_ = arquivos.Caminhos_Cad.Receita_Cad
// 				_ = arquivos.Caminhos_Cad.CampoExtra_Cad

// 				itens := Cad.CadToItens(itensFile)

// 				if nomeDaPasta := "SYSTEL-ARQUIVOS/"; len(itens) > 3 {
// 					balancas_import_checked, balancas_import_error = file.Passo2(itens, nomeDaPasta, balancas_import)
// 					progress += (75 * (len(balancas_import) - len(balancas_import_error)) / len(balancas_import))
// 					fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 					flusher.Flush()
// 				}

// 			} else {
// 				arquivos := file.ReadMGVJson()
// 				itensFile := arquivos.Caminhos.Itens
// 				receitaFile := arquivos.Caminhos.Receita
// 				nutriFile := arquivos.Caminhos.Nutricional
// 				campoextra := arquivos.Caminhos.CampoExtra
// 				fornFile := arquivos.Caminhos.Fornecedor
// 				fracionaFile := arquivos.Caminhos.Fracionador
// 				taraFile := arquivos.Caminhos.Tara
// 				conservaFile := arquivos.Caminhos.Conservantes
// 				somente_preco_form := r.FormValue("somente_preco")
// 				somente_preco := somente_preco_form == "on"
// 				erroLeitura, arquivo, nomeDaPasta, dict_nutri, info, dict_forn, dict_aler, dict_fraciona, dict_conserva, dict_tara := file.Passo1(itensFile, receitaFile, nutriFile, fracionaFile, fornFile, taraFile, conservaFile, campoextra)
// 				if erroLeitura != nil {
// 					progress = -1
// 					fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 					flusher.Flush()
// 					time.Sleep(time.Millisecond * 500)
// 					//http.Redirect(w, r, "/Leitura500", 404)
// 					return

// 				} else {
// 					progress += 25
// 					fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 					flusher.Flush()
// 					balancas_import_checked, balancas_import_error = file.Passo2(arquivo, nomeDaPasta, balancas_import)
// 					progress += (25 * (len(balancas_import) - len(balancas_import_error)) / len(balancas_import))
// 					fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 					flusher.Flush()

// 					balancas_import_status, err := json.Marshal(balancas_import_checked)
// 					if err != nil {
// 						fmt.Fprintf(w, "data: {\"error\": \"%s\"}\n\n", err.Error())
// 						flusher.Flush()
// 						return
// 					}
// 					fmt.Fprintf(w, "data: {\"step2\": %s}\n\n", balancas_import_status)
// 					flusher.Flush()

// 					time.Sleep(time.Millisecond * 500)
// 				}
// 				if err == nil && !somente_preco && len(balancas_import_checked) > 0 {
// 					balancas_import_status_extras, balancas_extras_error := file.Passo3(arquivo, nomeDaPasta, balancas_import_checked, dict_nutri, info, dict_forn, dict_aler, dict_fraciona, dict_conserva, dict_tara)
// 					progress += (25 * (len(balancas_import) - len(balancas_extras_error)) / len(balancas_import))
// 					fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 					flusher.Flush()
// 					time.Sleep(time.Millisecond * 500)
// 					balancas_import_status, err := json.Marshal(balancas_import_status_extras)
// 					if err != nil {
// 						fmt.Fprintf(w, "data: {\"error\": \"%s\"}\n\n", err.Error())
// 						flusher.Flush()
// 						return
// 					}
// 					fmt.Fprintf(w, "data: {\"step3\": %s}\n\n", balancas_import_status)
// 					flusher.Flush()
// 				}
// 				if len(balancas_import_checked) > 0 && somente_preco {
// 					progress += 25
// 					fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 					flusher.Flush()
// 					time.Sleep(time.Millisecond * 500)
// 				}
// 				// balancas_import_checked, balancas_import_error = file.EnviaParaBalancas(itensFile, receitaFile, nutriFile, fracionaFile, fornFile, taraFile, conservaFile, campoextra, balancas_import, somente_preco)
// 			}

// 			// // Envio de evento de progresso
// 			// progress = progress + (60 * (len(balancas_import) - len(balancas_import_error)) / len(balancas_import))

// 			fmt.Fprintf(w, "data: {\"progress\": %d}\n\n", progress)
// 			flusher.Flush()

// 			// Envio de evento de conclusão com balanças importadas
// 			balancasImportCheckedJSON, err := json.Marshal(balancas_import_checked)
// 			if err != nil {
// 				fmt.Fprintf(w, "data: {\"error\": \"%s\"}\n\n", err.Error())
// 				flusher.Flush()
// 				return
// 			}

// 			// Verifica se houve balanças não importadas e envia os erros
// 			var errorJSON []byte
// 			if len(balancas_import_checked) < len(balancas_import) {
// 				errorJSON, err = json.Marshal(balancas_import_error)
// 				if err != nil {
// 					fmt.Fprintf(w, "data: {\"error\": \"%s\"}\n\n", err.Error())
// 					flusher.Flush()
// 					return
// 				}
// 			}

// 			time.Sleep(time.Millisecond * 500)

// 			fmt.Fprintf(w, "data: {\"incomplete\": %s}\n\n", errorJSON)
// 			flusher.Flush()

// 			time.Sleep(time.Millisecond * 500)

// 			fmt.Fprintf(w, "data: {\"complete\": %s}\n\n", balancasImportCheckedJSON)
// 			flusher.Flush()

// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}
// }
