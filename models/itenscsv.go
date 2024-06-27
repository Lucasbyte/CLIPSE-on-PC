package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Csv(caminho string) {
	//arquivo, err := os.Open("C:/Users/MAXWELL/Documents/DEV/Clipse/itensCSV.csv")
	arquivo, err := os.Open(caminho)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer arquivo.Close()

	reader := csv.NewReader(arquivo)
	reader.Comma = ';'

	linhas, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo CSV:", err)
		return
	}

	produtos := make([]Produto, 0)

	for _, linha := range linhas {
		if len(linha) != 5 {
			fmt.Println("Formato CSV inválido:", linha)
			continue
		}

		id, _ := strconv.Atoi(linha[0])
		plu, _ := strconv.Atoi(linha[0])
		preco, _ := strconv.ParseFloat(linha[2], 64)
		//descricao := linha[1]
		venda, _ := strconv.Atoi(linha[3])
		validade, _ := strconv.Atoi(linha[4])

		produto := Produto{
			Id:        id,
			Plu:       plu,
			Descricao: linha[1],
			Preco:     preco,
			Venda:     venda,
			Validade:  validade,
		}

		produtos = append(produtos, produto)
	}

	for _, p := range produtos {
		//fmt.Printf("%+v\n", p)

		existe, err := ExisteProduto(p.Plu)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(existe)
		if existe {
			EditProduct(p.Descricao, p.Preco, p.Plu, p.Venda, p.Validade, "import")
		} else {
			fmt.Println("Teste")
			fmt.Println(p.Descricao, p.Preco, p.Plu, p.Venda, p.Validade)
			CriaNovoProduto(p.Descricao, p.Preco, p.Plu, p.Venda, p.Validade)
		}
		existe, _ = ExisteProduto(p.Plu)
		fmt.Println(existe)
	}
}
