package main

import (
	"fmt"

	"github.com/igor-kupczynski/golang-kata-1/v1/library"
)

func printCatalogue(catalogue *library.Catalogue) {
	fmt.Println("# Authors")
	for _, author := range catalogue.Authors {
		fmt.Printf("- %s\n", author)
	}
	fmt.Println("")

	fmt.Println("# Books")
	for _, item := range catalogue.Items {
		if item.Kind == library.Book {
			fmt.Printf("- %s\n", item)
		}
	}
	fmt.Println("")

	fmt.Println("# Magazines")
	for _, item := range catalogue.Items {
		if item.Kind == library.Magazine {
			fmt.Printf("- %s\n", item)
		}
	}
	fmt.Println("")
}

func printWithIsbn(catalogue *library.Catalogue, isbn string) {
	item, ok := catalogue.Items[isbn]
	if !ok {
		fmt.Printf("No items with isbn: %s\n\n", isbn)
		return
	}
	fmt.Printf("Found %s: %s\n\n", item.Kind, item)
}
