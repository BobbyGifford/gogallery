package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   = "views/layouts/"
	TemplateExt = ".gohtml"
	TemplateDir = "views/"
)

func NewView(layout string, files ...string) *View {
	addTemplateExt(files)
	addTemplatePath(files)

	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func addTemplatePath(files []string) {
	for i, v := range files {
		files[i] = TemplateDir + v
	}
}

func addTemplateExt(files []string) {
	for i, v := range files {
		files[i] = v + TemplateExt
	}
}
