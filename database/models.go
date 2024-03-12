package database

type SearchResults struct {
	Query   string `json:"query"`
	Results struct {
		Exact interface{} `json:"exact"`
		Other []struct {
			Name string `json:"name"`
		} `json:"other"`
	} `json:"results"`
	Suite string `json:"suite"`
}

type List struct {
	Packages []struct {
		Name string `json:"name"`
	} `json:"packages"`
}

// func (data List) Value() (driver.Value, error) {
// 	return json.Marshal(data)
// }

// func (list *List) Scan(value interface{}) error {
// 	encoded_json, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}

// 	return json.Unmarshal(encoded_json, &list)
// }
