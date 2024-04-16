package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/lucasbyte/go-clipse/db"
	"github.com/lucasbyte/go-clipse/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

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

func File(w http.ResponseWriter, r *http.Request) {
	file := models.NewFile()
	temp.ExecuteTemplate(w, "File", file)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		formCod := r.FormValue("codigo")
		descricao := r.FormValue("descricao")
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
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		formCod := r.FormValue("codigo")
		descricao := r.FormValue("descricao")
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
			models.EditProduct(descricao, preco, codigo, venda, validade)
		}
	}
	http.Redirect(w, r, "/", 301)
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
	http.Redirect(w, r, "/", 301)
}

func Push(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Post file page: ")
	}
	http.Redirect(w, r, "/", 301)
}
