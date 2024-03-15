package handlers

import (
	"fmt"
	"main/database"
	"net/http"
)

func TestRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Ping")
}

func Search(w http.ResponseWriter, r *http.Request) {
	search_query := r.URL.Query().Get("query")

	if len(search_query) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if search_query == "" {
		w.WriteHeader(http.StatusAccepted)
		w.Write(database.ReadByName(search_query, 20))
		return
	}
	if search_query != "" {
		w.WriteHeader(http.StatusAccepted)
		w.Write(database.ReadByName(search_query, 20))
		return
	}
}

func Package(w http.ResponseWriter, r *http.Request) {
	packag := r.URL.Query().Get("package")
	// version := r.URL.Query().Get("version")

	if len(packag) > 50 || packag == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if packag != "" {
		w.WriteHeader(http.StatusAccepted)
		return
	}
}
