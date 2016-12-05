package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const version = "0.0.1"
const templateDir = "templates/*"

// Must be able to compile all template files.
var templates = template.Must(template.ParseGlob(templateDir))

func main() {
	//	http.HandleFunc("/api/register", registerHandler) // To handle all new application loads

	http.HandleFunc("/", HomeHandler)              // Display landing page... eventually.
	http.HandleFunc("/resources/", includeHandler) // Loads css/js/etc. straight through.

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//	log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	log.Fatal(srv.ListenAndServe())
}

// TODO: Check for those already logged in and forward to /admin
// Loads the login page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%q\n", r)
	err := templates.ExecuteTemplate(w, "main", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// For resource files like js, images, etc.
// Just a straight through file server.
func includeHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[1:]
	http.ServeFile(w, r, filename)
}
