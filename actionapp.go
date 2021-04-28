package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tbal999/jsonbase"
)

const testAPI = "INPUT_TEST_API_KEY"
const testSpace = "INPUT_TEST_SPACE"

func iterfiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	for _, file := range files {
		log.Println(file.Name())
	}
	log.Println("Complete")
}

func checkEnv() bool {
	apiKey, exists := os.LookupEnv(testAPI)
	if !exists {
		log.Printf("Environment variable not set for %s", testAPI)
		return false
	} else {
		log.Printf("API KEY: %s", apiKey)
	}

	space, exists := os.LookupEnv(testSpace)
	if !exists {
		log.Printf("Environment variable not set for %s", testSpace)
		return false
	} else {
		log.Printf("SPACE: %s", space)
	}
	return true
}

func main() {
	if ok := checkEnv(); ok {
		iterfiles(os.Args[1])
	}

	fmt.Println(jsonbase.Temptable)
}
