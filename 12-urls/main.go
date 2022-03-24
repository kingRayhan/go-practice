package main

import (
	"fmt"
	"net/url"
)

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

const endpoint string = "https://api.techdiary.dev/api/articles?tag=js&tag=go"

func main() {
	fmt.Println("-------------- Go URL Module --------------")

	// parse url
	result, err := url.Parse(endpoint)
	checkNilError(err)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Fragment:", result.Fragment)

	query := result.Query()
	fmt.Println("Query:", query)
	fmt.Println("query['tag']:", query["tag"])
	fmt.Println("query['tag'][0]:", query["tag"][0])
	fmt.Println("query['tag'][1]:", query["tag"][1])

	// build url
	customMadeUrl := &url.URL{
		Scheme:   "https",
		Host:     "api.techdiary.dev",
		Path:     "/api/articles",
		RawQuery: "tag=js&tag=go&limit=10",
	}
	fmt.Println("Custom Made URL:", customMadeUrl.String())
}
