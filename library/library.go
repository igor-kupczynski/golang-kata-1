package library

import (
	"fmt"
	"strings"
	"time"
)

type Author struct {
	Email     string
	FirstName string
	LastName  string
}

func (a Author) String() string {
	return fmt.Sprintf("%s %s, email: %s", a.FirstName, a.LastName, a.Email)
}

const (
	Book     = "book"
	Magazine = "magazine"
)

type Item struct {
	Kind        string // book or magazine
	Title       string
	Isbn        string
	Authors     []*Author
	Description *string
	PublishedAt *time.Time
}

func (x Item) String() string {
	authors := make([]string, 0, len(x.Authors))
	for _, author := range x.Authors {
		authors = append(authors, author.FirstName+" "+author.LastName)
	}
	if x.Kind == Book {
		return fmt.Sprintf("“%s” by %s, isbn: %s",
			x.Title, strings.Join(authors, " & "), x.Isbn)
	}
	return fmt.Sprintf("“%s” by %s, isbn: %s, published: %s",
		x.Title, strings.Join(authors, " & "), x.Isbn, x.PublishedAt.Format("2006-01-02"))
}
