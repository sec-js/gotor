package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/KingAkeem/goTor/server/gobot"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func logErrMsg(w io.Writer, err error) {
	errMsg := fmt.Sprintf("Unable to get links. Error: %+v", err)
	log.Print(errMsg)
	fmt.Fprint(w, errMsg)
}

func getLinksHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	for {
		link := r.URL.Query().Get("link")
		links, err := gobot.GetLinks(link)
		if err != nil {
			logErrMsg(w, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(links)
		if err != nil {
			logErrMsg(w, err)
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
