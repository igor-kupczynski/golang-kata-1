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
	item, ok := catalogue.GetByIsbn(isbn)
	if !ok {
		fmt.Printf("No items with isbn: %s\n\n", isbn)
		return
	}
	fmt.Printf("Found %s: %s\n\n", item.Kind, item)
}

func printWithEmail(catalogue *library.Catalogue, email string) {
	matching := catalogue.FindByAuthorEmail(email)
	fmt.Printf("# Items for author %s\n", email)
	if len(matching) == 0 {
		fmt.Printf("- no items found\n\n")
		return
	}
	for _, item := range matching {
		fmt.Printf("- %s: %s\n", item.Kind, item)
	}
	fmt.Println("")
}

func printSortedByTitle(catalogue *library.Catalogue) {
	items := catalogue.ListSortedByTitle()
	fmt.Printf("# All items (sorted by title)\n")
	for _, item := range items {
		fmt.Printf("- %s: %s\n", item.Kind, item)
	}
	fmt.Println("")
}
