package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	gobot "github.com/KingAkeem/goTor/server/goBot"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getLinksHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "GET" {
		website := r.URL.Query().Get("url")
		links, err := gobot.GetLinks(website)
		if err != nil {
			log.Printf("Unable to retrieve links for %s. Error: %v", website, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&links)
		if err != nil {
			log.Printf("Unable to write response. Error: %+v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", getLinksHandler)
	fmt.Println("Serving on localhost:3050")
	log.Fatal(http.ListenAndServe(":3050", nil))
}
