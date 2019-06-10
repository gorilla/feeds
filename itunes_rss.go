package feeds

// itunes rss support
// validation done according to spec here:
//    https://help.apple.com/itc/podcasts_connect/#/itcb54353390

import (
	"encoding/xml"
	"fmt"
	"time"
)

// private wrapper around the RssFeed which gives us the <rss>..</rss> xml
type ItunesRssFeedXml struct {
	XMLName          xml.Name `xml:"rss"`
	Version          string   `xml:"version,attr"`
	ContentNamespace string   `xml:"xmlns:content,attr"`
	Channel          *ItunesRssFeed
	ItunesNamespace  string `xml:"xmlns:itunes,attr"`
}

type ItunesRssFeed struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`       // required
	Link           string   `xml:"link"`        // required
	Description    string   `xml:"description"` // required
	Language       string   `xml:"language,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"`
	ManagingEditor string   `xml:"author,omitempty"` // Author used
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
	IImage         *ItunesImage `xml:"itunes:image"`
	ICategory      *ItunesCategory
	IExplicit      string `xml:"itunes:explicit"`
	IAuthor        string `xml:"itunes:author,omitempty"`
	IOwner         *ItunesOwner
	IType          string           `xml:"itunes:type,omitempty"`
	INewFeedUrl    string           `xml:"itunes:new-feed-url,omitempty"`
	IBlock         string           `xml:"itunes:block,omitempty"`
	IComplete      string           `xml:"itunes:complete,omitempty"`
	Items          []*ItunesRssItem `xml:"item,omitempty"`
}

type ItunesImage struct {
	Href string `xml:"href,attr"`
}

type ItunesCategory struct {
	XMLName  xml.Name `xml:"itunes:category"`
	Text     string   `xml:"text,attr"`
	Category *ItunesSubCategory
}

type ItunesSubCategory struct {
	XMLName xml.Name `xml:"itunes:category"`
	Text    string   `xml:"text,attr"`
}

type ItunesOwner struct {
	XMLName xml.Name `xml:"itunes:owner"`
	Name    string   `xml:"itunes:name,omitempty"`
	Email   string   `xml:"itunes:email,omitempty"`
}

type ItunesRssItem struct {
	XMLName      xml.Name           `xml:"item"`
	Title        string             `xml:"title"` // required
	Link         string             `xml:"link"`  // required
	Description  *ItunesDescription // required
	Content      *RssContent
	Author       string `xml:"author,omitempty"`
	Category     string `xml:"category,omitempty"`
	Comments     string `xml:"comments,omitempty"`
	Enclosure    *RssEnclosure
	Guid         string `xml:"guid,omitempty"`    // Id used
	PubDate      string `xml:"pubDate,omitempty"` // created or updated
	Source       string `xml:"source,omitempty"`
	ITitle       string `xml:"itunes:title"`
	IDuration    string `xml:"itunes:duration,omitempty"`
	IImage       string `xml:"itunes:image,omitempty"`
	IExplicit    string `xml:"itunes:explicit,omitempty"`
	IEpisode     string `xml:"itunes:episode,omitempty"`
	ISeason      string `xml:"itunes:season,omitempty"`
	IEpisodeType string `xml:"itunes:episodeType,omitempty"`
	IBlock       string `xml:"itunes:block,omitempty"`
}

type ItunesDescription struct {
	XMLName xml.Name `xml:"description"`
	Content *RssContent
}

type ItunesRss struct {
	*Feed
}

// create a new ItunesRssItem with a generic Item struct's data
func newItunesRssItem(i *Item) *ItunesRssItem {
	item := &ItunesRssItem{
		Title:       i.Title,
		ITitle:      i.Title,
		Link:        i.Link.Href,
		Description: &ItunesDescription{Content: &RssContent{Content: i.Description}},
		Guid:        i.Id,
		PubDate:     anyTimeFormat(time.RFC1123Z, i.Created, i.Updated),
	}

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

// create a new ItunesRssFeed with a generic Feed struct's data
func (r *ItunesRss) ItunesRssFeed() *ItunesRssFeed {
	pub := anyTimeFormat(time.RFC1123Z, r.Created, r.Updated)
	build := anyTimeFormat(time.RFC1123Z, r.Updated)
	author := ""
	ownerName := ""
	ownerEmail := ""
	if r.Author != nil {
		author = r.Author.Email
		ownerEmail = r.Author.Email
		ownerName = r.Author.Name
		if len(r.Author.Name) > 0 {
			author = fmt.Sprintf("%s (%s)", r.Author.Email, r.Author.Name)
		}
	}

	var image *RssImage
	if r.Image != nil {
		image = &RssImage{Url: r.Image.Url, Title: r.Image.Title, Link: r.Image.Link, Width: r.Image.Width, Height: r.Image.Height}
	}

	channel := &ItunesRssFeed{
		Title:         r.Title,
		Link:          r.Link.Href,
		Description:   r.Description,
		PubDate:       pub,
		LastBuildDate: build,
		Copyright:     r.Copyright,
		Image:         image,
		IAuthor:       author,
	}

	if ownerEmail != "" || ownerName != "" {
		owner := &ItunesOwner{Name: ownerName, Email: ownerEmail}
		channel.IOwner = owner
	}

	for _, i := range r.Items {
		channel.Items = append(channel.Items, newItunesRssItem(i))
	}

	return channel
}

// FeedXml returns an XML-Ready object for an ItunesRss object
func (r *ItunesRss) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	return r.ItunesRssFeed().FeedXml()

}

// FeedXml returns an XML-ready object for an ItunesRssFeed object
func (r *ItunesRssFeed) FeedXml() interface{} {
	return &ItunesRssFeedXml{
		Version:          "2.0",
		ContentNamespace: "http://purl.org/rss/1.0/modules/content/",
		Channel:          r,
		ItunesNamespace:  "http://www.itunes.com/dtds/podcast-1.0.dtd",
	}
}
