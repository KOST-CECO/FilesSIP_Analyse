package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// read all folder in actual folder and recurse through subsequent folders
func walkdir(dir string) {
	files, _ := ioutil.ReadDir(dir)
	// fmt.Println(dir)

	for _, f := range files {

		if f.IsDir() {
			// test for arelda SIP
			if !areldaSIP(dir + string(os.PathSeparator) + f.Name()) {
				// subsequent folder walk
				walkdir(dir + string(os.PathSeparator) + f.Name())
			}
		}
	}
}

// test folder for arelda SIP
func areldaSIP(dir string) bool {
	files, _ := ioutil.ReadDir(dir)

	// only 2 subdirectories
	if len(files) == 2 {

		// only subdirectories "content" and "header"
		if files[0].Name()+" "+files[1].Name() == "content header" {
			// fmt.Println("SIP: " + dir)

			// read metadata.xml
			metadata, err := ioutil.ReadFile(dir + string(os.PathSeparator) + "header" + string(os.PathSeparator) + "metadata.xml")
			if err != nil {
				fmt.Println("metadata.xml not found")
				return false
			}

			meta := string(metadata)
			// metadata.xml contains "ablieferungFilesSIP"
			if strings.Contains(meta, "ablieferungFilesSIP") {
				// fmt.Println("FilesSIP: " + dir)
				if !strings.Contains(meta, "<dateiRef>") {
					fmt.Println("FilesSIP ohne DateiRef: " + dir)
					return true
				}
			}
		}
	}
	return false
}

// read prog arguments, test directory and start directory walk
func main() {

	// read initial parameter from command line
	if len(os.Args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " folder")
		os.Exit(-1)
	}

	// check for valid folder
	fo, foerr := os.Stat(os.Args[1])
	if foerr != nil {
		log.Fatal(os.Args[1] + " is not a valid folder name")
	}
	if !fo.IsDir() {
		log.Fatal(os.Args[1] + " is not a folder")
	}

	// recurse through subsequent folders
	walkdir(os.Args[1])

}
