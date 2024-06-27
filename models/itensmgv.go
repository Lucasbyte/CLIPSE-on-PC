package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ItensMGV(caminho string) {
	//arquivo, err := os.Open("C:\\Users\\MAXWELL\\Desktop\\TEST\\itensSystel.TXT")
	arquivo, err := os.Open(caminho)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer arquivo.Close()

	scanner := bufio.NewScanner(arquivo)

	produtos := make([]Produto, 0)

	for scanner.Scan() {
		linha := scanner.Text()

		if len(linha) < 34 {
			fmt.Println("Linha do arquivo TXT inválida:", linha)
			continue
		}

		venda, _ := strconv.Atoi(string(linha[2]))
		plu, _ := strconv.Atoi(linha[3:9])

		precoStr := linha[9:15]
		preco, _ := strconv.ParseFloat(precoStr, 64)
		preco = preco / 100
		validade, _ := strconv.Atoi(linha[15:18])
		descricao := strings.TrimSpace(linha[18:33])
		fmt.Println(linha[2:9], linha[9:15], linha[15:18], linha[18:33])
		produto := Produto{
			Plu:       plu,
			Descricao: descricao,
			Preco:     preco,
			Venda:     venda,
			Validade:  validade,
		}

		produtos = append(produtos, produto)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo TXT:", err)
		return
	}

	for _, p := range produtos {
		if p.Plu <= 200 {
			existe, err := ExisteProduto(p.Plu)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(existe)

			if existe {
				Ehigual, err := ComparaDB(p.Plu, p.Descricao, p.Preco, p.Venda, p.Validade)
				if err != nil {
					fmt.Println(err)
					return
				}
				if Ehigual {
					continue
				}
				EditProduct(p.Descricao, p.Preco, p.Plu, p.Venda, p.Validade, "import")
			} else {
				CriaNovoProduto(p.Descricao, p.Preco, p.Plu, p.Venda, p.Validade)
			}

			existe, _ = ExisteProduto(p.Plu)
			//mt.Println(existe)
		}
	}
}
