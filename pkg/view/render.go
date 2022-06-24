package view

import "net/http"

func (v *View) Render(w http.ResponseWriter, d Data) error {
	return v.Template.ExecuteTemplate(w, v.Bootstrap, d)
}
