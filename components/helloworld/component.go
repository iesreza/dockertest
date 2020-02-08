package helloworld

import (
	"github.com/iesreza/foundation/system"
	"html/template"
	"net/http"
)

type component struct {
	Path      string
	Assets    string
	Views     string
	Templates *template.Template
}

func (component component) GetTemplates() *template.Template {
	return component.Templates
}

func (component component) ViewPath() string {
	return component.Views
}

func (component component) AssetsPath() string {
	return component.Assets
}

var Component = component{
	Path:   "components/helloworld/",
	Assets: "components/helloworld/assets/",
	Views:  "components/helloworld/views/",
}

func Register() {
	Component.Register()

}

func (component *component) Register() {
	system.Components["helloworld"] = component

}

func (component *component) Menu() {

}

func (component *component) Install() {
	panic("implement me")
}

func (component *component) Uninstall() {
	panic("implement me")
}

func (component *component) Update() {
	panic("implement me")
}

func (component *component) ComputeHash() {
	panic("implement me")
}


func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}
