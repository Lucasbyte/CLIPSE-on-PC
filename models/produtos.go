package models

import (
	"fmt"

	"github.com/lucasbyte/go-clipse/db"
)

type Produto struct {
	Id        int
	Plu       int
	Descricao string
	Preco     float64
	Validade  int
}

func ExisteProduto(plu int) (bool, error) {
	db := db.ConectDb()
	query := fmt.Sprintf("select * from produtos where plu = %d", plu)
	fmt.Println(query)
	_, err := db.Query(query)
	if err != nil {
		return false, err
	}
	return true, nil
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectDb()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		fmt.Println(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, plu, validade int
		var descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &plu, &descricao, &preco, &validade)
		if err != nil {
			fmt.Println(err.Error())
		}

		p.Plu = plu
		p.Descricao = descricao
		p.Preco = preco
		p.Validade = validade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
func CriaNovoProduto(descricao string, preco float64, plu, validade int) {
	db := db.ConectDb()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(plu, descricao, preco, validade) values($1, $2, $3, $4)")
	if err != nil {
		fmt.Println(err.Error())
	}

	insereDadosNoBanco.Exec(plu, descricao, preco, validade)
	defer db.Close()

}

func EditProduct(descricao string, preco float64, plu, validade int) {
	db := db.ConectDb()
	query := "UPDATE produtos SET descricao = ?, preco = ?, validade = ? WHERE plu = ?"

	insereDadosNoBanco, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	insereDadosNoBanco.Exec(descricao, preco, validade, plu)
	defer db.Close()
}

func DeletProduct(plu int) {
	db := db.ConectDb()
	query := "DELETE FROM produtos WHERE plu = ?"

	insereDadosNoBanco, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	insereDadosNoBanco.Exec(plu)
	defer db.Close()
}
