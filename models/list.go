package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type List struct {
	Packages []struct {
		Name string `json:"name"`
	} `json:"packages"`
}

// Make the List struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (data List) Value() (driver.Value, error) {
	return json.Marshal(data)
}

// Make the List struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (list *List) Scan(value interface{}) error {
	encoded_json, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(encoded_json, &list)
}
