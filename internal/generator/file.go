package generator

import (
	"os"
	"text/template"
)

func GenerateFileFromTemplate(tmplPath, outputPath string, data interface{}) error {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
