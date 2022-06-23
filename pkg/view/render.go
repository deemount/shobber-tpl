package view

import "net/http"

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	d := Data{
		Data: data,
	}
	return v.Template.ExecuteTemplate(w, v.Bootstrap, d)
}
