package view

import "html/template"

type View struct {
	Template  *template.Template
	Bootstrap string
}

func NewView(dir string, bootstrap string, tmpls ...string) *View {
	tmpls = appendTemplates(dir, tmpls...)
	t, err := parseTemplates(tmpls...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template:  t,
		Bootstrap: bootstrap,
	}
}
