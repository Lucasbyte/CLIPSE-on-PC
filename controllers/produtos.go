package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/lucasbyte/go-clipse/db"
	"github.com/lucasbyte/go-clipse/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	err := temp.ExecuteTemplate(w, "Index", todosOsProdutos)
	if err != nil {
		db.ConectDb()
		todosOsProdutos := models.BuscaTodosOsProdutos()
		temp.ExecuteTemplate(w, "Index", todosOsProdutos)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Update", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Delete", todosOsProdutos)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		formCod := r.FormValue("codigo")
		descricao := strings.ToUpper(r.FormValue("descricao"))
		formPreco := r.FormValue("preco")
		formVenda := r.FormValue("venda-select")
		formValidade := r.FormValue("validade")

		preco, err := strconv.ParseFloat(formPreco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		venda, err := strconv.Atoi(formVenda)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		} else {
			fmt.Println(venda)
		}

		codigo, err := strconv.Atoi(formCod)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		validade, err := strconv.Atoi(formValidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		models.CriaNovoProduto(descricao, preco, codigo, venda, validade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		formCod := r.FormValue("codigo")
		descricao := strings.ToUpper(r.FormValue("descricao"))
		formPreco := r.FormValue("preco")
		formVenda := r.FormValue("venda-select")
		formValidade := r.FormValue("validade")

		preco, err := strconv.ParseFloat(formPreco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		venda, err := strconv.Atoi(formVenda)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		} else {
			fmt.Println(venda)
		}

		codigo, err := strconv.Atoi(formCod)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		validade, err := strconv.Atoi(formValidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		pluExist, err := models.ExisteProduto(codigo)
		if err != nil {
			fmt.Println(err)
		}

		if pluExist {
			models.EditProduct(descricao, preco, codigo, venda, validade, "user")
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Drop(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Teste")
		for i := 0; i <= 200; i++ {
			formName := fmt.Sprintf("plu%d", i)
			formCod := r.FormValue(formName)
			if formCod != "" {
				codigo, err := strconv.Atoi(formCod)
				if err != nil {
					log.Println("Erro na conversão do quantidade:", err)
				}
				fmt.Println(formCod)
				models.DeletProduct(codigo)
			} else {
				fmt.Println("Nada: ", formName)
			}
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func EnviarDados(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		models.ObterCodigosFaltantes()
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
