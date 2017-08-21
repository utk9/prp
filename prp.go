package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func getRelativePath(current string, target string) string {
	targetSlice := strings.Split(strings.Trim(target, "/"), "/")
	currentSlice := strings.Split(strings.Trim(current, "/"), "/")

	branchIdx := 0

	var minLen int
	if len(currentSlice) < len(targetSlice) {
		minLen = len(currentSlice)
	} else {
		minLen = len(targetSlice)
	}

	for branchIdx < minLen {
		if currentSlice[branchIdx] != targetSlice[branchIdx] {
			break
		}
		branchIdx++
	}

	var buffer bytes.Buffer

	buffer.WriteString("./")

	for i := branchIdx; i < len(currentSlice); i++ {
		buffer.WriteString("../")
	}

	for i := branchIdx; i < len(targetSlice); i++ {
		buffer.WriteString(targetSlice[i])
		buffer.WriteString("/")
	}

	return strings.TrimRight(buffer.String(), "/")

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must provide absolute path to target file/directory")
		os.Exit(1)
	}

	var to string
	var from string

	if len(os.Args) < 3 {
		to = os.Args[1]
		var err error
		from, err = os.Getwd()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		from = os.Args[1]
		to = os.Args[2]
	}


	if from[0] != '/' || to[0] != '/' {
		fmt.Println("Must provide absolute path(s) (from root directory)")
		os.Exit(1)
	}

	relativePath := getRelativePath(from, to)

	fmt.Println(relativePath)
}
