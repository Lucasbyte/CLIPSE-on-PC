package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CAD(caminho string) {
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
		if len(linha) < 6 {
			fmt.Println("Linha do arquivo TXT inválida:", linha)
			continue
		}
		char_venda := linha[7:8]
		linha_codigo := linha[0:6]
		linha_descricao := linha[7:22]
		linha_validade := linha[36:39]
		linha_preco := linha[29:36]
		fmt.Println(linha_codigo, linha_descricao, linha_preco, linha_validade)
		venda := 0

		if char_venda == "U" || char_venda == "u" {
			venda = 1
		}
		fmt.Println(venda, char_venda)
		plu, _ := strconv.Atoi(linha_codigo)

		precoStr := linha_preco
		preco, _ := strconv.ParseFloat(precoStr, 64)
		preco = preco / 100
		validade, _ := strconv.Atoi(linha_validade)
		descricao := strings.TrimSpace(linha_descricao)

		produto := Produto{
			Plu:       plu,
			Descricao: descricao,
			Preco:     preco,
			Venda:     venda,
			Validade:  validade,
		}
		fmt.Println(produto)
		produtos = append(produtos, produto)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo TXT:", err)
		return
	}

	fmt.Println(produtos)

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
