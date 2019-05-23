package feeds

// rss support
// validation done according to spec here:
//    http://cyber.law.harvard.edu/rss/rss.html

import (
	"fmt"
	"time"
)

// private wrapper around the RssFeed which gives us the <rss>..</rss> xml
type ItunesRssFeedXml struct {
	*RssFeedXml
	Channel *ItunesRssFeed
}

type ItunesRssContent struct {
	*RssContent
}

type ItunesRssTextInput struct {
	*RssTextInput
}

type ItunesRssFeed struct {
	*RssFeed
	Iimage    string          `xml:"itunes:image"`
	ICategory *ItunesCategory `xml:"itunes:category"`
	TextInput *ItunesRssTextInput
	Items     []*ItunesRssItem `xml:"item"`
}

type ItunesCategory struct {
	Text     string             `xml:"text,attr"`
	Category *ItunesSubCategory `xml:"itunes:category"`
}

type ItunesSubCategory struct {
	Text string `xml:"text,attr"`
}

type ItunesRssItem struct {
	*RssItem
	Content   *ItunesRssContent
	Enclosure *ItunesRssEnclosure
}

type ItunesRssEnclosure struct {
	*RssEnclosure
}

type ItunesRss struct {
	*Feed
}

// create a new ItunesRssItem with a generic Item struct's data
func newItunesRssItem(i *Item) *ItunesRssItem {
	item := &ItunesRssItem{}

	item.Title = i.Title
	item.Link = i.Link.Href
	item.Description = i.Description
	item.Guid = i.Id
	item.PubDate = anyTimeFormat(time.RFC1123Z, i.Created, i.Updated)

	if len(i.Content) > 0 {
		item.Content = &ItunesRssContent{}
		item.Content.Content = i.Content
	}
	if i.Source != nil {
		item.Source = i.Source.Href
	}

	// Define a closure
	if i.Enclosure != nil && i.Enclosure.Type != "" && i.Enclosure.Length != "" {
		item.Enclosure = &ItunesRssEnclosure{}
		item.Enclosure.Url = i.Enclosure.Url
		item.Enclosure.Type = i.Enclosure.Type
		item.Enclosure.Length = i.Enclosure.Length
	}

	if i.Author != nil {
		item.Author = i.Author.Name
	}
	return item
}

// create a new RssFeed with a generic Feed struct's data
func (r *ItunesRss) ItunesRssFeed() *ItunesRssFeed {
	pub := anyTimeFormat(time.RFC1123Z, r.Created, r.Updated)
	build := anyTimeFormat(time.RFC1123Z, r.Updated)
	author := ""
	if r.Author != nil {
		author = r.Author.Email
		if len(r.Author.Name) > 0 {
			author = fmt.Sprintf("%s (%s)", r.Author.Email, r.Author.Name)
		}
	}

	var image *RssImage
	if r.Image != nil {
		image = &RssImage{}
		image.Url = r.Image.Url
		image.Title = r.Image.Title
		image.Link = r.Image.Link
		image.Width = r.Image.Width
		image.Height = r.Image.Height
	}

	channel := &ItunesRssFeed{}

	channel.Title = r.Title
	channel.Link = r.Link.Href
	channel.Description = r.Description
	channel.ManagingEditor = author
	channel.PubDate = pub
	channel.LastBuildDate = build
	channel.Copyright = r.Copyright
	channel.Image = image

	for _, i := range r.Items {
		channel.Items = append(channel.Items, newItunesRssItem(i))
	}
	return channel
}

// FeedXml returns an XML-Ready object for an Rss object
func (r *ItunesRss) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	return r.ItunesRssFeed().FeedXml()

}

// FeedXml returns an XML-ready object for an RssFeed object
func (r *ItunesRssFeed) FeedXml() interface{} {
	itunesRssFeedXml := &ItunesRssFeedXml{}
	itunesRssFeedXml.Version = "2.0"
	itunesRssFeedXml.Channel = r
	itunesRssFeedXml.ContentNamespace = "http://purl.org/rss/1.0/modules/content/"
	return itunesRssFeedXml
}
