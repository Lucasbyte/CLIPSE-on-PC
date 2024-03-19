package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createDatabaseIfNotExists(db *sql.DB) error {
	// Verifica se o arquivo do banco de dados já existe
	_, err := os.Stat(".db")
	if os.IsNotExist(err) {
		// Se o arquivo não existir, cria o banco de dados e a tabela
		_, err := db.Exec(`
            CREATE TABLE IF NOT EXISTS produtos (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
				plu INTEGER UNIQUE CHECK(plu <= 200),
            	descricao TEXT CHECK(length(descricao) <= 13),
                validade INTEGER CHECK(validade <= 200),
				preco DOUBLE CHECK(preco < 1000)
            );
        `)
		if err != nil {
			return err
		}
		fmt.Println("Banco de dados e tabela criados com sucesso.")
	} else if err != nil {
		// Em caso de erro ao verificar a existência do arquivo
		return err
	}

	return nil
}

func init() {
	// Abre a conexão com o banco de dados
	db, err := sql.Open("sqlite3", "andine.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Chama a função para criar o banco se não existir
	err = createDatabaseIfNotExists(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Resto do código...
}

func ConectDb() *sql.DB {
	db, err := sql.Open("sqlite3", "andine.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
