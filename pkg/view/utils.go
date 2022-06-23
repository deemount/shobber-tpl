package view

import (
	"html/template"
	"path/filepath"
)

func getTemplates(dir string) []string {
	tmpls, err := filepath.Glob(dir + "/*.tmpl")
	if err != nil {
		panic(err)
	}
	return tmpls
}

func appendTemplates(dir string, tmpls ...string) []string {
	return append(getTemplates(dir), tmpls...)
}

func parseTemplates(tmpls ...string) (*template.Template, error) {
	return template.ParseFiles(tmpls...)
}
