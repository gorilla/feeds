package feeds

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"sort"
	"time"
)

// Link represents an RSS/Atom link element.
type Link struct {
	Href, Rel, Type, Length string
}

// Author represents an author of a feed or item.
type Author struct {
	Name, Email string
}

// Image represents an image or logo for a feed.
type Image struct {
	Url, Title, Link string
	Width, Height    int
}

// Enclosure represents an enclosure element in an RSS/Atom item.
type Enclosure struct {
	Url, Length, Type string
}

// Item represents a single feed item or story.
type Item struct {
	Title       string
	Link        *Link
	Source      *Link
	Author      *Author
	Description string // used as description in rss, summary in atom
	Id          string // used as guid in rss, id in atom
	IsPermaLink string // an optional parameter for guid in rss
	Updated     time.Time
	Created     time.Time
	Enclosure   *Enclosure
	Content     string
}

// Feed represents a syndication feed containing items and metadata.
type Feed struct {
	Title       string
	Link        *Link
	Description string
	Author      *Author
	Updated     time.Time
	Created     time.Time
	Id          string
	Subtitle    string
	Items       []*Item
	Copyright   string
	Image       *Image
}

// Add adds a new Item to a Feed.
func (f *Feed) Add(item *Item) {
	f.Items = append(f.Items, item)
}

// anyTimeFormat returns the first non-zero time formatted as a string or "".
func anyTimeFormat(format string, times ...time.Time) string {
	for _, t := range times {
		if !t.IsZero() {
			return t.Format(format)
		}
	}
	return ""
}

// XmlFeed is an interface used by ToXML to get an object suitable for exporting XML.
type XmlFeed interface {
	FeedXml() interface{}
}

// ToXML converts a feed object (either a Feed, AtomFeed, or RssFeed) into XML.
// It returns an error if XML marshaling fails.
func ToXML(feed XmlFeed) (string, error) {
	x := feed.FeedXml()
	data, err := xml.MarshalIndent(x, "", "  ")
	if err != nil {
		return "", err
	}
	// strip empty line from default xml header
	s := xml.Header[:len(xml.Header)-1] + string(data)
	return s, nil
}

// WriteXML writes a feed object (either a Feed, AtomFeed, or RssFeed) as XML into
// the writer. Returns an error if XML marshaling fails.
func WriteXML(feed XmlFeed, w io.Writer) error {
	x := feed.FeedXml()
	// write default xml header, without the newline
	if _, err := w.Write([]byte(xml.Header[:len(xml.Header)-1])); err != nil {
		return err
	}
	e := xml.NewEncoder(w)
	e.Indent("", "  ")
	return e.Encode(x)
}

// creates an Atom representation of this feed
func (f *Feed) ToAtom() (string, error) {
	a := &Atom{f}
	return ToXML(a)
}

// WriteAtom writes an Atom representation of this feed to the writer.
func (f *Feed) WriteAtom(w io.Writer) error {
	return WriteXML(&Atom{f}, w)
}

// creates an Rss representation of this feed
func (f *Feed) ToRss() (string, error) {
	r := &Rss{f}
	return ToXML(r)
}

// WriteRss writes an RSS representation of this feed to the writer.
func (f *Feed) WriteRss(w io.Writer) error {
	return WriteXML(&Rss{f}, w)
}

// ToJSON creates a JSON Feed representation of this feed
func (f *Feed) ToJSON() (string, error) {
	j := &JSON{f}
	return j.ToJSON()
}

// WriteJSON writes an JSON representation of this feed to the writer.
func (f *Feed) WriteJSON(w io.Writer) error {
	j := &JSON{f}
	feed := j.JSONFeed()

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(feed)
}

// Sort sorts the Items in the feed with the given less function.
func (f *Feed) Sort(less func(a, b *Item) bool) {
	lessFunc := func(i, j int) bool {
		return less(f.Items[i], f.Items[j])
	}
	sort.SliceStable(f.Items, lessFunc)
}
