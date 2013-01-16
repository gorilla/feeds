package syndicate

import (
	"encoding/xml"
	"time"
)

type Link struct {
	Href, Rel string
}

type Author struct {
	Name, Email string
}

type Item struct {
	Title       string
	Link        *Link
	Author      *Author
	Description string // used as description in rss, summary in atom
	Id          string // used as guid in rss, id in atom
	Updated     time.Time
	Created     time.Time
}

type Feed struct {
	Title       string
	Link        *Link
	Description string
	Author      *Author
	Updated     time.Time
	Id          string
	Subtitle    string
	Created     time.Time
	Items       []*Item
}

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

type XmlFeed interface {
	FeedXml() interface{}
}

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

func (f *Feed) ToAtom() (string, error) {
	a := &Atom{f}
	return ToXML(a)
}

func (f *Feed) ToRss() (string, error) {
	r := &Rss{f}
	return ToXML(r)
}
