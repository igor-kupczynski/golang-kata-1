package main

import (
	"flag"
	"log"
)

func main() {
	var printAll = flag.Bool("printAll", false, "print all in the catalogue")
	flag.Parse()

	catalogue, err := readCatalogue()
	if err != nil {
		log.Fatal(catalogue)
	}

	if *printAll {
		printCatalogue(catalogue)
	}
}
