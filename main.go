package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

// var defaultKegPath string = os.UserHomeDir() + "keg"

func main() {
	var kegPath string
	var inputRegex string
	var caseSensitive string

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err, "Some error due to os.UserHomeDir; Couldn't decifer home directory for user")
		log.Fatal(err)
	}

	flag.StringVar(&kegPath, "p", homedir+"/keg/", "path: the KEG path on the machine.")
	flag.StringVar(&inputRegex, "reg", "reg", "regex: regular expression to be matched in keg nodes.")
	flag.StringVar(&caseSensitive, "c", "", "case: default is insensitive; any string will make it case-sensitive")

	flag.Parse()

	files, err := ioutil.ReadDir(kegPath)
	if err != nil {
		log.Print(err, "Some error occured, while reading kegPath")
		log.Fatal(err)
	}

	for _, file := range files {
		if _, err := strconv.Atoi(file.Name()); err == nil {
			if file.IsDir() {
				nodeName := kegPath + "/" + file.Name()
				searchNodeMatch(nodeName, inputRegex, caseSensitive)
			}
		} else {
			continue
		}
	}
}

func searchNodeMatch(nodeName, regex, caseSensitive string) {
	filesNode, err := ioutil.ReadDir(nodeName)
	if err != nil {
		log.Println(err, "error at searchNodeMatch")
		log.Fatal(err)
	}
	for _, fileNode := range filesNode {
		if !fileNode.IsDir() {
			searchMatch(nodeName+"/README.md", regex, caseSensitive)
		}
	}
}

func searchMatch(filepath, regex, caseSensitive string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err, "error opening fileNode")
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if caseSensitive != "" {
		re, err := regexp.Compile(regex)
		if err != nil {
			log.Println(err, "Error with regex")
			log.Fatal(err)
		}
		matchText(filepath, scanner, re)
	} else {
		re, err := regexp.Compile("(?i)" + regex)
		if err != nil {
			log.Println(err, "Error with regex")
			log.Fatal(err)
		}
		matchText(filepath, scanner, re)
	}

}

func matchText(filepath string, scanner *bufio.Scanner, re *regexp.Regexp) {
	for scanner.Scan() {
		if re.MatchString(scanner.Text()) {
			fmt.Println(filepath)
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err, "scanner error")
		log.Fatal(err)
	}
}
