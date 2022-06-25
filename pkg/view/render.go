package view

import (
	"net/http"

	"github.com/deemount/shobber-tpl/pkg/models"
)

type RenderInterface View
type RenderData View
type RenderNil View

func (vi *RenderInterface) Render(w http.ResponseWriter, d interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return vi.Template.ExecuteTemplate(w, vi.Bootstrap, d)
}

func (vd *RenderData) Render(w http.ResponseWriter, d models.Data) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return vd.Template.ExecuteTemplate(w, vd.Bootstrap, d)
}

func (vn *RenderNil) Render(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return vn.Template.ExecuteTemplate(w, vn.Bootstrap, nil)
}
