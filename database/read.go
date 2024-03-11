package database

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v5"
)

// import
// cache map search_string to json
// cache map maintianer to json
// cache map package_string to json

// create metrics table
// count searches
// count maintainer views
// count package views
// count package copies

func ReadSearch(search_string string) []byte {
	rows, err := connection.Query(context.Background(),
		SEARCH_PACKAGES,
		search_string,
	)
	if err != nil {
		log.Fatalf("Unable to query database: %v", err)
	}

	var jsonValue string
	var jsonData map[string]interface{}
	for rows.Next() {
		err = rows.Scan(&jsonValue)
		if err != nil {
			log.Fatalf("Unable to scan json value: %v", err)
		}

		// Store the JSON value in a variable
		if err := json.Unmarshal([]byte(jsonValue), &jsonData); err != nil {
			log.Fatalf("Unable to unmarshal json value: %v", err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}

	return []byte(jsonValue)
}
func ReadPackage(search_string string)    {}
func ReadMaintainer(search_string string) {}

func ReadPaginate(index int) []byte {

	var rows pgx.Rows
	var err error
	if index < 1 {
		rows, err = connection.Query(context.Background(),
			PAGINATE_PACKAGES_START,
		)
		if err != nil {
			log.Fatalf("Unable to query database: %v", err)
		}
	} else {
		rows, err = connection.Query(context.Background(),
			PAGINATE_PACKAGES,
			index,
		)
		if err != nil {
			log.Fatalf("Unable to query database: %v", err)
		}
	}

	var jsonValue string
	var jsonData map[string]interface{}
	for rows.Next() {
		err = rows.Scan(&jsonValue)
		if err != nil {
			log.Fatalf("Unable to scan json value: %v", err)
		}

		// Store the JSON value in a variable
		if err := json.Unmarshal([]byte(jsonValue), &jsonData); err != nil {
			log.Fatalf("Unable to unmarshal json value: %v", err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}

	return []byte(jsonValue)
}

// func ReadUser(username string) string {
// 	row, err := connection.Query(context.Background(),
// 		SELECT_USER,
// 		username,
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for row.Next() {
// 		var id int32
// 		var result string
// 		var password string
// 		err = row.Scan(&id, &result, &password)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if result == username {
// 			return result
// 		}
// 	}
// 	return "error reading username."
// }

// func ReadPassword(username string) string {
// 	row, err := connection.Query(context.Background(),
// 		SELECT_USER,
// 		username,
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for row.Next() {
// 		var id int32
// 		var result string
// 		var password string
// 		err = row.Scan(&id, &result, &password)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return password
// 	}
// 	return "Read password error."
// }

// func UpdateUser(username string, new_username string, new_password string) {
// 	_, err := connection.Exec(context.Background(),
// 		UPDATE_USER,
// 		new_username,
// 		new_password,
// 		username,
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func DeleteUser(username string) {
// 	_, err := connection.Exec(context.Background(),
// 		DELETE_USER,
// 		username,
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
