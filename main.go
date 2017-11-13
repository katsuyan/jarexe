package main

import (
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	flag "github.com/spf13/pflag"
)

func main() {

	var (
		exeFileName = flag.String("name", "exejar_default", "exe file name")
	)

	flag.Parse()

	jarFileName := flag.Args()[0]
	jarFile, _ := os.Open(jarFileName)

	rep := regexp.MustCompile(`.*\/(.*).jar`)
	defaultExeFileName := rep.ReplaceAllString(jarFileName, "$1")

	if *exeFileName == "exejar_default" {
		*exeFileName = defaultExeFileName
	}

	shStr := `#!/bin/sh
	exec java -jar "$0" "$@"
	exit 1`

	reader := io.MultiReader(strings.NewReader(shStr), jarFile)

	b, _ := ioutil.ReadAll(reader)
	ioutil.WriteFile(*exeFileName, b, os.ModePerm)
}
