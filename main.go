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

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Couldn't decifer home directory for user")
	}

	flag.StringVar(&kegPath, "p", homedir+"/keg/", "path: the KEG path on the machine.")
	flag.StringVar(&inputRegex, "reg", "reg", "regex: regular expression to be matched in keg nodes.")

	flag.Parse()

	// fmt.Println(kegPath, inputRegex)

	files, err := ioutil.ReadDir(kegPath)
	if err != nil {
		log.Print(err, "Some error occured")
	}

	for _, file := range files {
		if _, err := strconv.Atoi(file.Name()); err == nil {
			if file.IsDir() {
				nodeName := kegPath + "/" + file.Name()
				// fmt.Println(nodeName)
				searchNodeMatch(nodeName, inputRegex)
			}
		} else {
			continue
		}
	}
}

func searchNodeMatch(nodeName, regex string) {
	filesNode, err := ioutil.ReadDir(nodeName)
	if err != nil {
		// log.Fatal(err)
		log.Println(err, "error at searchNodeMatch")
	}
	for _, fileNode := range filesNode {
		// fmt.Println(fileNode.Name(), fileNode.IsDir()) // -> Directories inside keg node
		if !fileNode.IsDir() {
			fmt.Println(nodeName, regex)
			searchMatch(nodeName+"/README.md", regex)
		}
	}
}

func searchMatch(filepath, regex string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err, "error opening fileNode")
		// log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		re := regexp.MustCompile(regex)
		matches := re.FindStringSubmatch(scanner.Text())
		fmt.Println(matches[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
