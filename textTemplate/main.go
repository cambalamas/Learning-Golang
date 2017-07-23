package main

import (
	"log"
	"os"
	"text/template"
)

const templ = `{{.Var1}} and {{.Var2}}`

// Data contains the info
type Data struct {
	Var1 int
	Var2 string
}

func main() {
	log.SetOutput(os.Stderr)

	data := Data{1, "hi"}
	report, err := template.New("report").Parse(templ)

	if err != nil {
		log.Fatalf("Could not parse the template: %q", err)
	}

	if err := report.Execute(os.Stdout, data); err != nil {
		log.Fatalf("Could not fill the template: %q", err)
	}
}
