package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	checkNilError(err)

	return string(bytes)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	file, err := os.Create("./hello.txt")
	checkNilError(err)

	length, err := io.WriteString(file, "Hello World")
	checkNilError(err)

	fmt.Printf("%d bytes written\n", length)

	fileContent := readFile("./large-lorem.txt")
	fmt.Println(fileContent)

	file.Close()
}
