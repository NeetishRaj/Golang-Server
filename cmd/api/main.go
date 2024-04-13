package main

import (
	"fmt"
	"html/template"
	"net/http"

	internal_api "cdk-converter/internal/api"
)

func main() {

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Server is live and active!")
	})

	// Static file hosting : web/
	fs := http.FileServer(http.Dir("web/public/"))
	http.Handle("/", fs)

	tmpl := template.Must(template.ParseFiles("web/templates/converter.html"))
	http.HandleFunc("/converter", func(res http.ResponseWriter, req *http.Request) {
		data := internal_api.ConverterDataPage{
			Title:    "Python CDK",
			Filename: "sample_file.zip",
		}
		tmpl.Execute(res, data)
	})

	http.ListenAndServe(":3003", nil)
}
