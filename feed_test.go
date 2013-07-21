package feeds

import (
	"encoding/xml"
	"testing"
	"time"
)

var atomOutput = `<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title>jmoiron.net blog</title>
  <id>http://jmoiron.net/blog</id>
  <updated>2013-01-16T21:52:35-05:00</updated>
  <rights>This work is copyright © Benjamin Button</rights>
  <subtitle>discussion about tech, footie, photos</subtitle>
  <link href="http://jmoiron.net/blog"></link>
  <author>
    <name>Jason Moiron</name>
    <email>jmoiron@jmoiron.net</email>
  </author>
  <entry>
    <title>Limiting Concurrency in Go</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/blog/limiting-concurrency-in-go/</id>
    <content type="html">A discussion on controlled parallelism in golang</content>
    <link href="http://jmoiron.net/blog/limiting-concurrency-in-go/"></link>
    <author>
      <name>Jason Moiron</name>
      <email>jmoiron@jmoiron.net</email>
    </author>
  </entry>
  <entry>
    <title>Logic-less Template Redux</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/blog/logicless-template-redux/</id>
    <content type="html">More thoughts on logicless templates</content>
    <link href="http://jmoiron.net/blog/logicless-template-redux/"></link>
  </entry>
  <entry>
    <title>Idiomatic Code Reuse in Go</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/blog/idiomatic-code-reuse-in-go/</id>
    <content type="html">How to use interfaces &lt;em&gt;effectively&lt;/em&gt;</content>
    <link href="http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"></link>
  </entry>
</feed>`

var rssOutput = `<?xml version="1.0" encoding="UTF-8"?><rss version="2.0">
  <channel>
    <title>jmoiron.net blog</title>
    <link>http://jmoiron.net/blog</link>
    <description>discussion about tech, footie, photos</description>
    <copyright>This work is copyright © Benjamin Button</copyright>
    <managingEditor>jmoiron@jmoiron.net (Jason Moiron)</managingEditor>
    <pubDate>2013-01-16T21:52:35-05:00</pubDate>
    <item>
      <title>Limiting Concurrency in Go</title>
      <link>http://jmoiron.net/blog/limiting-concurrency-in-go/</link>
      <description>A discussion on controlled parallelism in golang</description>
      <author>Jason Moiron</author>
      <pubDate>2013-01-16T21:52:35-05:00</pubDate>
    </item>
    <item>
      <title>Logic-less Template Redux</title>
      <link>http://jmoiron.net/blog/logicless-template-redux/</link>
      <description>More thoughts on logicless templates</description>
      <pubDate>2013-01-16T21:52:35-05:00</pubDate>
    </item>
    <item>
      <title>Idiomatic Code Reuse in Go</title>
      <link>http://jmoiron.net/blog/idiomatic-code-reuse-in-go/</link>
      <description>How to use interfaces &lt;em&gt;effectively&lt;/em&gt;</description>
      <pubDate>2013-01-16T21:52:35-05:00</pubDate>
    </item>
  </channel>
</rss>`

func TestFeed(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2013-01-16T21:52:35-05:00")
	feed := &Feed{
		Title:       "jmoiron.net blog",
		Link:        &Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &Author{"Jason Moiron", "jmoiron@jmoiron.net"},
		Created:     now,
		Copyright:   "This work is copyright © Benjamin Button",
	}

	feed.Items = []*Item{
		&Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &Author{"Jason Moiron", "jmoiron@jmoiron.net"},
			Created:     now,
		},
		&Item{
			Title:       "Logic-less Template Redux",
			Link:        &Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&Item{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}
	atom, _ := feed.ToAtom()
	rss, _ := feed.ToRss()

	if atom != atomOutput {
		t.Errorf("Atom not what was expected.  Got:\n%s\n\nExpected:\n%s\n", atom, atomOutput)
	}

	if rss != rssOutput {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", rss, rssOutput)
	}
}

func TestParsing(t *testing.T) {

	atomFeed := &AtomFeed{}
	err := xml.Unmarshal([]byte(atomOutput), atomFeed)
	if err != nil {
		t.Log("Error when unmarshaling:", err.Error())
	}

	atom, err := ToXML(atomFeed)
	if err != nil {
		t.Log("Error when remarshalling", err.Error())
	}

	if atom != atomOutput {
		t.Errorf("Atom not what was expected.  Got:\n%s\n\nExpected:\n%s\n", atom, atomOutput)
	}

	rssFeed := &rssFeedXml{}
	err = xml.Unmarshal([]byte(rssOutput), rssFeed)
	if err != nil {
		t.Log("Error when unmarshaling:", err.Error())
	}

	rss, err := ToXML(rssFeed.Channel)
	if err != nil {
		t.Log("Error when remarshalling", err.Error())
	}

	if rss != rssOutput {
		t.Errorf("rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", rss, rssOutput)
	}
}
