package feeds

import (
	"encoding/json"
	"encoding/xml"
	"io"
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
	Extension   []interface{}
}

type Feed struct {
	Attrs       []xml.Attr
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

func (f *Feed) AddAttribute(name, nsURI string) {
	f.Attrs = append(f.Attrs, xml.Attr{
		Name:  xml.Name{Local: name},
		Value: nsURI,
	})
}

func (i *Item) AddExtension(extend interface{}) {
	i.Extension = append(i.Extension, extend)
}

func (i *Item) AddExtensionString(name string, nsURI string, value string) {
	i.Extension = append(i.Extension, struct {
		XMLName xml.Name
		Text    string `xml:",chardata"`
	}{
		XMLName: xml.Name{Local: name, Space: nsURI},
		Text:    value,
	})
}

func (i *Item) AddExtensionInt(name string, nsURI string, value int) {
	i.Extension = append(i.Extension, struct {
		XMLName xml.Name
		Number  int `xml:",chardata"`
	}{
		XMLName: xml.Name{Local: name, Space: nsURI},
		Number:  value,
	})
}

func (i *Item) AddExtensionUint(name string, nsURI string, value uint) {
	i.Extension = append(i.Extension, struct {
		XMLName xml.Name
		Number  uint `xml:",chardata"`
	}{
		XMLName: xml.Name{Local: name, Space: nsURI},
		Number:  value,
	})
}

func (i *Item) AddExtensionFloat64(name string, nsURI string, value float64) {
	i.Extension = append(i.Extension, struct {
		XMLName xml.Name
		Number  float64 `xml:",chardata"`
	}{
		XMLName: xml.Name{Local: name, Space: nsURI},
		Number:  value,
	})
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

// Write a feed object (either a Feed, AtomFeed, or RssFeed) as XML into
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

// Writes an Atom representation of this feed to the writer.
func (f *Feed) WriteAtom(w io.Writer) error {
	return WriteXML(&Atom{f}, w)
}

// creates an Rss representation of this feed
func (f *Feed) ToRss() (string, error) {
	r := &Rss{f}
	return ToXML(r)
}

// Writes an RSS representation of this feed to the writer.
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
