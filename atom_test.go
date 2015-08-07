package feeds

import "testing"

var multipleLinksOutput = `<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title>jmoiron.net atom</title>
  <id>http://jmoiron.net/atom</id>
  <updated>2013-01-16T21:52:35-05:00</updated>
  <link href="http://jmoiron.net/atom" rel="self"></link>
  <link href="http://jmoiron.net/atom/1" rel="next-archive"></link>
  <entry>
    <title>Limiting Concurrency in Go</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/atom/limiting-concurrency-in-go/</id>
    <link href="http://jmoiron.net/blog/limiting-concurrency-in-go/" rel="self"></link>
    <link href="http://jmoiron.net/blog/limiting-concurrency-in-go.xml" rel="alternative"></link>
  </entry>
</feed>`

func TestMultipleLinksInFeed(t *testing.T) {
	feed := &AtomFeed{
		Xmlns:   "http://www.w3.org/2005/Atom",
		Id:      "http://jmoiron.net/atom",
		Title:   "jmoiron.net atom",
		Link:    &AtomLink{Href: "http://jmoiron.net/atom", Rel: "self"},
		Links:   []*AtomLink{&AtomLink{Href: "http://jmoiron.net/atom/1", Rel: "next-archive"}},
		Updated: "2013-01-16T21:52:35-05:00",
		Entries: []*AtomEntry{
			&AtomEntry{
				Id:      "tag:jmoiron.net,2013-01-16:/atom/limiting-concurrency-in-go/",
				Updated: "2013-01-16T21:52:35-05:00",
				Title:   "Limiting Concurrency in Go",
				Link:    &AtomLink{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/", Rel: "self"},
				Links:   []*AtomLink{&AtomLink{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go.xml", Rel: "alternative"}},
			},
		},
	}
	atom, _ := ToXML(feed)

	if atom != multipleLinksOutput {
		t.Errorf("Atom not what was expected.  Got:\n%s\n\nExpected:\n%s\n", atom, multipleLinksOutput)
	}
}
