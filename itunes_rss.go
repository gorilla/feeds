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
	Channel         *ItunesRssFeed
	ItunesNamespace string `xml:"xmlns:itunes,attr"`
}

type ItunesRssFeed struct {
	*RssFeed
	Iimage      string           `xml:"itunes:image"`
	ICategory   *ItunesCategory  `xml:"itunes:category"`
	IExplicit   string           `xml:"itunes:explicit"`
	IAuthor     string           `xml:"itunes:author,omitempty"`
	IOwner      *ItunesOwner     `xml:"itunes:owner,omitempty"`
	IType       string           `xml:"itunes:type,omitempty"`
	INewFeedUrl string           `xml:"itunes:new-feed-url,omitempty"`
	IBlock      string           `xml:"itunes:block,omitempty"`
	IComplete   string           `xml:"itunes:complete,omitempty"`
	Items       []*ItunesRssItem `xml:"item,omitempty"`
}

type ItunesCategory struct {
	Text     string             `xml:"text,attr"`
	Category *ItunesSubCategory `xml:"itunes:category"`
}

type ItunesSubCategory struct {
	Text string `xml:"text,attr"`
}

type ItunesOwner struct {
	Name  string `xml:"itunes:name,omitempty"`
	Email string `xml:"itunes:email,omitempty"`
}

type ItunesRssItem struct {
	*RssItem
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
		item.Content = &RssContent{Content: i.Content}
	}
	if i.Source != nil {
		item.Source = i.Source.Href
	}

	// Define a closure
	if i.Enclosure != nil && i.Enclosure.Type != "" && i.Enclosure.Length != "" {
		item.Enclosure = &RssEnclosure{Url: i.Enclosure.Url, Type: i.Enclosure.Type, Length: i.Enclosure.Length}
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
		image = &RssImage{Url: r.Image.Url, Title: r.Image.Title, Link: r.Image.Link, Width: r.Image.Width, Height: r.Image.Height}
	}

	channelRss := &RssFeed{
		Title:          r.Title,
		Link:           r.Link.Href,
		Description:    r.Description,
		ManagingEditor: author,
		PubDate:        pub,
		LastBuildDate:  build,
		Copyright:      r.Copyright,
		Image:          image,
	}

	channel := &ItunesRssFeed{RssFeed: channelRss}

	// for _, i := range r.Items {
	// 	channel.Items = append(channel.Items, newItunesRssItem(i))
	// }
	return channel
}

// FeedXml returns an XML-Ready object for an Rss object
func (r *ItunesRss) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	return r.ItunesRssFeed().FeedXml()

}

// FeedXml returns an XML-ready object for an RssFeed object
func (r *ItunesRssFeed) FeedXml() interface{} {
	rssFeedXml := &RssFeedXml{
		Version:          "2.0",
		ContentNamespace: "http://purl.org/rss/1.0/modules/content/",
	}
	return ItunesRssFeedXml{rssFeedXml, r, "http://www.itunes.com/dtds/podcast-1.0.dtd"}
}
