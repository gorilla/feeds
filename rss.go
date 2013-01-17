package syndicate

import (
	"encoding/xml"
	"fmt"
	//"net/url"
	"time"
)

// private wrapper around the RssFeed which gives us the <rss>..</rss> xml
type rssFeedXml struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *RssFeed
}

// rss feed object which 
type RssFeed struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`       // required
	Link           string   `xml:"link"`        // required
	Description    string   `xml:"description"` // required
	Language       string   `xml:"language,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty"` // Author used
	WebMaster      string   `xml:"webMaster,omitempty"`
	PubDate        string   `xml:"pubDate,omitempty"`       // created or updated
	LastBuildDate  string   `xml:"lastBuildDate,omitempty"` // updated used
	Category       string   `xml:"category,omitempty"`
	Generator      string   `xml:"generator,omitempty"`
	Docs           string   `xml:"docs,omitempty"`
	Cloud          string   `xml:"cloud,omitempty"`
	Ttl            string   `xml:"ttl,omitempty"`
	Image          string   `xml:"image,omitempty"`
	Rating         string   `xml:"rating,omitempty"`
	TextInput      string   `xml:"textInput,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"`
	SkipDays       string   `xml:"skipDays,omitempty"`
	Items          []*RssItem
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`       // required
	Link        string   `xml:"link"`        // required
	Description string   `xml:"description"` // required
	Author      string   `xml:"author,omitempty"`
	Category    string   `xml:"category,omitempty"`
	Comments    string   `xml:"comments,omitempty"`
	Enclosure   *RssEnclosure
	Guid        string `xml:"guid,omitempty"`    // Id used
	PubDate     string `xml:"pubDate,omitempty"` // created or updated
	Source      string `xml:"source,omitempty"`
}

type RssEnclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	Url     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

type Rss struct {
	*Feed
}

func newRssItem(i *Item) *RssItem {
	item := &RssItem{
		Title:       i.Title,
		Link:        i.Link.Href,
		Description: i.Description,
		Guid:        i.Id,
		PubDate:     anyTimeFormat(time.RFC3339, i.Created, i.Updated),
	}
	return item
}

func (r *Rss) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	pub := anyTimeFormat(time.RFC3339, r.Created, r.Updated)
	build := anyTimeFormat(time.RFC3339, r.Updated)
	author := r.Author.Email
	if len(r.Author.Name) > 0 {
		author = fmt.Sprintf("%s (%s)", r.Author.Email, r.Author.Name)
	}
	feed := &rssFeedXml{
		Version: "2.0",
		Channel: &RssFeed{
			Title:          r.Title,
			Link:           r.Link.Href,
			Description:    r.Description,
			ManagingEditor: author,
			PubDate:        pub,
			LastBuildDate:  build,
		},
	}
	for _, i := range r.Items {
		feed.Channel.Items = append(feed.Channel.Items, newRssItem(i))
	}
	return feed

}
