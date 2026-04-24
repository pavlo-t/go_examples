package main

import (
	"os"
	"text/template"
)

// text/template package offers built-in support for creating dynamic content or showing customized output.
// html/template provides the same API but has additional security features and should be used for generating HTML.
//
//goland:noinspection GoUnhandledErrorResult
func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// we can use the template.Must function to panic in case Parse returns an error
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})

	// helper to build templates
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// if the data is a struct we can use the {{.FieldName}} action
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct{ Name string }{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse"})

	// if/else provide conditional execution for templates.
	// A value is considered false if it’s the default value of a type, such as 0, an empty string, nil pointer, etc.
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// range blocks let us loop through slices, arrays, maps or channels
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
}
