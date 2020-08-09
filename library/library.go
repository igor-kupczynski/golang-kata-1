package library

import (
	"fmt"
	"sort"
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

type byTitle []Item

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }

type Catalogue struct {
	Authors []Author
	Items   map[string]Item
}

func (c *Catalogue) GetByIsbn(isbn string) (Item, bool) {
	item, ok := c.Items[isbn]
	return item, ok
}

func (c *Catalogue) FindByAuthorEmail(email string) []Item {
	// This is slow and can be improved
	matching := make([]Item, 0)
	for _, item := range c.Items {
		for _, author := range item.Authors {
			if author.Email == email {
				matching = append(matching, item)
				break
			}
		}
	}
	return matching
}

func (c *Catalogue) ListSortedByTitle() []Item {
	items := make([]Item, 0, len(c.Items))
	for _, item := range c.Items {
		items = append(items, item)
	}
	sort.Sort(byTitle(items))
	return items
}
