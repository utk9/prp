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

	target := os.Args[1]

	current, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	relativePath := getRelativePath(current, target)

	fmt.Println(relativePath)
}
