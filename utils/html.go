package utils

import (
	"bytes"
	"html/template"
)

func ParsedHTMLTemplate(templatePath string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var tplOutput bytes.Buffer

	err = tmpl.Execute(&tplOutput, data)
	if err != nil {
		return "", err
	}

	return tplOutput.String(), nil
}
