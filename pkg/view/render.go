package view

import (
	"net/http"

	"github.com/deemount/shobber-tpl/pkg/models"
)

func (v *View) Render(w http.ResponseWriter, d models.Data) error {
	return v.Template.ExecuteTemplate(w, v.Bootstrap, d)
}
