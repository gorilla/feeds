package syndicate

import (
	"fmt"
	"testing"
	"time"
)

func TestFeed(t *testing.T) {
	feed := &Feed{
		Title:       "jmoiron.net blog",
		Link:        &Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &Author{"Jason Moiron", "jmoiron@jmoiron.net"},
		Created:     time.Now(),
	}

	fmt.Println(feed.ToAtom())
	fmt.Println(feed.ToRss())
}
