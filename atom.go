package syndicate

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
)

// Generates Atom feed as XML

const ns = "http://www.w3.org/2005/Atom"

type atomSummary struct {
	S    string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type atomEntry struct {
	XMLName xml.Name `xml:"entry"`
	Title   string   `xml:"title"`
	Link    *atomLink
	Updated string       `xml:"updated"`
	Id      string       `xml:"id"`
	Summary *atomSummary `xml:"summary"`
}

type atomLink struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
}

type atomFeed struct {
	XMLName xml.Name `xml:"feed"`
	Ns      string   `xml:"xmlns,attr"`
	Title   string   `xml:"title"`
	Link    *atomLink
	Id      string `xml:"id"`
	Updated string `xml:"updated"`
	Entries []*atomEntry
}

type Atom struct {
	*Feed
}

func newAtomEntry(i *Item) *atomEntry {
	id := i.Id
	// assume the description is html
	s := &atomSummary{i.Description, "html"}

	// try to get a single timestamp, since we only have one  in atom 
	ts := i.Updated
	if ts.IsZero() {
		ts = i.Created
	}
	// <id>tag:blog.kowalczyk.info,2012-09-11:/item/1.html</id>
	if len(id) == 0 {
		// if there's no id set, try to create one, either from data or just a uuid
		if len(i.Link.Href) > 0 && (!i.Created.IsZero() || !i.Updated.IsZero()) {
			dateStr := ts.Format("2006-01-02")
			host, path := i.Link.Href, "/invalid.html"
			if url, err := url.Parse(i.Link.Href); err == nil {
				host, path = url.Host, url.Path
			}
			id = fmt.Sprintf("tag:%s,%s:%s", host, dateStr, path)
		} else {
			id = "urn:uuid:" + NewUUID().String()
		}
	}
	x := &atomEntry{
		Title:   i.Title,
		Link:    &atomLink{Href: i.Link.Href, Rel: i.Link.Rel},
		Summary: s,
		Id:      id,
		Updated: i.Updated.Format(time.RFC3339)}
	return x
}

func (a *Atom) FeedXml() interface{} {
	ts := a.Updated
	if ts.IsZero() {
		ts = a.Created
	}
	feed := &atomFeed{
		Ns:      ns,
		Title:   a.Title,
		Link:    &atomLink{Href: a.Link.Href, Rel: a.Link.Rel},
		Id:      a.Link.Href,
		Updated: ts.Format(time.RFC3339)}
	for _, e := range a.Items {
		feed.Entries = append(feed.Entries, newAtomEntry(e))
	}

	return feed
}
