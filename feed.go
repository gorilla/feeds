package syndicate

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Link struct {
	Href string
	Rel  string
}

type Author struct {
	Name  string
	Email string
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
	Author      *Author
	Created     time.Time
	Updated     time.Time
	Description string
	Subtitle    string
	Id          string
	Items       []*Item
}

func (f *Feed) Add(item *Item) {
	f.Items = append(f.Items, item)
}

type XmlFeed interface {
	FeedXml() interface{}
}

func ToXML(feed XmlFeed) (string, error) {
	x := feed.FeedXml()
	data, err := xml.MarshalIndent(x, " ", " ")
	if err != nil {
		return "", err
	}
	s := xml.Header[:len(xml.Header)-1] + string(data)
	return s, nil
}

func (f *Feed) ToAtom() (string, error) {
	a := &Atom{f}
	return ToXML(a)
}

func (f *Feed) ToRss(version ...float64) (string, error) {
	vers := 2.0
	if len(version) > 0 {
		vers = version[0]
	}
	/*
		r := &RssFeed{f}
		return ToXML(r)
	*/
	return fmt.Sprint(vers), nil
}
