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

// iterate through a bunch of files - filepath fed via argument 1
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

// checkConfluenceEnv is a placeholder function for checking the required env vars are set
func checkConfluenceEnv() bool {
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
	if ok := checkConfluenceEnv(); ok {
		iterfiles(os.Args[1])
	}

	fmt.Println(jsonbase.Temptable)
}
