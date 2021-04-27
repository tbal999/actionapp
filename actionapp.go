package main

import (
	"fmt"
	"os"
)

func main() {
	commands := os.Args[1:]
	fmt.Println("You have taken in these args:")
	fmt.Println(commands)
}
