package main

import (
	"flag"
	"log"
)

func main() {
	var printAll = flag.Bool("printAll", false, "print all authors and items in the catalogue")
	var byTitle = flag.Bool("byTitle", false, "print all items in the catalogue sorted by title")
	var findIsbn = flag.String("findIsbn", "", "find a book or magazine by isbn")
	var findEmail = flag.String("findEmail", "", "find a book or magazine by email of one of the authors")

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

	if *findEmail != "" {
		printWithEmail(catalogue, *findEmail)
	}

	if *byTitle {
		printSortedByTitle(catalogue)
	}
}
