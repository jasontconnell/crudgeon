package process

import (
	"bytes"
	"log"
	"text/template"
)

func processTemplate(tmplstring string, data any) string {
	tmpl, err := template.New("outputFormat").Parse(tmplstring)
	if err != nil {
		log.Println(err)
		return ""
	}
	buffer := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buffer, "outputFormat", data)
	return buffer.String()
}
