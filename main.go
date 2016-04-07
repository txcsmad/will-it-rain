package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	austinLatitude  = 30.2672
	austinLongitude = -97.7431
)

var (
	indexTemplateStr = `
    <html>
        <head>
            <title>Will it rain in Austin today?</title>
        <head>
        <body>
            <center>
            <p>Will it rain in Austin today?</p>
            <h1>
            {{ if .Rain }}
                Yes!
            {{ else }}
                Nope.
            {{ end }}
            </h1>
            </center>
        </body>
    </html>
    `
	indexTemplate *template.Template
	forecastIOKey string
)

func init() {
	indexTemplate = template.Must(template.New("index").Parse(indexTemplateStr))
	forecastIOKey = os.Getenv("FORECASTIO_API_KEY")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	f, err := getForecast()

	if err != nil {
		http.Error(w, "error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	indexTemplate.Execute(w, map[string]bool{
		"Rain": f.Daily.Data[0].PrecipitationProbability > 0.5,
	})
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
