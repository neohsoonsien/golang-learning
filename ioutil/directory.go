package ioutil

import (
    "fmt"
    "io/ioutil"
    "log"
	"os"
)

func Directory() {

    // input current working directory
	directory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

    // read in all the files from the current working directory
    files, err := ioutil.ReadDir(directory)
    if err != nil {
        log.Fatal(err)
    }

    // print all the files from current direcory
    for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
    }
}