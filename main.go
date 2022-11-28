package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	files, err := ioutil.ReadDir("/home/buddhilw/keg/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if _, err := strconv.Atoi(file.Name()); err == nil {
			if file.IsDir() {
				nodeName := "/home/buddhilw/keg/" + file.Name()
				searchNodeMatch(nodeName, inputRegex)
			}
		}

	}
}

func searchNodeMatch(nodeName, regex string) {
	filesNode, err := ioutil.ReadDir(nodeName)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileNode := range filesNode {
		fmt.Println(fileNode.Name(), fileNode.IsDir()) // -> Directories inside keg node
		if !fileNode.IsDir() {
			searchMatch(nodeName, regex)
		}
	}
}

func searchMatch(filepath, regex string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		re := regexp.MustCompile(regex)
		matches := re.FindStringSubmatch(scanner.Text())
		fmt.Println(matches[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
