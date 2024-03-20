package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const URL = "https://popcon.debian.org/by_inst"

func main() {
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	counter := 0

	lines := strings.Split((string(body)), "\n")
	for _, line := range lines[11 : len(lines)-3] {

		if counter == 20 {
			break
		}
		counter++

		fields := strings.Fields(line)
		rank := fields[0]
		name := fields[1]
		installs := fields[2]
		split := strings.Split(line, ("("))
		last_element := split[len(split)-1]
		maintainer := strings.ReplaceAll(last_element, ")", "")

		fmt.Println(line)
		fmt.Println(rank)
		fmt.Println(name)
		fmt.Println(installs)
		fmt.Println(maintainer)
	}
}
