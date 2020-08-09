package library

import (
	"reflect"
	"sort"
	"testing"
)

func TestCatalogue_FindByAuthorEmail(t *testing.T) {

	authorA := Author{
		Email: "a@example.com",
	}
	authorB := Author{
		Email: "b@example.com",
	}
	bookByA := Item{
		Kind:    Book,
		Title:   "Big book of small items",
		Isbn:    "1234",
		Authors: []*Author{&authorA},
	}
	bookByAB := Item{
		Kind:    Book,
		Title:   "Cryptography in 24 hours",
		Isbn:    "2345",
		Authors: []*Author{&authorA, &authorB},
	}
	bookByB := Item{
		Kind:    Book,
		Title:   "My life as TikTok influencer",
		Isbn:    "5678",
		Authors: []*Author{&authorB},
	}
	magazineByB := Item{
		Kind:    Magazine,
		Title:   "V like Vitalize",
		Isbn:    "9012",
		Authors: []*Author{&authorB},
	}

	tests := []struct {
		name  string
		items map[string]Item
		email string
		want  []Item
	}{
		{
			name: "Returns nothing if no items by the author",
			items: map[string]Item{
				bookByB.Isbn: bookByB,
			},
			email: authorA.Email,
			want:  []Item{},
		},
		{
			name: "Returns books and magazines by the author",
			items: map[string]Item{
				bookByA.Isbn:     bookByA,
				bookByAB.Isbn:    bookByAB,
				bookByB.Isbn:     bookByB,
				magazineByB.Isbn: magazineByB,
			},
			email: authorB.Email,
			want:  []Item{bookByAB, bookByB, magazineByB},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Catalogue{
				Items: tt.items,
			}
			got := c.FindByAuthorEmail(tt.email)
			sort.Sort(byTitle(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByAuthorEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
