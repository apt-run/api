package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"net/http"
	"strconv"
)

func TestRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Ping")
}

func Search(w http.ResponseWriter, r *http.Request) {
	search_query := r.URL.Query().Get("query")
	str := r.URL.Query().Get("limit")
	fmt.Println(str)

	var err error
	search_limit := 0
	if str != "" {
		search_limit, err = strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
	}

	if len(search_query) > 50 || search_limit > 100 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if search_query == "" {
		w.WriteHeader(http.StatusAccepted)
		w.Write(database.ReadPaginate(20))
		return
	}
	if search_query != "" {
		w.WriteHeader(http.StatusAccepted)
		w.Write(database.ReadSearch(search_query))
		return
	}
}

func Package(w http.ResponseWriter, r *http.Request) {
	packag := r.URL.Query().Get("query")
	fmt.Println("GET query parameters were:", packag)

	if len(packag) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request. Query length is too long.\n%d", len(packag))
		return
	}
	if packag == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, len(packag))
		return
	}
	if packag != "" {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": 200,
			"source":     "debian",
			"package":    packag,
			"data":       "json",
		})
		return
	}
}
