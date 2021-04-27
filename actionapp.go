package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	}
	iterfiles(os.Args[1])
}
