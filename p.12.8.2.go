package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("log.txt")
	if err != nil {
		panic(err)

	}
	defer file.Close()

	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(resultBytes))
}
