package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

const rawTemplate = `type {{ .Name }} {{ .Type }}

func New{{ .Name }}(v {{ .Type }}) *{{ .Name }} {
	return (*{{ .Name }})(&v)
}

func (v *{{ .Name }}) Value() {{ .Type }} {
	if v == nil {
		var zero {{ .Type }}
		return zero
	}
	return {{ .Type }}(*v)
}

func (v *{{ .Name }}) {{ .Name }}() {{ .Type }} {
	return v.Value()
}

func (v *{{ .Name }}) IsNull() bool {
	return v == nil
}

func (v *{{ .Name }}) Ok() ({{ .Type }}, bool) {
	return v.Value(), v.IsNull()
}

`

type Data struct {
	Name string
	Type string
}

func main() {
	tpl, err := template.New("default").Parse(rawTemplate)
	if err != nil {
		log.Fatal(err)
	}

	types := []string{
		"bool",
		"string",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",

		"byte",

		"rune",

		"float32", "float64",

		"complex64", "complex128",
	}

	file, err := os.OpenFile("./generated.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("package nulls\n\n")

	for _, t := range types {
		err = tpl.Execute(file, Data{Name: titleCase(t), Type: t})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func titleCase(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}
