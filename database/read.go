package database

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v5"
)

func ReadByMaintianer(maintainer string, limit int) []byte {
	var rows pgx.Rows
	rows, err := connection.Query(context.Background(),
		SELECT_MAINTAINERS,
		maintainer,
		limit,
	)
	if err != nil {
		log.Fatalf("Unable to query database: %v", err)
	}

	var jsonValue string
	var jsonData map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(&jsonValue); err != nil {
			log.Fatalf("Unable to scan json value: %v", err)
		}
		if err := json.Unmarshal([]byte(jsonValue), &jsonData); err != nil {
			log.Fatalf("Unable to unmarshal json value: %v", err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}

	return []byte(jsonValue)
}

func ReadByInstalls(limit int) []byte {
	var rows pgx.Rows
	rows, err := connection.Query(context.Background(),
		SELECT_INSTALLS,
		limit,
	)
	if err != nil {
		log.Fatalf("Unable to query database: %v", err)
	}

	var jsonValue string
	var jsonData map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(&jsonValue); err != nil {
			log.Fatalf("Unable to scan json value: %v", err)
		}
		if err := json.Unmarshal([]byte(jsonValue), &jsonData); err != nil {
			log.Fatalf("Unable to unmarshal json value: %v", err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}

	return []byte(jsonValue)
}

func ReadByName(name string, limit int) []byte {
	var rows pgx.Rows
	rows, err := connection.Query(context.Background(),
		SELECT_NAMES,
		name,
		limit,
	)
	if err != nil {
		log.Fatalf("Unable to query database: %v", err)
	}

	var jsonValue string
	var jsonData map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(&jsonValue); err != nil {
			log.Fatalf("Unable to scan json value: %v", err)
		}
		if err := json.Unmarshal([]byte(jsonValue), &jsonData); err != nil {
			log.Fatalf("Unable to unmarshal json value: %v", err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}

	return []byte(jsonValue)
}

// func ReadSourcesSearchList(search_string string, limit int) []byte {
// 	var err error
// 	var rows pgx.Rows
// 	if limit > 0 {
// 		rows, err = connection.Query(context.Background(),
// 			SEARCH_PACKAGES,
// 			search_string,
// 			limit,
// 		)
// 		if err != nil {
// 			log.Fatalf("Unable to query database: %v", err)
// 		}
// 	} else {
// 		rows, err = connection.Query(context.Background(),
// 			SEARCH_PACKAGES,
// 			search_string,
// 			20,
// 		)
// 		if err != nil {
// 			log.Fatalf("Unable to query database: %v", err)
// 		}
// 	}

// 	var jsonValue string
// 	var jsonData map[string]interface{}
// 	for rows.Next() {
// 		err = rows.Scan(&jsonValue)
// 		if err != nil {
// 			log.Fatalf("Unable to scan json value: %v", err)
// 		}
// 		if err := json.Unmarshal([]byte(jsonValue), &jsonData); err != nil {
// 			log.Fatalf("Unable to unmarshal json value: %v", err)
// 		}
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Fatalf("Error reading rows: %v", err)
// 	}

// 	return []byte(jsonValue)
// }

// func ReadSourcesList(index int) []byte {
// 	var err error
// 	var rows pgx.Rows
// 	if index < 1 {
// 		rows, err = connection.Query(context.Background(),
// 			PAGINATE_PACKAGES_START,
// 		)
// 		if err != nil {
// 			log.Fatalf("Unable to query database: %v", err)
// 		}
// 	} else {
// 		rows, err = connection.Query(context.Background(),
// 			PAGINATE_PACKAGES,
// 			index,
// 		)
// 		if err != nil {
// 			log.Fatalf("Unable to query database: %v", err)
// 		}
// 	}

// 	var jsonValue string
// 	var jsonData map[string]interface{}
// 	for rows.Next() {
// 		err = rows.Scan(&jsonValue)
// 		if err != nil {
// 			log.Fatalf("Unable to scan json value: %v", err)
// 		}
// 		if err := json.Unmarshal([]byte(jsonValue), &jsonData); err != nil {
// 			log.Fatalf("Unable to unmarshal json value: %v", err)
// 		}
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Fatalf("Error reading rows: %v", err)
// 	}

// 	return []byte(jsonValue)
// }
