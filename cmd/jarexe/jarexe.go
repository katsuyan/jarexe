package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
		log.Fatal("Not enough arguments. ex)jarexe standalone.jar")
	}

	jarFileName := flag.Args()[0]
	jarFile, err := os.Open(jarFileName)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	defer jarFile.Close()

	defaultExeFileName := filepath.Base(jarFileName)
	if ext := filepath.Ext(defaultExeFileName); ext == ".jar" {
		defaultExeFileName = defaultExeFileName[:len(defaultExeFileName)-4]
	}

	if *exeFileName == "" {
		*exeFileName = defaultExeFileName
	}

	shStr := fmt.Sprintf(`#!/bin/sh
		exec java %s -jar "$0" "$@"
		exit $?`, *javaOptions)

	reader := io.MultiReader(strings.NewReader(shStr), jarFile)

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	ioutil.WriteFile(*exeFileName, b, os.ModePerm)
}
