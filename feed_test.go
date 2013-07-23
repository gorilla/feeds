package feeds

import (
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

var atomInput = `<?xml version="1.0" encoding="utf-8"?><feed xmlns="http://www.w3.org/2005/Atom"><title>Event stream 'bf24aec7-7357-4194-492b-263090208d78'</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78</id><updated>2013-07-22T06:16:53.323897Z</updated><author><name>EventStore</name></author><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78" rel="self" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/head/backward/20" rel="first" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/0/forward/20" rel="last" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/80/backward/20" rel="next" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/101/forward/20" rel="previous" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/metadata" rel="metadata" /><entry><title>100@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/100</id><updated>2013-07-22T06:16:53.323897Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/100" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/100" rel="alternate" /></entry><entry><title>99@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/99</id><updated>2013-07-22T06:16:53.323891Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/99" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/99" rel="alternate" /></entry><entry><title>98@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/98</id><updated>2013-07-22T06:16:53.323885Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/98" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/98" rel="alternate" /></entry><entry><title>97@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/97</id><updated>2013-07-22T06:16:53.323878Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/97" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/97" rel="alternate" /></entry><entry><title>96@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/96</id><updated>2013-07-22T06:16:53.323871Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/96" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/96" rel="alternate" /></entry><entry><title>95@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/95</id><updated>2013-07-22T06:16:53.323865Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/95" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/95" rel="alternate" /></entry><entry><title>94@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/94</id><updated>2013-07-22T06:16:53.323859Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/94" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/94" rel="alternate" /></entry><entry><title>93@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/93</id><updated>2013-07-22T06:16:53.323852Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/93" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/93" rel="alternate" /></entry><entry><title>92@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/92</id><updated>2013-07-22T06:16:53.323846Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/92" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/92" rel="alternate" /></entry><entry><title>91@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/91</id><updated>2013-07-22T06:16:53.32384Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/91" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/91" rel="alternate" /></entry><entry><title>90@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/90</id><updated>2013-07-22T06:16:53.323834Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/90" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/90" rel="alternate" /></entry><entry><title>89@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/89</id><updated>2013-07-22T06:16:53.323827Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/89" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/89" rel="alternate" /></entry><entry><title>88@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/88</id><updated>2013-07-22T06:16:53.323821Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/88" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/88" rel="alternate" /></entry><entry><title>87@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/87</id><updated>2013-07-22T06:16:53.323813Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/87" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/87" rel="alternate" /></entry><entry><title>86@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/86</id><updated>2013-07-22T06:16:53.323803Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/86" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/86" rel="alternate" /></entry><entry><title>85@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/85</id><updated>2013-07-22T06:16:53.323797Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/85" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/85" rel="alternate" /></entry><entry><title>84@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/84</id><updated>2013-07-22T06:16:53.323791Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/84" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/84" rel="alternate" /></entry><entry><title>83@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/83</id><updated>2013-07-22T06:16:53.323784Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/83" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/83" rel="alternate" /></entry><entry><title>82@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/82</id><updated>2013-07-22T06:16:53.323778Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/82" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/82" rel="alternate" /></entry><entry><title>81@bf24aec7-7357-4194-492b-263090208d78</title><id>http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/81</id><updated>2013-07-22T06:16:53.323771Z</updated><author><name>EventStore</name></author><summary>github.com/pjvds/go-cqrs/tests/events/UsernameChanged</summary><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/81" rel="edit" /><link href="http://localhost:2113/streams/bf24aec7-7357-4194-492b-263090208d78/81" rel="alternate" /></entry></feed>`

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

func TestAtomFeedParsing(t *testing.T) {
	atom, err := ParseAtomFeed(atomInput)
	if err != nil {
		t.Errorf("Unexpected error while parsing atom: %v", err)
	}

	if title := atom.Title; title != "Event stream 'bf24aec7-7357-4194-492b-263090208d78'" {
		t.Errorf("unexpected title: %v", title)
	}

	if l := len(atom.Entries); l > 0 {
		t.Errorf("Unexpected entries count: %v", l)
	}
}

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
