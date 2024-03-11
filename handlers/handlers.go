package handlers

import (
	"encoding/json"
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
		w.Write(GET_PACKAGE_LIST())
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

func GET_PACKAGE_LIST() []byte {
	response, err := http.Get("https://sources.debian.org/api/list")
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
