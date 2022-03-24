package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readHttpResponseBody(res *http.Response) string {
	bytes, err := ioutil.ReadAll(res.Body)
	checkNilError(err)

	return string(bytes)
}

func main() {
	fmt.Println("-------------- Go HTTP Request --------------")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	checkNilError(err)
	defer response.Body.Close()

	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Headers:", response.Header)

	fmt.Println("Response Body:", readHttpResponseBody(response))
}
