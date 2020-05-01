// Package bookmarks is used to parse Netscape compatible bookmark file exports
package bookmarks

import (
	"io"
	"strconv"
	"time"

	"golang.org/x/net/html"
)

// Attribute contains named element attributes
type Attribute struct {
	Key string
	Val string
}

// Bookmark contains single bookmark
type Bookmark struct {
	Name       string
	Href       string
	AddDate    string
	Attributes []Attribute
}

// Parse parses read stream to list of Bookmarks
func Parse(r io.Reader) ([]Bookmark, error) {
	var bookmarks []Bookmark
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			{
				if z.Err() != io.EOF {
					return bookmarks, z.Err()
				}
				return bookmarks, nil
			}
		case html.StartTagToken:
			{
				tk := z.Token()
				isAnchor := tk.Data == "a"
				if isAnchor {
					bookmarks = append(bookmarks, parseLink(tk, z))
				}
			}
		}
	}
}

// ParseTimestamp converts timestamp to time.Time
func ParseTimestamp(ts string) (time.Time, error) {
	i, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Now(), err
	}
	return time.Unix(i, 0), nil
}

func parseLink(tk html.Token, z *html.Tokenizer) Bookmark {
	bookmark := Bookmark{
		Attributes: make([]Attribute, 0),
	}
	for _, a := range tk.Attr {
		if a.Key == "href" {
			bookmark.Href = a.Val
			continue
		}
		if a.Key == "add_date" {
			bookmark.AddDate = a.Val
			continue
		}
		bookmark.Attributes = append(bookmark.Attributes, Attribute{a.Key, a.Val})
	}
	for {
		tt := z.Next()
		switch tt {
		case html.TextToken:
			{
				bookmark.Name = string(z.Text())
			}
		case html.EndTagToken:
			{
				return bookmark
			}
		}
	}
}
