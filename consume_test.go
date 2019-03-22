package feeds

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

var testRssFeedXML = RssFeedXml{
	XMLName:          xml.Name{Space: "", Local: "rss"},
	Version:          "2.0",
	ContentNamespace: "",
	Channel: &RssFeed{
		XMLName:        xml.Name{Space: "", Local: "channel"},
		Title:          "Lorem ipsum feed for an interval of 1 minutes",
		Link:           "http://example.com/",
		Description:    "This is a constantly updating lorem ipsum feed",
		Language:       "",
		Copyright:      "Michael Bertolacci, licensed under a Creative Commons Attribution 3.0 Unported License.",
		ManagingEditor: "",
		WebMaster:      "",
		PubDate:        "Tue, 30 Oct 2018 23:22:00 GMT",
		LastBuildDate:  "Tue, 30 Oct 2018 23:22:37 GMT",
		Category:       "",
		Generator:      "RSS for Node",
		Docs:           "",
		Cloud:          "",
		Ttl:            60,
		Rating:         "",
		SkipHours:      "",
		SkipDays:       "",
		Image:          (*RssImage)(nil),
		TextInput:      (*RssTextInput)(nil),
		Items: []*RssItem{
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:22:00+00:00",
				Link:        "http://example.com/test/1540941720",
				Description: "Exercitation ut Lorem sint proident.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941720",
				PubDate:     "Tue, 30 Oct 2018 23:22:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:21:00+00:00",
				Link:        "http://example.com/test/1540941660",
				Description: "Ea est do quis fugiat exercitation.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941660",
				PubDate:     "Tue, 30 Oct 2018 23:21:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:20:00+00:00",
				Link:        "http://example.com/test/1540941600",
				Description: "Ipsum velit cillum ad laborum sit nulla exercitation consequat sint veniam culpa veniam voluptate incididunt.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941600",
				PubDate:     "Tue, 30 Oct 2018 23:20:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:19:00+00:00",
				Link:        "http://example.com/test/1540941540",
				Description: "Ullamco pariatur aliqua consequat ea veniam id qui incididunt laborum.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941540",
				PubDate:     "Tue, 30 Oct 2018 23:19:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:18:00+00:00",
				Link:        "http://example.com/test/1540941480",
				Description: "Velit proident aliquip aliquip anim mollit voluptate laboris voluptate et occaecat occaecat laboris ea nulla.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941480",
				PubDate:     "Tue, 30 Oct 2018 23:18:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:17:00+00:00",
				Link:        "http://example.com/test/1540941420",
				Description: "Do in quis mollit consequat id in minim laborum sint exercitation laborum elit officia.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941420",
				PubDate:     "Tue, 30 Oct 2018 23:17:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:16:00+00:00",
				Link:        "http://example.com/test/1540941360",
				Description: "Irure id sint ullamco Lorem magna consectetur officia adipisicing duis incididunt.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941360",
				PubDate:     "Tue, 30 Oct 2018 23:16:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:15:00+00:00",
				Link:        "http://example.com/test/1540941300",
				Description: "Sunt anim excepteur esse nisi commodo culpa laborum exercitation ad anim ex elit.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941300",
				PubDate:     "Tue, 30 Oct 2018 23:15:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:14:00+00:00",
				Link:        "http://example.com/test/1540941240",
				Description: "Excepteur aliquip fugiat ex labore nisi.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941240",
				PubDate:     "Tue, 30 Oct 2018 23:14:00 GMT",
				Source:      "",
			},
			{
				XMLName:     xml.Name{Space: "", Local: "item"},
				Title:       "Lorem ipsum 2018-10-30T23:13:00+00:00",
				Link:        "http://example.com/test/1540941180",
				Description: "Id proident adipisicing proident pariatur aute pariatur pariatur dolor dolor in voluptate dolor.",
				Content:     (*RssContent)(nil),
				Author:      "",
				Category:    "",
				Comments:    "",
				Enclosure:   (*RssEnclosure)(nil),
				Guid:        "http://example.com/test/1540941180",
				PubDate:     "Tue, 30 Oct 2018 23:13:00 GMT",
				Source:      "",
			},
		},
	},
}

var testAtomFeedXML = AtomFeed{
	XMLName:  xml.Name{Space: "", Local: "feed"},
	Xmlns:    "",
	Title:    "Lorem ipsum feed for an interval of 1 minutes",
	Id:       "",
	Updated:  "",
	Category: "",
	Icon:     "",
	Logo:     "",
	Rights:   "",
	Subtitle: "",
	Link: &AtomLink{
		XMLName: xml.Name{Space: "", Local: "link"},
		Href:    "",
		Rel:     "",
		Type:    "",
		Length:  "",
	},
	Author: &AtomAuthor{
		XMLName:    xml.Name{Space: "", Local: "author"},
		AtomPerson: AtomPerson{},
	},
	Contributor: (*AtomContributor)(nil),
	Entries: []*AtomEntry{
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:22:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941720"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:21:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941660"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:20:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941600"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:19:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941540"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:18:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941480"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:17:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941420"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:16:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941360"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:15:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941300"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:14:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941240"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
		{
			XMLName:     xml.Name{Space: "", Local: "entry"},
			Xmlns:       "",
			Title:       "Lorem ipsum 2018-10-30T23:13:00+00:00",
			Updated:     "",
			Id:          "",
			Category:    "",
			Content:     (*AtomContent)(nil),
			Rights:      "",
			Source:      "",
			Published:   "",
			Contributor: (*AtomContributor)(nil),
			Links: []AtomLink{{
				XMLName: xml.Name{Space: "", Local: "link"},
				Href:    "http://example.com/test/1540941180"}},
			Summary: (*AtomSummary)(nil),
			Author:  (*AtomAuthor)(nil),
		},
	},
}

func TestRssUnmarshal(t *testing.T) {
	var xmlFeed RssFeedXml
	xmlFile, err := os.Open("test.rss")
	if err != nil {
		panic("AHH file bad")
	}
	bytes, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(bytes, &xmlFeed)

	if !reflect.DeepEqual(testRssFeedXML, xmlFeed) {
		diffs := pretty.Diff(testRssFeedXML, xmlFeed)
		t.Log(pretty.Println(diffs))
		t.Error("object was not unmarshalled correctly")
	}

}

func TestAtomUnmarshal(t *testing.T) {
	var xmlFeed AtomFeed
	xmlFile, err := os.Open("test.atom")
	if err != nil {
		panic("AHH file bad")
	}
	bytes, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(bytes, &xmlFeed)

	if !reflect.DeepEqual(testAtomFeedXML, xmlFeed) {
		diffs := pretty.Diff(testAtomFeedXML, xmlFeed)
		t.Log(pretty.Println(diffs))
		t.Error("object was not unmarshalled correctly")
	}
}
