package generator

import (
	"os"
	"text/template"
)

func renderTemplate(templatePath, destinationPath, serviceName string) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(destinationPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := map[string]string{
		"ServiceName": serviceName,
	}

	tmpl.Execute(f, data)
}
