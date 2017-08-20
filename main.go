package main

import (
	"os"
	"fmt"
)

func main() {
	if (len(os.Args) < 2) {
		fmt.Println("Must provide absolute path to target file/directory")
		os.Exit(1)
	}

	target := os.Args[1]

	current, err := os.Getwd()

	if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	fmt.Println(target)
	fmt.Println(current)

}