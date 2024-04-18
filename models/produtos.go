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
	Venda     int
	Validade  int
}

func ExisteProduto(plu int) (bool, error) {
	db := db.ConectDb()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM produtos WHERE plu = ?", plu).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectDb()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos ORDER BY plu")
	if err != nil {
		fmt.Println(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, plu, validade, venda int
		var descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &plu, &descricao, &venda, &validade, &preco)
		if err != nil {
			fmt.Println(err.Error())
		}

		p.Plu = plu
		p.Descricao = descricao
		p.Preco = preco
		p.Venda = venda
		p.Validade = validade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
func CriaNovoProduto(descricao string, preco float64, plu, venda, validade int) {
	db := db.ConectDb()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(plu, descricao, preco, venda, validade) values($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := insereDadosNoBanco.Exec(plu, descricao, preco, venda, validade)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	defer db.Close()

}

func EditProduct(descricao string, preco float64, plu, venda, validade int) {
	db := db.ConectDb()
	query := "UPDATE produtos SET descricao = ?, preco = ?, venda = ?, validade = ? WHERE plu = ?"

	insereDadosNoBanco, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	insereDadosNoBanco.Exec(descricao, preco, venda, validade, plu)
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
