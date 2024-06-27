package controllers

import "html/template"

// SetTemplates configura o template compilado a partir do embed.FS
func SetTemplates(t *template.Template) {
	temp = t
}
