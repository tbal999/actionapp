package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/tbal999/jsonbase"
)

const testUser = "INPUT_TEST_USERNAME"
const testAPI = "INPUT_TEST_API_KEY"
const testSpace = "INPUT_TEST_SPACE"

// iterate through a bunch of files - filepath fed via argument 1
func iterfiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println("Complete")
}

// checkConfluenceEnv is a placeholder function for checking the required env vars are set
func checkConfluenceEnv() {
	username, exists := os.LookupEnv(testUser)
	if !exists {
		log.Printf("Environment variable not set for %s", testUser)
	} else {
		log.Printf("USER: %s", username)
	}

	apiKey, exists := os.LookupEnv(testAPI)
	if !exists {
		log.Printf("Environment variable not set for %s", testAPI)
	} else {
		log.Printf("API KEY: %s", apiKey)
	}

	space, exists := os.LookupEnv(testSpace)
	if !exists {
		log.Printf("Environment variable not set for %s", testSpace)
	} else {
		log.Printf("SPACE: %s", space)
	}
}

func init() {
	checkConfluenceEnv()	
}

func main() {
	commands := os.Args[2:]
	fmt.Println("You have taken in these args:")
	for index := range commands {
		fmt.Println(commands[index])
		if commands[index] == "asdFoj230203ufdfsSDud123" { // random string to be used as secret
			fmt.Println("Secret received!") // https://docs.github.com/en/rest/reference/actions#secrets
		}
	}
	iterfiles(os.Args[1])
	fmt.Println(jsonbase.Temptable) // just to demonstrate you can run go app with external packages.
}
