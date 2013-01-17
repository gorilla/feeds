package syndicate

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
)

// Generates Atom feed as XML

const ns = "http://www.w3.org/2005/Atom"

type AtomPerson struct {
	Name  string `xml:"name,omitempty"`
	Uri   string `xml:"uri,omitempty"`
	Email string `xml:"email,omitempty"`
}

type AtomSummary struct {
	Content string `xml:",chardata"`
	Type    string `xml:"type,attr"`
}

type AtomAuthor struct {
	XMLName xml.Name `xml:"author"`
	AtomPerson
}

type AtomContributor struct {
	XMLName xml.Name `xml:"contributor"`
	AtomPerson
}

type AtomEntry struct {
	XMLName xml.Name `xml:"entry"`
	Title   string   `xml:"title"`
	Link    *AtomLink
	Updated string       `xml:"updated"`
	Id      string       `xml:"id"`
	Summary *AtomSummary `xml:"summary,omitempty"`
	Author  *AtomAuthor
}

type AtomLink struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr,omitempty"`
}

type AtomFeed struct {
	XMLName     xml.Name `xml:"feed"`
	Xmlns       string   `xml:"xmlns,attr"`
	Category    string   `xml:"category,omitempty"`
	Icon        string   `xml:"icon,omitempty"`
	Logo        string   `xml:"logo,omitempty"`
	Rights      string   `xml:"rights,omitempty"`
	Title       string   `xml:"title"`
	Subtitle    string   `xml:"subtitle,omitempty"`
	Id          string   `xml:"id,omitempty"`
	Updated     string   `xml:"updated"`
	Link        *AtomLink
	Contributor *AtomContributor
	Entries     []*AtomEntry
}

type Atom struct {
	*Feed
}

func newAtomEntry(i *Item) *AtomEntry {
	id := i.Id
	// assume the description is html
	s := &AtomSummary{i.Description, "html"}

	if len(id) == 0 {
		// if there's no id set, try to create one, either from data or just a uuid
		if len(i.Link.Href) > 0 && (!i.Created.IsZero() || !i.Updated.IsZero()) {
			dateStr := anyTimeFormat("2006-01-02", i.Updated, i.Created)
			host, path := i.Link.Href, "/invalid.html"
			if url, err := url.Parse(i.Link.Href); err == nil {
				host, path = url.Host, url.Path
			}
			id = fmt.Sprintf("tag:%s,%s:%s", host, dateStr, path)
		} else {
			id = "urn:uuid:" + NewUUID().String()
		}
	}
	var name, email string
	if i.Author != nil {
		name, email = i.Author.Name, i.Author.Email
	}

	x := &AtomEntry{
		Title:   i.Title,
		Link:    &AtomLink{Href: i.Link.Href, Rel: i.Link.Rel},
		Summary: s,
		Id:      id,
		Updated: anyTimeFormat(time.RFC3339, i.Updated, i.Created),
	}
	if len(name) > 0 || len(email) > 0 {
		x.Author = &AtomAuthor{AtomPerson: AtomPerson{Name: name, Email: email}}
	}
	return x
}

func (a *Atom) FeedXml() interface{} {
	updated := anyTimeFormat(time.RFC3339, a.Updated, a.Created)
	feed := &AtomFeed{
		Xmlns:    ns,
		Title:    a.Title,
		Link:     &AtomLink{Href: a.Link.Href, Rel: a.Link.Rel},
		Subtitle: a.Description,
		Id:       a.Link.Href,
		Updated:  updated,
		Rights:   a.Copyright,
	}
	for _, e := range a.Items {
		feed.Entries = append(feed.Entries, newAtomEntry(e))
	}
	return feed
}

// support the ToXML function for AtomFeeds directly
func (a *AtomFeed) FeedXml() interface{} {
	return a
}
