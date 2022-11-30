package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/jwalton/gchalk"
)

// var defaultKegPath string = os.UserHomeDir() + "keg"
var colorWithGchalk = gchalk.Red

func main() {
	gchalk.SetLevel(gchalk.LevelAnsi16m)

	var kegPath string
	var inputRegex string
	var caseSensitive string
	var surrounding int

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err, "Some error due to os.UserHomeDir; Couldn't decifer home directory for user")
		log.Fatal(err)
	}

	flag.StringVar(&kegPath, "p", homedir+"/keg/", "path: the KEG path on the machine.")
	flag.StringVar(&inputRegex, "reg", "reg", "regex: regular expression to be matched in keg nodes.")
	flag.StringVar(&caseSensitive, "c", "", "case: default is insensitive; any string will make it case-sensitive.")
	flag.IntVar(&surrounding, "s", 10, "surrounding: context-size of text-display around matching regex.")

	flag.Parse()
	// files, err := ioutil.ReadDir(kegPath) -> deprecated
	dirOpen, err := os.Open(kegPath)
	if err != nil {
		log.Print(err, "Some error occured, while opening kegPath")
		log.Fatal(err)
	}
	files, err := dirOpen.Readdir(0)
	if err != nil {
		log.Print(err, "Some error occured, while reading kegPath")
		log.Fatal(err)
	}

	for _, file := range files {
		if _, err := strconv.Atoi(file.Name()); err == nil {
			if file.IsDir() {
				nodeName := kegPath + "/" + file.Name()
				searchNodeMatch(kegPath, nodeName, inputRegex, caseSensitive, surrounding)
			}
		} else {
			continue
		}
	}
}

func searchNodeMatch(kegPath, nodeName, regex, caseSensitive string, surrounding int) {
	// filesNode, err := ioutil.ReadDir(nodeName) -> deprecated
	dirOpen, err := os.Open(nodeName)
	if err != nil {
		log.Print(err, "Some error occured, while opening node")
		log.Fatal(err)
	}

	if mapfile, maperr := os.Open(dirOpen.Name()); maperr == nil {
		defer mapfile.Close()
		info, _ := mapfile.Stat()
		if info.IsDir() {
			searchMatch(kegPath+"/"+string(info.Name())+"/README.md", regex, caseSensitive, surrounding)
		}
	}
}

func searchMatch(filepath, regex, caseSensitive string, surrounding int) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		log.Println(err, "error opening fileNode content")
		log.Fatal(err)
	}

	if caseSensitive != "" {
		re, err := regexp.Compile(regex)
		if err != nil {
			log.Println(err, "Error with regex")
			log.Fatal(err)
		}
		matchText(filepath, fileContent, re, surrounding)
	} else {
		re, err := regexp.Compile("(?i)" + regex)
		if err != nil {
			log.Println(err, "Error with regex")
			log.Fatal(err)
		}
		matchText(filepath, fileContent, re, surrounding)
	}
}

func matchText(filepath string, fileContent []byte, re *regexp.Regexp, surrounding int) {
	matches := re.FindAllStringSubmatchIndex(string(fileContent), -1)
	if len(matches) != 0 {
		for _, match := range matches {
			fmt.Println(match)
			if match[0]-surrounding < 0 && match[1]+surrounding <= len(fileContent) {
				// Case the left-context would go over the left-limit of the file, but not right:
				// return from the beginning of the file, until right-context.
				fmt.Println(gchalk.Blue(filepath+":"),
					string(fileContent[:match[0]])+
						gchalk.Red(string(fileContent[match[0]:match[1]]))+
						string(fileContent[match[1]:match[1]+surrounding]))
			} else if match[0]-surrounding >= 0 && match[1]+surrounding > len(fileContent) {
				// Case the left-context would go over the right-limit of the file, but not left:
				// return from the left-context of the file, until end.
				fmt.Println(gchalk.Blue(filepath+":"),
					string(fileContent[match[0]-surrounding:match[0]])+
						gchalk.Red(string(fileContent[match[0]:match[1]]))+
						string(fileContent[match[1]:]))
			} else if match[0]-surrounding >= 0 && match[1]+surrounding > len(fileContent) {
				// Case the left and right-context would go over the limit:
				fmt.Println(gchalk.Blue(filepath+":"),
					string(fileContent[:match[0]])+
						gchalk.Red(string(fileContent[match[0]:match[1]]))+
						string(fileContent[match[1]:]))
			} else {
				// Default
				fmt.Println(gchalk.Blue(filepath+":"),
					string(fileContent[match[0]-surrounding:match[0]])+
						gchalk.Red(string(fileContent[match[0]:match[1]]))+
						string(fileContent[match[1]:match[1]+surrounding]))
			}
		}
	}
}
