## gorilla/feeds

Web feed generator library.

[![Build Status](https://travis-ci.org/gorilla/feeds.png?branch=master)](https://travis-ci.org/gorilla/feeds)

### Goals

 * simple interface to create both Atom & RSS 2.0 feeds
 * full support for Atom/RSS2.0 spec elements
 * ability to modify particulars for each spec

### Usage

```go

package main

import (
	"fmt"
	"time"

	"github.com/gorilla/feeds"
)

func main() {

	now := time.Now()
	feed := &feeds.Feed{
		Title:       "jmoiron.net blog",
		Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
			Created:     now,
		},
		&feeds.Item{
			Title:       "Logic-less Template Redux",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&feeds.Item{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}

	atom, _ := feed.ToAtom()
	rss, _ := feed.ToRss()

	fmt.Println(atom, "\n", rss)

}

```

