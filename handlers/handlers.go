package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func TestRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Ping")
}

func Search(w http.ResponseWriter, r *http.Request) {
	search_input := r.URL.Query().Get("query")
	fmt.Println("GET query parameters were:", search_input)

	if len(search_input) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request. Query length is too long.\n%d", len(search_input))
		return
	}
	if search_input == "" {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, len(search_input))
		return
	}
	if search_input != "" {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, search_input)
		return
	}
}

func Package(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Package")
}

func GET_Package_List(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://sources.debian.org/api/list")
	if err != nil {
		log.Fatalln(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
	fmt.Fprint(w, data)
}
