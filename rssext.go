package feeds

// rssext support
// validation done according to spec here:
//    http://cyber.law.harvard.edu/rss/rss.html

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"time"
)

// private wrapper around the RssFeedExt which gives us the <rss>..</rss> xml
type RssFeedExtXml struct {
	XMLName          xml.Name   `xml:"rss"`
	Version          string     `xml:"version,attr"`
	ContentNamespace string     `xml:"xmlns:content,attr"`
	CustomNamespaces []xml.Attr `xml:",attr"`
	Channel          *RssFeedExt
	Extensions       []Extension `xml:",any"`
}

type RssImageExt struct {
	XMLName    xml.Name    `xml:"image"`
	Url        string      `xml:"url"`
	Title      string      `xml:"title"`
	Link       string      `xml:"link"`
	Width      int         `xml:"width,omitempty"`
	Height     int         `xml:"height,omitempty"`
	Extensions []Extension `xml:",any"`
}

type RssTextInputExt struct {
	XMLName     xml.Name    `xml:"textInput"`
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	Name        string      `xml:"name"`
	Link        string      `xml:"link"`
	Extensions  []Extension `xml:",any"`
}

type RssFeedExt struct {
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
	Image          *RssImageExt
	TextInput      *RssTextInputExt
	Items          []*RssItemExt `xml:"item"`
	Extensions     []Extension   `xml:",any"`
}

type RssItemExt struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`       // required
	Link        string   `xml:"link"`        // required
	Description string   `xml:"description"` // required
	Content     *RssContent
	Author      string `xml:"author,omitempty"`
	Category    string `xml:"category,omitempty"`
	Comments    string `xml:"comments,omitempty"`
	Enclosure   *RssEnclosure
	Guid        *RssGuid    // Id used
	PubDate     string      `xml:"pubDate,omitempty"` // created or updated
	Source      string      `xml:"source,omitempty"`
	Extensions  []Extension `xml:",any"`
}

type RssExt struct {
	*Feed
}

// create a new RssItemExt with a generic Item struct's data
func newRssItemExt(i *Item) *RssItemExt {
	item := &RssItemExt{
		Title:       i.Title,
		Description: i.Description,
		PubDate:     anyTimeFormat(time.RFC1123Z, i.Created, i.Updated),
	}
	if i.Id != "" {
		item.Guid = &RssGuid{Id: i.Id, IsPermaLink: i.IsPermaLink}
	}
	if i.Link != nil {
		item.Link = i.Link.Href
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

// create a new RssFeedExt with a generic Feed struct's data
func (r *RssExt) RssFeedExt() *RssFeedExt {
	pub := anyTimeFormat(time.RFC1123Z, r.Created, r.Updated)
	build := anyTimeFormat(time.RFC1123Z, r.Updated)
	author := ""
	if r.Author != nil {
		author = r.Author.Email
		if len(r.Author.Name) > 0 {
			author = fmt.Sprintf("%s (%s)", r.Author.Email, r.Author.Name)
		}
	}

	var image *RssImageExt
	if r.Image != nil {
		image = &RssImageExt{Url: r.Image.Url, Title: r.Image.Title, Link: r.Image.Link, Width: r.Image.Width, Height: r.Image.Height}
	}

	var href string
	if r.Link != nil {
		href = r.Link.Href
	}
	channel := &RssFeedExt{
		Title:          r.Title,
		Link:           href,
		Description:    r.Description,
		ManagingEditor: author,
		PubDate:        pub,
		LastBuildDate:  build,
		Copyright:      r.Copyright,
		Image:          image,
	}
	for _, i := range r.Items {
		channel.Items = append(channel.Items, newRssItemExt(i))
	}
	return channel
}

// FeedXml returns an XML-Ready object for an RssExt object
func (r *RssExt) FeedXml() interface{} {
	// only generate version 2.0 feeds for now
	return r.RssFeedExt().FeedXml()

}

// FeedXml returns an XML-ready object for an RssFeedExt object
func (r *RssFeedExt) FeedXml() interface{} {
	return &RssFeedExtXml{
		Version:          "2.0",
		Channel:          r,
		ContentNamespace: "http://purl.org/rss/1.0/modules/content/",
	}
}

// UnmarshalRssFeedExt parses the XML data into an RssFeedExtXml, preserving xmlns attributes.
func UnmarshalRssFeedExt(data []byte) (*RssFeedExtXml, error) {
	feed := &RssFeedExtXml{}

	// Standard unmarshal for content
	if err := xml.Unmarshal(data, feed); err != nil {
		return nil, err
	}

	// xml.Unmarshal strips xmlns attributes, so we parse them manually using RawToken
	dec := xml.NewDecoder(bytes.NewReader(data))
	for {
		t, err := dec.RawToken()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if start, ok := t.(xml.StartElement); ok {
			if start.Name.Local == "rss" {
				for _, attr := range start.Attr {
					if attr.Name.Space == "xmlns" || (attr.Name.Space == "" && attr.Name.Local == "xmlns") {
						if attr.Name.Space == "xmlns" && attr.Name.Local == "content" {
							feed.ContentNamespace = attr.Value
						} else {
							// To ensure round-trip compatibility with xml.Marshal (which strips unused namespaces),
							// we treat these as regular attributes by including the prefix in Local.
							if attr.Name.Space == "xmlns" {
								attr.Name.Local = "xmlns:" + attr.Name.Local
								attr.Name.Space = ""
							}
							feed.CustomNamespaces = append(feed.CustomNamespaces, attr)
						}
					}
				}
				break
			}
		}
	}
	return feed, nil
}

type Extension struct {
	XMLName  xml.Name
	Attrs    []xml.Attr  `xml:",any,attr"`
	Children []Extension `xml:",any"`
	Value    string      `xml:",chardata"`
}

func (e *Extension) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Filter xmlns attributes
	var filteredAttrs []xml.Attr
	for _, attr := range start.Attr {
		if attr.Name.Space == "xmlns" || attr.Name.Local == "xmlns" {
			continue
		}
		filteredAttrs = append(filteredAttrs, attr)
	}
	// Update start.Attr so DecodeElement uses the filtered list
	start.Attr = filteredAttrs

	type extensionAlias Extension
	var v extensionAlias

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	*e = Extension(v)
	e.Value = strings.TrimSpace(e.Value)
	return nil
}
