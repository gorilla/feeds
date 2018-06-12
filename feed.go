package feeds

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"strings"
	"time"
)

// default indentation level for feeds
const defaultIndent = 2

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

// SerializeOpts represents serialization options for a feed.
type SerializeOpts struct {
	Indent int // how many spaces to indent
}

// get the indentation string for this option set
func (s SerializeOpts) indentString() string {
	return strings.Repeat(" ", s.Indent)
}

// DefaultOptions are the default serialization options for the module. This
// object can be mutated to update the default options for the module.
var DefaultOptions = &SerializeOpts{Indent: defaultIndent}

// Feed represents a generic XML feed.
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
	Options     *SerializeOpts
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

// like ToXML, but with options
func toXMLWithOptions(feed XmlFeed, opts *SerializeOpts) (string, error) {
	x := feed.FeedXml()
	data, err := xml.MarshalIndent(x, "", opts.indentString())
	if err != nil {
		return "", err
	}
	// strip empty line from default xml header
	s := xml.Header[:len(xml.Header)-1] + string(data)
	return s, nil
}

// ToXML turns a feed object (either a Feed, AtomFeed, or RssFeed)
// into XML, and returns an error if XML marshaling fails.
func ToXML(feed XmlFeed) (string, error) {
	return toXMLWithOptions(feed, DefaultOptions)
}

// like WriteXML, but with options
func writeXMLWithOptions(feed XmlFeed, w io.Writer, opts *SerializeOpts) error {
	x := feed.FeedXml()
	// write default xml header, without the newline
	if _, err := w.Write([]byte(xml.Header[:len(xml.Header)-1])); err != nil {
		return err
	}
	e := xml.NewEncoder(w)
	e.Indent("", opts.indentString())
	return e.Encode(x)
}

// WriteXML writes a feed object (either a Feed, AtomFeed, or RssFeed) as XML
// into the writer. Returns an error if XML marshaling fails.
func WriteXML(feed XmlFeed, w io.Writer) error {
	return writeXMLWithOptions(feed, w, DefaultOptions)
}

// return Options, or DefaultOptions if Options is unset
func (f *Feed) options() *SerializeOpts {
	if f.Options == nil {
		return DefaultOptions
	}
	return f.Options
}

// creates an Atom representation of this feed
func (f *Feed) ToAtom() (string, error) {
	a := &Atom{f}
	return toXMLWithOptions(a, f.options())
}

// Writes an Atom representation of this feed to the writer.
func (f *Feed) WriteAtom(w io.Writer) error {
	return writeXMLWithOptions(&Atom{f}, w, f.options())
}

// creates an Rss representation of this feed
func (f *Feed) ToRss() (string, error) {
	r := &Rss{f}
	return toXMLWithOptions(r, f.options())
}

// Writes an RSS representation of this feed to the writer.
func (f *Feed) WriteRss(w io.Writer) error {
	return writeXMLWithOptions(&Rss{f}, w, f.options())
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
	e.SetIndent("", f.options().indentString())
	return e.Encode(feed)
}
