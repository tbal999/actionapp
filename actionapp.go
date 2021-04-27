package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/tbal999/jsonbase"
)

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
