package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	const jsonBody = `{ "message": "Hello Reader!" }`
	var jsonMap map[string]string
	r := strings.NewReader(jsonBody)
	json.NewDecoder(r).Decode(&jsonMap)

	fmt.Println(jsonMap)
}
