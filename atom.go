package syndicate

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
)

// Generates Atom feed as XML

const ns = "http://www.w3.org/2005/Atom"

type AtomSummary struct {
	S    string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type AtomEntry struct {
	XMLName     xml.Name `xml:"entry"`
	Title       string   `xml:"title"`
	Link        *AtomLink
	Updated     string       `xml:"updated"`
	Id          string       `xml:"id"`
	Summary     *AtomSummary `xml:"summary,omitempty"`
	AuthorName  string       `xml:"author>name,omitempty"`
	AuthorEmail string       `xml:"author>email,omitempty"`
}

type AtomLink struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr,omitempty"`
}

type AtomFeed struct {
	XMLName xml.Name `xml:"feed"`
	Ns      string   `xml:"xmlns,attr"`
	Title   string   `xml:"title"`
	Link    *AtomLink
	Id      string `xml:"id,omitempty"`
	Updated string `xml:"updated"`
	Summary string `xml:"summary,omitempty"`
	Entries []*AtomEntry
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
		Title:       i.Title,
		Link:        &AtomLink{Href: i.Link.Href, Rel: i.Link.Rel},
		Summary:     s,
		Id:          id,
		Updated:     i.Updated.Format(time.RFC3339),
		AuthorName:  name,
		AuthorEmail: email,
	}
	return x
}

func (a *Atom) FeedXml() interface{} {
	updated := anyTimeFormat(time.RFC3339, a.Updated, a.Created)
	feed := &AtomFeed{
		Ns:      ns,
		Title:   a.Title,
		Link:    &AtomLink{Href: a.Link.Href, Rel: a.Link.Rel},
		Summary: a.Description,
		Id:      a.Link.Href,
		Updated: updated,
	}
	for _, e := range a.Items {
		feed.Entries = append(feed.Entries, newAtomEntry(e))
	}
	return feed
}
