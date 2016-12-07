package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"html/template"
	"log"
	"net/http"
	"time"
)

const version = "0.0.1"
const templateDir = "templates/*"

// Must be able to compile all template files.
var templates = template.Must(template.ParseGlob(templateDir))
var url = "localhost"
var database = "coinalert"
var collection = "devices"

var db *mgo.Collection

type Price struct {
	Current string `json:"currentPrice"`
}

var price Price

func main() {
	fmt.Printf("Starting CoinAlert version %s\n", version)

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(database).C(collection)
	fmt.Printf("Connected to MongoDB\n")

	price.Current, err = CurrentPrice()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Initial price for BTC: %s\n", price.Current)
	go priceUpdate()

	http.HandleFunc("/api/register", registerHandler)    // To handle all new application loads
	http.HandleFunc("/api/current", currentPriceHandler) // Returns current price of BTC in USA$

	http.HandleFunc("/", HomeHandler)              // Display landing page... eventually.
	http.HandleFunc("/resources/", includeHandler) // Loads css/js/etc. straight through.

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//	log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%q\n", r)
	err := templates.ExecuteTemplate(w, "main", price)
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

func currentPriceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Add("Allowed", "GET")
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	buf, err := json.Marshal(price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", buf)
}

func priceUpdate() {
	ticker := time.NewTicker(5 * time.Second)

	// Keep trying until we're timed out or got a result or got an error
	for {
		select {
		// Got a timeout! fail with a timeout error
		case <-ticker.C:
			p, err := CurrentPrice()
			if err != nil {
				fmt.Printf("Error getting current price: %s\n", err.Error())

			} else {
				price.Current = p
				fmt.Printf("Price updated to $%s\n", p)
			}
		}
	}
}
