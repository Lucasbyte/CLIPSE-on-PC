package models

import (
	"strings"

	"github.com/lucasbyte/go-clipse/file"
)

type File struct {
	Tipo    string
	Caminho string
}

func (f File) MudaTipo(tipo string) {
	f.Tipo = tipo
}

func NewFile() File {
	caminho := file.FinderFile()
	file := File{Caminho: caminho}
	return file
}

func (f File) LerArquivoDados() {
	caminho := strings.Replace(f.Caminho, "\\", "/", -1)
	if f.Tipo == "CSV" {
		Csv(caminho)
	} else if f.Tipo == "Txitens" {
		Txitens(caminho)
	} else if f.Tipo == "ItensMGV" {
		ItensMGV(caminho)
	}
}
