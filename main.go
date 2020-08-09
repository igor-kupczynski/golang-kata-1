package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/igor-kupczynski/golang-kata-1/v1/library"
)

func main() {
	authors, err := readAuthors()
	if err != nil {
		panic(err)
	}

	books, err := readBooks(authors)
	if err != nil {
		panic(err)
	}

	magazines, err := readMagazines(authors)
	if err != nil {
		panic(err)
	}

	fmt.Println("# Authors")
	for _, v := range authors {
		fmt.Printf("- %s\n", v)
	}
	fmt.Println("")

	fmt.Println("# Books")
	for _, v := range books {
		fmt.Printf("- %s\n", v)
	}
	fmt.Println("")

	fmt.Println("# Magazines")
	for _, v := range magazines {
		fmt.Printf("- %s\n", v)
	}
	fmt.Println("")
}

func welcomeMessage() string {
	return "Hello world!"
}

func readAuthors() (map[string]library.Author, error) {
	file, err := os.Open("resources/authors.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	r.FieldsPerRecord = 3

	readHeaders := false
	authors := make(map[string]library.Author, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// skip header row
		if !readHeaders {
			readHeaders = true
			continue
		}
		authors[record[0]] = library.Author{Email: record[0], FirstName: record[1], LastName: record[2]}
	}

	return authors, nil
}

func readBooks(authorsDb map[string]library.Author) (map[string]library.Item, error) {
	file, err := os.Open("resources/books.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	r.FieldsPerRecord = 4

	readHeaders := false
	books := make(map[string]library.Item, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// skip header row
		if !readHeaders {
			readHeaders = true
			continue
		}

		emails := strings.Split(record[2], ",")
		authors := make([]*library.Author, 0, len(emails))
		for _, email := range emails {
			author, ok := authorsDb[email]
			if !ok {
				author = library.Author{
					Email:     email,
					FirstName: "Unknown",
					LastName:  "Unknown",
				}
			}
			authors = append(authors, &author)
		}

		books[record[1]] = library.Item{
			Kind:        library.Book,
			Title:       record[0],
			Isbn:        record[1],
			Authors:     authors,
			Description: &record[3],
		}
	}

	return books, nil
}

func readMagazines(authorsDb map[string]library.Author) (map[string]library.Item, error) {
	file, err := os.Open("resources/magazines.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	r.FieldsPerRecord = 4

	readHeaders := false
	magazines := make(map[string]library.Item, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// skip header row
		if !readHeaders {
			readHeaders = true
			continue
		}

		emails := strings.Split(record[2], ",")
		authors := make([]*library.Author, 0, len(emails))
		for _, email := range emails {
			author, ok := authorsDb[email]
			if !ok {
				author = library.Author{
					Email:     email,
					FirstName: "Unknown",
					LastName:  "Unknown",
				}
			}
			authors = append(authors, &author)
		}

		var publishedAt *time.Time
		if t, err := time.Parse("02.01.2006", record[3]); err == nil {
			publishedAt = &t
		}
		magazines[record[1]] = library.Item{
			Kind:        library.Magazine,
			Title:       record[0],
			Isbn:        record[1],
			Authors:     authors,
			PublishedAt: publishedAt,
		}
	}

	return magazines, nil
}
