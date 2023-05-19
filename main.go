package main

import (
	"fmt"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

// Create a new Jet set
var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./templates"),
	jet.InDevelopmentMode(),
)

// Load your template files
var tmpl, err = views.GetTemplate("index.jet.html")

func indexHandler(w http.ResponseWriter, r *http.Request) {
	type JetStruct struct {
		Title   string
		Heading string
		Body    string
	}
	data := JetStruct{
		Title:   "My Page",
		Heading: "Welcome to my page",
		Body:    "This is some example text.",
	}
	vars := make(jet.VarMap)
	vars.Set("Title", data.Title)
	vars.Set("Heading", data.Heading)
	vars.Set("Body", data.Body)

	// Parse the template file
	var err error
	err = tmpl.Execute(w, vars, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func main() {
	// Start the HTTP server
	http.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started on http://localhost:8080")
}
