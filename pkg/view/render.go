package view

import (
	"net/http"

	"github.com/deemount/shobber-tpl/pkg/models"
)

func (v *View) RenderInterface(w http.ResponseWriter, d interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return v.Template.ExecuteTemplate(w, v.Bootstrap, d)
}

func (v *View) RenderData(w http.ResponseWriter, d models.Data) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return v.Template.ExecuteTemplate(w, v.Bootstrap, d)
}

func (v *View) RenderNil(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return v.Template.ExecuteTemplate(w, v.Bootstrap, nil)
}
