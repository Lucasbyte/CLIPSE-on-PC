package models

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Txitens(caminho string) {
	//arquivo, err := os.Open("C:\\Users\\MAXWELL\\Desktop\\ARQUIVOS DE ITENS\\txitens.txt")
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

		if len(linha) < 17 {
			fmt.Println("Linha do arquivo TXT inválida:", linha)
			continue
		}

		venda, _ := strconv.Atoi(string(linha[4]))
		plu, _ := strconv.Atoi(linha[5:11])

		precoStr := linha[11:17]
		preco, _ := strconv.ParseFloat(precoStr, 64)
		preco = preco / 100
		validadeStr := linha[17:20]
		validade, _ := strconv.Atoi(validadeStr)
		descricao := ""
		if len(linha) < 35 {
			descricao = strings.TrimSpace(linha[20:])
		} else {
			descricao = strings.TrimSpace(linha[20:35])
		}
		fmt.Println(linha[5:11], venda, precoStr, validadeStr, descricao)
		produto := Produto{
			Id:        plu, // Seu código para definir o ID do produto
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
				EditProduct(p.Descricao, p.Preco, p.Plu, p.Venda, p.Validade)
				fmt.Println(p.Plu)
			} else {
				CriaNovoProduto(p.Descricao, p.Preco, p.Plu, p.Venda, p.Validade)
			}

			existe, _ = ExisteProduto(p.Plu)
			//mt.Println(existe)
		}
	}
}
