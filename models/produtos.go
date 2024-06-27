package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lucasbyte/go-clipse/db"
)

type Produto struct {
	Id        int
	Plu       int
	Descricao string
	Preco     float64
	Venda     int
	Validade  int
	CreatedAt time.Time
	UpdatedAt time.Time
	UpdatedBy string
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

func ComparaDB(plu int, desc string, preco float64, venda int, validade int) (bool, error) {
	db := db.ConectDb()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM produtos WHERE plu = ? AND descricao = ? AND venda = ? AND validade = ? AND preco = ?", plu, desc, venda, validade, preco).Scan(&count)
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

		var user string

		var updatedAt time.Time
		var createdAt time.Time

		err = selectDeTodosOsProdutos.Scan(&id, &plu, &descricao, &venda, &validade, &preco, &createdAt, &updatedAt, &user)
		if err != nil {
			fmt.Println(err.Error())
		}

		p.Plu = plu
		p.Descricao = descricao
		p.Preco = preco
		p.Venda = venda
		p.Validade = validade
		p.CreatedAt = createdAt
		p.UpdatedAt = updatedAt
		p.UpdatedBy = user

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(descricao string, preco float64, plu, venda, validade int) {
	db := db.ConectDb()

	descricao = strings.ToUpper(descricao)

	user := "import"

	insereDadosNoBanco, err := db.Prepare("insert into produtos(plu, descricao, preco, venda, validade, createdAt, updatedAt, updateBy) values($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := insereDadosNoBanco.Exec(plu, descricao, preco, venda, validade, time.Now(), time.Now(), user)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
	defer db.Close()

}

func EditProduct(descricao string, preco float64, plu, venda, validade int, user string) {
	db := db.ConectDb()
	query := "UPDATE produtos SET descricao = ?, preco = ?, venda = ?, validade = ?, updatedAt = ?, updateBy = ? WHERE plu = ?"
	descricao = strings.ToUpper(descricao)

	insereDadosNoBanco, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	insereDadosNoBanco.Exec(descricao, preco, venda, validade, time.Now(), user, plu)
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

func ObterCodigosFaltantes() ([]string, error) {
	// Consulta os valores da coluna "plu"
	db := db.ConectDb()

	rows, err := db.Query("SELECT plu FROM produtos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Cria um mapa para armazenar os códigos presentes na tabela
	codigos := make(map[int]bool)
	for rows.Next() {
		var codigo int
		if err := rows.Scan(&codigo); err != nil {
			return nil, err
		}
		codigos[codigo] = true
	}

	// Cria uma slice para armazenar os códigos faltantes
	var codigosFaltantes []string
	for i := 1; i <= 200; i++ {
		if !codigos[i] {
			codigo := fmt.Sprintf("%03d", i)
			codigosFaltantes = append(codigosFaltantes, codigo)
		}
	}
	return codigosFaltantes, nil
}

func ObterProduto(id int) (*Produto, error) {
	db := db.ConectDb()
	// Consulta os dados das colunas "plu", "descricao", "preco", "venda" e "validade" para um único produto com base no ID
	query := "SELECT * FROM produtos WHERE plu = ? LIMIT 1"
	row := db.QueryRow(query, id)

	var produto Produto
	// err = selectDeTodosOsProdutos.Scan(&id, &plu, &descricao, &venda, &validade, &preco, &createdAt, &updatedAt, &user)

	err := row.Scan(&produto.Id, &produto.Plu, &produto.Descricao, &produto.Venda, &produto.Validade, &produto.Preco, &produto.CreatedAt, &produto.UpdatedAt, &produto.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("produto não encontrado")
		}
		return nil, err
	}

	return &produto, nil
}

func ObterDadosProdutos() ([]string, error) {
	db := db.ConectDb()
	// Consulta os dados das colunas "plu", "descricao", "preco", "venda" e "validade"
	rows, err := db.Query("SELECT plu, descricao, preco, venda, validade FROM produtos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []string
	for rows.Next() {
		var produto Produto
		if err := rows.Scan(&produto.Plu, &produto.Descricao, &produto.Preco, &produto.Venda, &produto.Validade); err != nil {
			return nil, err
		}
		//fmt.Sprintf("\x03%s%s%s%s\x20", plu, desc, valor, venda)
		//precoStr := strings.Replace(fmt.Sprintf("%06.2f", produto.Preco), ".", "", -1)
		precoStr := transformarPreco(produto.Preco)

		//fmt.Println(precoStr)
		strVenda := "\x10"
		if produto.Venda == 1 {
			fmt.Println("un")
			strVenda = "\x20"
		}
		descricao := preencherDescricao(produto.Descricao, 15)
		produtoStr := fmt.Sprintf("\x03%03d%s00%s%03d%s", produto.Plu, strings.ToUpper(descricao), precoStr, produto.Validade, strVenda)
		produtoEcodigo := produtoStr + "£" + fmt.Sprint(produto.Plu)
		produtos = append(produtos, produtoEcodigo)
	}

	return produtos, nil
}

func transformarPreco(preco float64) string {
	//fmt.Println(preco)
	precoStr := fmt.Sprintf("%.2f", preco)
	//fmt.Println(precoStr)
	//log.Fatal("TESTE")
	precoStrSemPonto := strings.Replace(precoStr, ".", "", -1)
	precoComZeros := fmt.Sprintf("%06s", precoStrSemPonto)
	return precoComZeros
}

func preencherDescricao(descricao string, tamanho int) string {
	descricaoPreenchida := descricao
	if len(descricaoPreenchida) < tamanho {
		descricaoPreenchida += strings.Repeat(" ", tamanho-len(descricaoPreenchida))
	}
	return strings.ToUpper(descricaoPreenchida)
}
