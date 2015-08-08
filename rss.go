package feeds

// rss support
// validation done according to spec here:
//    http://cyber.law.harvard.edu/rss/rss.html

import (
	"encoding/xml"
	"fmt"
	"time"
)

// private wrapper around the RssFeed which gives us the <rss>..</rss> xml
type rssFeedXml struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *RssFeed
}

type RssImage struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Width   int      `xml:"width,omitempty"`
	Height  int      `xml:"height,omitempty"`
}

type RssTextInput struct {
	XMLName     xml.Name `xml:"textInput"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Name        string   `xml:"name"`
	Link        string   `xml:"link"`
}

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
	Ttl            int      `xml:"ttl,omitempty"`
	Rating         string   `xml:"rating,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"`
	SkipDays       string   `xml:"skipDays,omitempty"`
	Image          *RssImage
	TextInput      *RssTextInput
	Items          []*RssItem
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"` // required
	Link        string   `xml:"link,omitempty"`
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

// create a new RssItem with a generic Item struct's data
func newRssItem(i *Item) *RssItem {
	item := &RssItem{
		Title:       i.Title,
		Description: i.Description,
		Guid:        i.Id,
		PubDate:     anyTimeFormat(time.RFC822, i.Created, i.Updated),
	}
	if i.Author != nil {
		item.Author = i.Author.Name
	}
	if i.Link != nil {
		item.Link = i.Link.Href
	}
	if i.Enclosure != nil {
		item.Enclosure = &RssEnclosure{
			Url:    i.Enclosure.Url,
			Length: i.Enclosure.Length,
			Type:   i.Enclosure.Type,
		}
	}

	return item
}

// create a new RssFeed with a generic Feed struct's data
func (r *Rss) RssFeed() *RssFeed {
	pub := anyTimeFormat(time.RFC822, r.Created, r.Updated)
	build := anyTimeFormat(time.RFC822, r.Updated)
	author := ""
	if r.Author != nil {
		author = r.Author.Email
		if len(r.Author.Name) > 0 {
			author = fmt.Sprintf("%s (%s)", r.Author.Email, r.Author.Name)
		}
	}

	channel := &RssFeed{
		Title:          r.Title,
		Link:           r.Link.Href,
		Description:    r.Description,
		ManagingEditor: author,
		PubDate:        pub,
		LastBuildDate:  build,
		Copyright:      r.Copyright,
	}
	for _, i := range r.Items {
		channel.Items = append(channel.Items, newRssItem(i))
	}
	return channel
}

// return an XML-Ready object for an Rss object
func (r *Rss) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	return r.RssFeed().FeedXml()

}

// return an XML-ready object for an RssFeed object
func (r *RssFeed) FeedXml() interface{} {
	return &rssFeedXml{Version: "2.0", Channel: r}
}
