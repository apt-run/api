package handlers

import (
	"encoding/json"
	"fmt"
	"io"
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

	search_limit := 0
	var err error
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
		w.Write(database.ReadPaginate(search_limit))
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

func GET_PACKAGE(packag string) []byte {
	response, err := http.Get("https://sources.debian.org/src/" + packag)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

// // Initialize a new Attrs struct and add some values.
// attrs := new(Attrs)
// attrs.Name = "Pesto"
// attrs.Ingredients = []string{"Basil", "Garlic", "Parmesan", "Pine nuts", "Olive oil"}
// attrs.Organic = false
// attrs.Dimensions.Weight = 100.00

// // The database driver will call the Value() method and and marshall the
// // attrs struct to JSON before the INSERT.
// _, err = db.Exec("INSERT INTO items (attrs) VALUES($1)", attrs)
// if err != nil {
//     log.Fatal(err)
// }

// // Similarly, we can also fetch data from the database, and the driver
// // will call the Scan() method to unmarshal the data to an Attr struct.
// item := new(Item)
// err = db.QueryRow("SELECT id, attrs FROM items ORDER BY id DESC LIMIT 1").Scan(&item.ID, &item.Attrs)
// if err != nil {
//     log.Fatal(err)
// }

// // You can then use the struct fields as normal...
// weightKg := item.Attrs.Dimensions.Weight / 1000
// log.Printf("Item: %d, Name: %s, Weight: %.2fkg", item.ID, item.Attrs.Name, weightKg)
