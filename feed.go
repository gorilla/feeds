package feeds

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"sort"
	"time"
)

type Link struct {
	Href, Rel, Type, Length string
}

type Author struct {
	Name, Email string
}

type Image struct {
	Url, Title, Link string
	Width, Height    int
}

type Enclosure struct {
	Url, Length, Type string
}

type Item struct {
	Title       string
	Link        *Link
	Source      *Link
	Author      *Author
	Description string // used as description in rss, summary in atom
	Id          string // used as guid in rss, id in atom
	Updated     time.Time
	Created     time.Time
	Enclosure   *Enclosure
	Content     string
}

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

// add a new Item to a Feed
func (f *Feed) Add(item *Item) {
	f.Items = append(f.Items, item)
}

// returns the first non-zero time formatted as a string or ""
func anyTimeFormat(format string, times ...time.Time) string {
	for _, t := range times {
		if !t.IsZero() {
			return t.Format(format)
		}
	}
	return ""
}

// interface used by ToXML to get a object suitable for exporting XML.
type XmlFeed interface {
	FeedXml() interface{}
}

// turn a feed object (either a Feed, AtomFeed, or RssFeed) into xml
// returns an error if xml marshaling fails
func ToXML(feed XmlFeed, header string) (string, error) {
	x := feed.FeedXml()
	data, err := xml.MarshalIndent(x, "", "  ")
	if err != nil {
		return "", err
	}
	s := header + string(data)
	return s, nil
}

// WriteXML writes a feed object (either a Feed, AtomFeed, or RssFeed) as XML into
// the writer. Returns an error if XML marshaling fails.
func WriteXML(feed XmlFeed, header string, w io.Writer) error {
	x := feed.FeedXml()
	// write default xml header, without the newline
	if _, err := w.Write([]byte(header)); err != nil {
		return err
	}
	e := xml.NewEncoder(w)
	e.Indent("", "  ")
	return e.Encode(x)
}

// ToAtom creates an Atom representation of this feed
func (f *Feed) ToAtom() (string, error) {
	return f.ToAtomWithHeader(xml.Header[:len(xml.Header)-1])
}

// ToAtomWithHeader creates an Atom representation of this feed with a custom header
func (f *Feed) ToAtomWithHeader(header string) (string, error) {
	a := &Atom{f}
	return ToXML(a, header)
}

// WriteAtom writes an Atom representation of this feed to the writer.
func (f *Feed) WriteAtom(w io.Writer) error {
	return WriteXML(&Atom{f}, xml.Header[:len(xml.Header)-1], w)
}

// WriteAtomWithHeader writes an Atom representation of this feed to the writer along with a custom header.
func (f *Feed) WriteAtomWithHeader(w io.Writer, header string) error {
	return WriteXML(&Atom{f}, header, w)
}

//ToRss creates an Rss representation of this feed
func (f *Feed) ToRss() (string, error) {
	r := &Rss{f}
	return ToXML(r, xml.Header[:len(xml.Header)-1])
}

//ToRssWithHeader creates an Rss representation of this feed with a custom header
func (f *Feed) ToRssWithHeader(header string) (string, error) {
	r := &Rss{f}
	return ToXML(r, header)
}

// WriteRss writes an RSS representation of this feed to the writer.
func (f *Feed) WriteRss(w io.Writer) error {
	return WriteXML(&Rss{f}, xml.Header[:len(xml.Header)-1], w)
}

// WriteRssWithHeader writes an RSS representation of this feed to the writer along with a custom header.
func (f *Feed) WriteRssWithHeader(w io.Writer, header string) error {
	return WriteXML(&Rss{f}, header, w)
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
