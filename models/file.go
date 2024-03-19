package models

import "github.com/lucasbyte/go-clipse/file"

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
