package feeds

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

var rssOutput1 = `<?xml version="1.0" encoding="UTF-8"?><rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/">
  <channel test="https://example.com/test/uri">
    <title>Example</title>
    <link>https://rss.example.com</link>
    <description>Example Example</description>
    <copyright>copyright © Example</copyright>
    <managingEditor>Example (Example)</managingEditor>
    <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    <item>
      <title>Example example</title>
      <link>https://example.com/link</link>
      <description>example</description>
      <content:encoded><![CDATA[Example example example, example example...]]></content:encoded>
      <author>Jason Moiron</author>
      <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
      <test:version>v1</test:version>
      <test:example>1234567890</test:example>
    </item>
  </channel>
</rss>`

var rssOutput2 = `<?xml version="1.0" encoding="UTF-8"?><rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/">
  <channel test="https://example.com/test/uri">
    <title>Example</title>
    <link>https://rss.example.com</link>
    <description>Example Example</description>
    <copyright>copyright © Example</copyright>
    <managingEditor>Example (Example)</managingEditor>
    <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
    <item>
      <title>Example example</title>
      <link>https://example.com/link</link>
      <description>example</description>
      <content:encoded><![CDATA[Example example example, example example...]]></content:encoded>
      <author>Jason Moiron</author>
      <pubDate>Wed, 16 Jan 2013 21:52:35 -0500</pubDate>
      <test:num>111</test:num>
      <test:example>
        <version>v1</version>
      </test:example>
    </item>
  </channel>
</rss>`

func TestRssFeedExtensions(t *testing.T) {

	now, err := time.Parse(time.RFC3339, "2013-01-16T21:52:35-05:00")
	if err != nil {
		t.Error(err)
	}
	tz := time.FixedZone("EST", -5*60*60)
	now = now.In(tz)

	feed := &Feed{
		Title:       "Example",
		Link:        &Link{Href: "https://rss.example.com"},
		Description: "Example Example",
		Author:      &Author{Name: "Example", Email: "Example"},
		Created:     now,
		Copyright:   "copyright © Example",
	}

	feed.Items = []*Item{
		{
			Title:       "Example example",
			Link:        &Link{Href: "https://example.com/link"},
			Description: "example",
			Author:      &Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
			Content:     "Example example example, example example...",
		},
	}

	feed.AddAttribute("test", "https://example.com/test/uri")
	feed.Items[0].AddExtensionString("test:version", "", "v1")
	feed.Items[0].AddExtensionInt("test:example", "", 1234567890)

	rss, err := feed.ToRss()
	if err != nil {
		t.Errorf("unexpected error encoding RSS: %v", err)
	}
	if rss != rssOutput1 {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", rss, rssOutput1)
	}

	var buf = new(bytes.Buffer)
	if err := feed.WriteRss(buf); err != nil {
		t.Errorf("unexpected error writing RSS: %v", err)
	}
	if got := buf.String(); got != rssOutput1 {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", got, rssOutput1)
	}
}

type extern struct {
	XMLName xml.Name `xml:"test:example"`
	Version string   `xml:"version"`
}

func TestRssFeedExtensions2(t *testing.T) {

	now, err := time.Parse(time.RFC3339, "2013-01-16T21:52:35-05:00")
	if err != nil {
		t.Error(err)
	}
	tz := time.FixedZone("EST", -5*60*60)
	now = now.In(tz)

	feed := &Feed{
		Title:       "Example",
		Link:        &Link{Href: "https://rss.example.com"},
		Description: "Example Example",
		Author:      &Author{Name: "Example", Email: "Example"},
		Created:     now,
		Copyright:   "copyright © Example",
	}

	feed.Items = []*Item{
		{
			Title:       "Example example",
			Link:        &Link{Href: "https://example.com/link"},
			Description: "example",
			Author:      &Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
			Content:     "Example example example, example example...",
		},
	}

	ex := extern{
		Version: "v1",
	}

	feed.AddAttribute("test", "https://example.com/test/uri")
	feed.Items[0].AddExtensionUint("test:num", "", 111)
	feed.Items[0].AddExtension(ex)

	rss, err := feed.ToRss()
	if err != nil {
		t.Errorf("unexpected error encoding RSS: %v", err)
	}
	if rss != rssOutput2 {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", rss, rssOutput2)
	}

	var buf = new(bytes.Buffer)
	if err := feed.WriteRss(buf); err != nil {
		t.Errorf("unexpected error writing RSS: %v", err)
	}
	if got := buf.String(); got != rssOutput2 {
		t.Errorf("Rss not what was expected.  Got:\n%s\n\nExpected:\n%s\n", got, rssOutput2)
	}
}
