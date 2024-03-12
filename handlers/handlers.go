package handlers

import (
	"fmt"
	"log"
	"main/database"
	"net/http"
	"strconv"
)

func TestRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Ping")
}

func Search(w http.ResponseWriter, r *http.Request) {
	search_query := r.URL.Query().Get("query")
	limit := r.URL.Query().Get("limit")
	search_limit := 0
	var err error

	if limit != "" {
		search_limit, err = strconv.Atoi(limit)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(search_query) > 50 || search_limit > 100 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if search_query == "" {
		w.WriteHeader(http.StatusAccepted)
		w.Write(database.ReadList(search_limit))
		return
	}
	if search_query != "" {
		w.WriteHeader(http.StatusAccepted)
		w.Write(database.ReadSearch(search_query, search_limit))
		return
	}
}

func Package(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Ping")
}
