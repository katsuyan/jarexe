package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	flag "github.com/spf13/pflag"
)

func main() {

	var (
		exeFileName = flag.String("name", "", "exe file name")
		javaOptions = flag.String("jop", "", "java options")
	)

	flag.Parse()

	if len(flag.Args()) == 0 {
		panic(errors.New("Not enough arguments. ex)jarexe standalone.jar"))
	}

	jarFileName := flag.Args()[0]
	jarFile, err := os.Open(jarFileName)
	if err != nil {
		log.Printf("error: %s", err.Error())
		panic(err.Error())
	}

	rep := regexp.MustCompile(`.*\/(.*).jar`)
	defaultExeFileName := rep.ReplaceAllString(jarFileName, "$1")

	if *exeFileName == "" {
		*exeFileName = defaultExeFileName
	}

	shStr := fmt.Sprintf(`#!/bin/sh
		exec java %s -jar "$0" "$@"
		exit $?`, *javaOptions)

	reader := io.MultiReader(strings.NewReader(shStr), jarFile)

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Printf("error: %s", err.Error())
		panic(err.Error())
	}
	ioutil.WriteFile(*exeFileName, b, os.ModePerm)
}
