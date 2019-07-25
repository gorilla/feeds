package feeds

import (
	"bytes"
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
    <content type="html">&lt;p&gt;Go&#39;s goroutines make it easy to make &lt;a href=&#34;http://collectiveidea.com/blog/archives/2012/12/03/playing-with-go-embarrassingly-parallel-scripts/&#34;&gt;embarrassingly parallel programs&lt;/a&gt;, but in many &amp;quot;real world&amp;quot; cases resources can be limited and attempting to do everything at once can exhaust your access to them.&lt;/p&gt;</content>
    <link href="http://jmoiron.net/blog/limiting-concurrency-in-go/" rel="alternate"></link>
    <summary type="html">A discussion on controlled parallelism in golang</summary>
    <author>
      <name>Jason Moiron</name>
      <email>jmoiron@jmoiron.net</email>
    </author>
  </entry>
  <entry>
    <title>Logic-less Template Redux</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/blog/logicless-template-redux/</id>
    <link href="http://jmoiron.net/blog/logicless-template-redux/" rel="alternate"></link>
    <summary type="html">More thoughts on logicless templates</summary>
  </entry>
  <entry>
    <title>Idiomatic Code Reuse in Go</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-16:/blog/idiomatic-code-reuse-in-go/</id>
    <link href="http://jmoiron.net/blog/idiomatic-code-reuse-in-go/" rel="alternate"></link>
    <link href="http://example.com/cover.jpg" rel="enclosure" type="image/jpg" length="123456"></link>
    <summary type="html">How to use interfaces &lt;em&gt;effectively&lt;/em&gt;</summary>
  </entry>
  <entry>
    <title>Never Gonna Give You Up Mp3</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:example.com,2013-01-16:/RickRoll.mp3</id>
    <link href="http://example.com/RickRoll.mp3" rel="alternate"></link>
    <link href="http://example.com/RickRoll.mp3" rel="enclosure" type="audio/mpeg" length="123456"></link>
    <summary type="html">Never gonna give you up - Never gonna let you down.</summary>
  </entry>
  <entry>
    <title>String formatting in Go</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:example.com,2013-01-16:/strings</id>
    <link href="http://example.com/strings" rel="alternate"></link>
    <summary type="html">How to use things like %s, %v, %d, etc.</summary>
  </entry>
</feed>`

var rssOutput = `<?xml version="1.0" encoding="UTF-8"?><rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/">
  <channel>
    <title>jmoiron.net blog</title>
    <link>http://jmoiron.net/blog</link>
    <description>discussion about tech, footie, photos</description>
    <copyright>This work is copyright © Benjamin Button</copyright>
    <managingEditor>jmoiron@jmoiron.net (Jason Moiron)</managingEditor>
    <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    <item>
      <title>Limiting Concurrency in Go</title>
      <link>http://jmoiron.net/blog/limiting-concurrency-in-go/</link>
      <description>A discussion on controlled parallelism in golang</description>
      <content:encoded><![CDATA[<p>Go's goroutines make it easy to make <a href="http://collectiveidea.com/blog/archives/2012/12/03/playing-with-go-embarrassingly-parallel-scripts/">embarrassingly parallel programs</a>, but in many &quot;real world&quot; cases resources can be limited and attempting to do everything at once can exhaust your access to them.</p>]]></content:encoded>
      <author>Jason Moiron</author>
      <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    </item>
    <item>
      <title>Logic-less Template Redux</title>
      <link>http://jmoiron.net/blog/logicless-template-redux/</link>
      <description>More thoughts on logicless templates</description>
      <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    </item>
    <item>
      <title>Idiomatic Code Reuse in Go</title>
      <link>http://jmoiron.net/blog/idiomatic-code-reuse-in-go/</link>
      <description>How to use interfaces &lt;em&gt;effectively&lt;/em&gt;</description>
      <enclosure url="http://example.com/cover.jpg" length="123456" type="image/jpg"></enclosure>
      <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    </item>
    <item>
      <title>Never Gonna Give You Up Mp3</title>
      <link>http://example.com/RickRoll.mp3</link>
      <description>Never gonna give you up - Never gonna let you down.</description>
      <enclosure url="http://example.com/RickRoll.mp3" length="123456" type="audio/mpeg"></enclosure>
      <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    </item>
    <item>
      <title>String formatting in Go</title>
      <link>http://example.com/strings</link>
      <description>How to use things like %s, %v, %d, etc.</description>
      <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    </item>
  </channel>
</rss>`

var jsonOutput = `{
  "version": "https://jsonfeed.org/version/1",
  "title": "jmoiron.net blog",
  "home_page_url": "http://jmoiron.net/blog",
  "description": "discussion about tech, footie, photos",
  "author": {
    "name": "Jason Moiron"
  },
  "items": [
    {
      "id": "",
      "url": "http://jmoiron.net/blog/limiting-concurrency-in-go/",
      "title": "Limiting Concurrency in Go",
      "content_html": "\u003cp\u003eGo's goroutines make it easy to make \u003ca href=\"http://collectiveidea.com/blog/archives/2012/12/03/playing-with-go-embarrassingly-parallel-scripts/\"\u003eembarrassingly parallel programs\u003c/a\u003e, but in many \u0026quot;real world\u0026quot; cases resources can be limited and attempting to do everything at once can exhaust your access to them.\u003c/p\u003e",
      "summary": "A discussion on controlled parallelism in golang",
      "date_published": "2013-01-16T21:52:35-05:00",
      "author": {
        "name": "Jason Moiron"
      }
    },
    {
      "id": "",
      "url": "http://jmoiron.net/blog/logicless-template-redux/",
      "title": "Logic-less Template Redux",
      "summary": "More thoughts on logicless templates",
      "date_published": "2013-01-16T21:52:35-05:00"
    },
    {
      "id": "",
      "url": "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/",
      "title": "Idiomatic Code Reuse in Go",
      "summary": "How to use interfaces \u003cem\u003eeffectively\u003c/em\u003e",
      "image": "http://example.com/cover.jpg",
      "date_published": "2013-01-16T21:52:35-05:00"
    },
    {
      "id": "",
      "url": "http://example.com/RickRoll.mp3",
      "title": "Never Gonna Give You Up Mp3",
      "summary": "Never gonna give you up - Never gonna let you down.",
      "date_published": "2013-01-16T21:52:35-05:00"
    },
    {
      "id": "",
      "url": "http://example.com/strings",
      "title": "String formatting in Go",
      "summary": "How to use things like %s, %v, %d, etc.",
      "date_published": "2013-01-16T21:52:35-05:00"
    }
  ]
}`

func TestFeed(t *testing.T) {
	now, err := time.Parse(time.RFC3339, "2013-01-16T21:52:35-05:00")
	if err != nil {
		t.Error(err)
	}
	tz := time.FixedZone("EST", -5*60*60)
	now = now.In(tz)

	feed := &Feed{
		Title:       "jmoiron.net blog",
		Link:        &Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
		Copyright:   "This work is copyright © Benjamin Button",
	}

	feed.Items = []*Item{
		{
			Title:       "Limiting Concurrency in Go",
			Link:        &Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
			Content:     `<p>Go's goroutines make it easy to make <a href="http://collectiveidea.com/blog/archives/2012/12/03/playing-with-go-embarrassingly-parallel-scripts/">embarrassingly parallel programs</a>, but in many &quot;real world&quot; cases resources can be limited and attempting to do everything at once can exhaust your access to them.</p>`,
		},
		{
			Title:       "Logic-less Template Redux",
			Link:        &Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Enclosure:   &Enclosure{Url: "http://example.com/cover.jpg", Length: "123456", Type: "image/jpg"},
			Created:     now,
		},
		{
			Title:       "Never Gonna Give You Up Mp3",
			Link:        &Link{Href: "http://example.com/RickRoll.mp3"},
			Enclosure:   &Enclosure{Url: "http://example.com/RickRoll.mp3", Length: "123456", Type: "audio/mpeg"},
			Description: "Never gonna give you up - Never gonna let you down.",
			Created:     now,
		},
		{
			Title:       "String formatting in Go",
			Link:        &Link{Href: "http://example.com/strings"},
			Description: "How to use things like %s, %v, %d, etc.",
			Created:     now,
		}}

	atom, err := feed.ToAtom()
	if err != nil {
		t.Errorf("unexpected error encoding Atom: %v", err)
	}
	if atom != atomOutput {
		t.Errorf("Atom not what was expected.  Got:\n%s\n\nExpected:\n%s\n", atom, atomOutput)
	}
	var buf bytes.Buffer
	if err := feed.WriteAtom(&buf); err != nil {
		t.Errorf("unexpected error writing Atom: %v", err)
	}
	if got := buf.String(); got != atomOutput {
		t.Errorf("Atom not what was expected.  Got:\n%s\n\nExpected:\n%s\n", got, atomOutput)
	}

	rss, err := feed.ToRss()
	if err != nil {
		t.Errorf("unexpected error encoding RSS: %v", err)
	}
	if rss != rssOutput {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", rss, rssOutput)
	}
	buf.Reset()
	if err := feed.WriteRss(&buf); err != nil {
		t.Errorf("unexpected error writing RSS: %v", err)
	}
	if got := buf.String(); got != rssOutput {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", got, rssOutput)
	}

	json, err := feed.ToJSON()
	if err != nil {
		t.Errorf("unexpected error encoding JSON: %v", err)
	}
	if json != jsonOutput {
		t.Errorf("JSON not what was expected.  Got:\n%s\n\nExpected:\n%s\n", json, jsonOutput)
	}
	buf.Reset()
	if err := feed.WriteJSON(&buf); err != nil {
		t.Errorf("unexpected error writing JSON: %v", err)
	}
	if got := buf.String(); got != jsonOutput+"\n" { //json.Encode appends a newline after the JSON output: https://github.com/golang/go/commit/6f25f1d4c901417af1da65e41992d71c30f64f8f#diff-50848cbd686f250623a2ef6ddb07e157
		t.Errorf("JSON not what was expected.  Got:\n||%s||\n\nExpected:\n||%s||\n", got, jsonOutput)
	}
}

var atomOutputSorted = `<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
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
    <updated>2013-01-18T21:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-18:/blog/limiting-concurrency-in-go/</id>
    <link href="http://jmoiron.net/blog/limiting-concurrency-in-go/" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>Logic-less Template Redux</title>
    <updated>2013-01-17T21:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-17:/blog/logicless-template-redux/</id>
    <link href="http://jmoiron.net/blog/logicless-template-redux/" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>Idiomatic Code Reuse in Go</title>
    <updated>2013-01-17T09:52:35-05:00</updated>
    <id>tag:jmoiron.net,2013-01-17:/blog/idiomatic-code-reuse-in-go/</id>
    <link href="http://jmoiron.net/blog/idiomatic-code-reuse-in-go/" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>Never Gonna Give You Up Mp3</title>
    <updated>2013-01-17T07:52:35-05:00</updated>
    <id>tag:example.com,2013-01-17:/RickRoll.mp3</id>
    <link href="http://example.com/RickRoll.mp3" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
  <entry>
    <title>String formatting in Go</title>
    <updated>2013-01-16T21:52:35-05:00</updated>
    <id>tag:example.com,2013-01-16:/strings</id>
    <link href="http://example.com/strings" rel="alternate"></link>
    <summary type="html"></summary>
  </entry>
</feed>`

var rssOutputSorted = `<?xml version="1.0" encoding="UTF-8"?><rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/">
  <channel>
    <title>jmoiron.net blog</title>
    <link>http://jmoiron.net/blog</link>
    <description>discussion about tech, footie, photos</description>
    <copyright>This work is copyright © Benjamin Button</copyright>
    <managingEditor>jmoiron@jmoiron.net (Jason Moiron)</managingEditor>
    <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    <item>
      <title>Limiting Concurrency in Go</title>
      <link>http://jmoiron.net/blog/limiting-concurrency-in-go/</link>
      <description></description>
      <pubDate>Fri, 18 Jan 2013 21:52:35 -0500</pubDate>
    </item>
    <item>
      <title>Logic-less Template Redux</title>
      <link>http://jmoiron.net/blog/logicless-template-redux/</link>
      <description></description>
      <pubDate>Thu, 17 Jan 2013 21:52:35 -0500</pubDate>
    </item>
    <item>
      <title>Idiomatic Code Reuse in Go</title>
      <link>http://jmoiron.net/blog/idiomatic-code-reuse-in-go/</link>
      <description></description>
      <pubDate>Thu, 17 Jan 2013 09:52:35 -0500</pubDate>
    </item>
    <item>
      <title>Never Gonna Give You Up Mp3</title>
      <link>http://example.com/RickRoll.mp3</link>
      <description></description>
      <pubDate>Thu, 17 Jan 2013 07:52:35 -0500</pubDate>
    </item>
    <item>
      <title>String formatting in Go</title>
      <link>http://example.com/strings</link>
      <description></description>
      <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    </item>
  </channel>
</rss>`

var jsonOutputSorted = `{
  "version": "https://jsonfeed.org/version/1",
  "title": "jmoiron.net blog",
  "home_page_url": "http://jmoiron.net/blog",
  "description": "discussion about tech, footie, photos",
  "author": {
    "name": "Jason Moiron"
  },
  "items": [
    {
      "id": "",
      "url": "http://jmoiron.net/blog/limiting-concurrency-in-go/",
      "title": "Limiting Concurrency in Go",
      "date_published": "2013-01-18T21:52:35-05:00"
    },
    {
      "id": "",
      "url": "http://jmoiron.net/blog/logicless-template-redux/",
      "title": "Logic-less Template Redux",
      "date_published": "2013-01-17T21:52:35-05:00"
    },
    {
      "id": "",
      "url": "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/",
      "title": "Idiomatic Code Reuse in Go",
      "date_published": "2013-01-17T09:52:35-05:00"
    },
    {
      "id": "",
      "url": "http://example.com/RickRoll.mp3",
      "title": "Never Gonna Give You Up Mp3",
      "date_published": "2013-01-17T07:52:35-05:00"
    },
    {
      "id": "",
      "url": "http://example.com/strings",
      "title": "String formatting in Go",
      "date_published": "2013-01-16T21:52:35-05:00"
    }
  ]
}`

func TestFeedSorted(t *testing.T) {
	now, err := time.Parse(time.RFC3339, "2013-01-16T21:52:35-05:00")
	if err != nil {
		t.Error(err)
	}
	tz := time.FixedZone("EST", -5*60*60)
	now = now.In(tz)

	feed := &Feed{
		Title:       "jmoiron.net blog",
		Link:        &Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
		Copyright:   "This work is copyright © Benjamin Button",
	}

	feed.Items = []*Item{
		{
			Title:   "Limiting Concurrency in Go",
			Link:    &Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Created: now.Add(time.Duration(time.Hour * 48)),
		},
		{
			Title:   "Logic-less Template Redux",
			Link:    &Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Created: now.Add(time.Duration(time.Hour * 24)),
		},
		{
			Title:   "Idiomatic Code Reuse in Go",
			Link:    &Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Created: now.Add(time.Duration(time.Hour * 12)),
		},
		{
			Title:   "Never Gonna Give You Up Mp3",
			Link:    &Link{Href: "http://example.com/RickRoll.mp3"},
			Created: now.Add(time.Duration(time.Hour * 10)),
		},
		{
			Title:   "String formatting in Go",
			Link:    &Link{Href: "http://example.com/strings"},
			Created: now,
		}}

	feed.Sort(func(a, b *Item) bool {
		return a.Created.After(b.Created)
	})
	atom, err := feed.ToAtom()
	if err != nil {
		t.Errorf("unexpected error encoding Atom: %v", err)
	}
	if atom != atomOutputSorted {
		t.Errorf("Atom not what was expected.  Got:\n%s\n\nExpected:\n%s\n", atom, atomOutputSorted)
	}
	var buf bytes.Buffer
	if err := feed.WriteAtom(&buf); err != nil {
		t.Errorf("unexpected error writing Atom: %v", err)
	}
	if got := buf.String(); got != atomOutputSorted {
		t.Errorf("Atom not what was expected.  Got:\n%s\n\nExpected:\n%s\n", got, atomOutputSorted)
	}

	rss, err := feed.ToRss()
	if err != nil {
		t.Errorf("unexpected error encoding RSS: %v", err)
	}

	if rss != rssOutputSorted {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", rss, rssOutputSorted)
	}
	buf.Reset()
	if err := feed.WriteRss(&buf); err != nil {
		t.Errorf("unexpected error writing RSS: %v", err)
	}
	if got := buf.String(); got != rssOutputSorted {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", got, rssOutputSorted)
	}

	json, err := feed.ToJSON()
	if err != nil {
		t.Errorf("unexpected error encoding JSON: %v", err)
	}
	if json != jsonOutputSorted {
		t.Errorf("JSON not what was expected.  Got:\n%s\n\nExpected:\n%s\n", json, jsonOutputSorted)
	}
	buf.Reset()
	if err := feed.WriteJSON(&buf); err != nil {
		t.Errorf("unexpected error writing JSON: %v", err)
	}
	if got := buf.String(); got != jsonOutputSorted+"\n" { //json.Encode appends a newline after the JSON output: https://github.com/golang/go/commit/6f25f1d4c901417af1da65e41992d71c30f64f8f#diff-50848cbd686f250623a2ef6ddb07e157
		t.Errorf("JSON not what was expected.  Got:\n||%s||\n\nExpected:\n||%s||\n", got, jsonOutputSorted)
	}
}

var jsonOutputHub = `{
  "version": "https://jsonfeed.org/version/1",
  "title": "feed title",
  "hubs": [
    {
      "type": "WebSub",
      "url": "https://websub-hub.example"
    }
  ]
}`

func TestJSONHub(t *testing.T) {
	feed := &JSONFeed{
		Version: "https://jsonfeed.org/version/1",
		Title:   "feed title",
		Hubs: []*JSONHub{
			&JSONHub{
				Type: "WebSub",
				Url:  "https://websub-hub.example",
			},
		},
	}
	json, err := feed.ToJSON()
	if err != nil {
		t.Errorf("unexpected error encoding JSON: %v", err)
	}
	if json != jsonOutputHub {
		t.Errorf("JSON not what was expected.  Got:\n%s\n\nExpected:\n%s\n", json, jsonOutputHub)
	}
}
