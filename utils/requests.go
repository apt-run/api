package utils

import (
	"io"
	"log"
	"net/http"
)

func GetPackage(packag string) []byte {
	response, err := http.Get("https://sources.debian.org/api/src/" + packag)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func GetList() []byte {
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

func GetSearchResult(search_string string) []byte {
	response, err := http.Get("https://sources.debian.org/api/search/" + search_string)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func GetPackageMetrics(packag string, version string) []byte {
	response, err := http.Get("https://sources.debian.org/api/info/package/" + packag + "/" + version)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}

func GetPopularityData(search_string string) []byte {
	response, err := http.Get("https://sources.debian.org/api/search/" + search_string)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
