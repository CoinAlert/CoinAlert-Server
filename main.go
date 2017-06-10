package main

import (
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
const url = "localhost"
const database = "coinalert"
const devices_collection = "devices"
const alerts_collection = "alerts"

func main() {
	fmt.Printf("Starting CoinAlert version %s\n", version)

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected to MongoDB\n")

	price := &Price{}
	price.Current, err = CurrentPrice()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Initial price for BTC: %s\n", price.Current)
	go priceUpdate(price)

	register := http.HandlerFunc(registerHandler(session))
	currentPrice := http.HandlerFunc(currentPriceHandler(price))
	alexa := http.HandlerFunc(alexaHandler(price))

	// API Routes
	http.Handle("/api/register", PostHandler(register))   // To handle all new application loads
	http.Handle("/api/current", GetHandler(currentPrice)) // Returns current price of BTC in USD
	http.Handle("/api/alexa", GetHandler(alexa))          // Returns current price of BTC in USD

	var templates = template.Must(template.ParseGlob(templateDir))

	// Web Routes
	http.HandleFunc("/", HomeHandler(price, templates)) // Display landing page... eventually.
	http.HandleFunc("/resources/", includeHandler)      // Loads css/js/etc. straight through.

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
	}

	//	log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))
	log.Fatal(srv.ListenAndServe())
}

// Return the home page.
func HomeHandler(price *Price, templates *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%q\n", r)
		err := templates.ExecuteTemplate(w, "main", price)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

}

// For resource files like js, images, etc.
// Just a straight through file server.
func includeHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[1:]
	http.ServeFile(w, r, filename)
}

func priceUpdate(price *Price) {
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
