package database

type List struct {
	Packages []struct {
		Name string `json:"name"`
	} `json:"packages"`
}

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

type Package struct {
	Package  string     `json:"package"`
	Path     string     `json:"path"`
	Pathl    [][]string `json:"pathl"`
	Suite    string     `json:"suite"`
	Type     string     `json:"type"`
	Versions []struct {
		Area    string   `json:"area"`
		Suites  []string `json:"suites"`
		Version string   `json:"version"`
	} `json:"versions"`
}

// type PackageStats struct {
// 	Rank       int    `json:"rank"`
// 	Name       string `json:"name"`
// 	Inst       int    `json:"inst"`
// 	Vote       int    `json:"vote"`
// 	Old        int    `json:"old"`
// 	Recent     int    `json:"recent"`
// 	NoFiles    int    `json:"no-files"`
// 	Maintainer string `json:"maintainer"`
// }

type PackageStats struct {
	Rank       int
	Name       string
	Inst       int
	Vote       int
	Old        int
	Recent     int
	NoFiles    int
	Maintainer string
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
