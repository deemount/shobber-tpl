package view

import "html/template"

type View struct {
	Template  *template.Template
	Bootstrap string
}

func NewView(dir string, bootstrap string, tmpls ...string) *View {
	tmpls = AppendTemplates(dir, tmpls...)
	t, err := ParseTemplates(tmpls...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template:  t,
		Bootstrap: bootstrap,
	}
}
