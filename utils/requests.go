package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPackage(packag string) {
	response, err := http.Get("https://sources.debian.org/api/src/" + packag)
	if err != nil {
		log.Fatalln(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v", data)
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
	fmt.Printf("%v", data)
	return data
}
