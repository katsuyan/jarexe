package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file1, _ := os.Open(os.Args[1])
	file2, _ := os.Open(os.Args[2])
	reader := io.MultiReader(file1, file2)

	b, _ := ioutil.ReadAll(reader)
	ioutil.WriteFile("runrun", b, os.ModePerm)
}
