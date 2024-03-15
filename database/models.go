package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// http:/ip_address/5432
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

// https://sources.debian.org/api/src/package_name/package_version
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

func (data Package) Value() (driver.Value, error) {
	return json.Marshal(data)
}

func (data *Package) Scan(value interface{}) error {
	encoded_json, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(encoded_json, &data)
}
