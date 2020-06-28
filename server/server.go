package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KingAkeem/goTor/server/gobot"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func logErrMsg(err error) {
	errMsg := fmt.Sprintf("Unable to get links. Error: %+v", err)
	log.Print(errMsg)
}

func getLinksHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "GET" {
		link := r.URL.Query().Get("link")
		links, err := gobot.GetLinks(link)
		if err != nil {
			logErrMsg(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(&links)
		if err != nil {
			logErrMsg(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}

func main() {
	http.HandleFunc("/", getLinksHandler)
	fmt.Println("Serving on localhost:3050")
	log.Fatal(http.ListenAndServe(":3050", nil))
}
