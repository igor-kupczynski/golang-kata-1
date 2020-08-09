package main

import (
	"flag"
	"log"
)

func main() {
	var printAll = flag.Bool("printAll", false, "print all in the catalogue")
	var findIsbn = flag.String("findIsbn", "", "find a book or magazine by isbn")
	flag.Parse()

	catalogue, err := readCatalogue()
	if err != nil {
		log.Fatal(catalogue)
	}

	if *printAll {
		printCatalogue(catalogue)
	}

	if *findIsbn != "" {
		printWithIsbn(catalogue, *findIsbn)
	}
}
