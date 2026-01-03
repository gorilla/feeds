package feeds

import (
	"encoding/xml"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

var testRssFeedExtXML = RssFeedExtXml{
	XMLName:          xml.Name{Space: "", Local: "rss"},
	Version:          "2.0",
	ContentNamespace: "http://purl.org/rss/1.0/modules/content/",
	CustomNamespaces: []xml.Attr{
		{Name: xml.Name{Local: "xmlns:atom"}, Value: "http://www.w3.org/2005/Atom"},
		{Name: xml.Name{Local: "xmlns:custom"}, Value: "http://www.example.com/custom"},
		{Name: xml.Name{Local: "xmlns:dc"}, Value: "http://purl.org/dc/elements/1.1/"},
	},
	Channel: &RssFeedExt{
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
		Image:          (*RssImageExt)(nil),
		TextInput:      (*RssTextInputExt)(nil),
		Items: []*RssItemExt{
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
				Guid:        &RssGuid{XMLName: xml.Name{Local: "guid"}, Id: "http://example.com/test/1540941720", IsPermaLink: "true"},
				PubDate:     "Tue, 30 Oct 2018 23:22:00 GMT",
				Source:      "",
				Extensions: []Extension{
					{
						XMLName: xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "creator"},
						Value:   "John Smith",
					},
					{
						XMLName: xml.Name{Space: "http://www.example.com/custom", Local: "group"},
						Attrs:   []xml.Attr{{Name: xml.Name{Local: "name"}, Value: "item name"}},
						Children: []Extension{
							{
								XMLName: xml.Name{Space: "http://www.example.com/custom", Local: "elem"},
								Value:   "item content",
							},
						},
					},
				},
			},
		},
		Extensions: []Extension{
			{
				XMLName: xml.Name{Local: "author"},
				Value:   "John Smith",
			},
			{
				XMLName: xml.Name{Space: "http://www.example.com/custom", Local: "group"},
				Attrs:   []xml.Attr{{Name: xml.Name{Local: "name"}, Value: "channel name"}},
				Children: []Extension{
					{
						XMLName: xml.Name{Space: "http://www.example.com/custom", Local: "elem"},
						Value:   "channel content",
					},
				},
			},
		},
	},
}

func TestRssExtUnmarshal(t *testing.T) {
	var xmlFeed RssFeedExtXml
	xmlFile, err := os.Open("testext.rss")
	if err != nil {
		panic("AHH file bad")
	}
	bytes, _ := io.ReadAll(xmlFile)
	feedPtr, err := UnmarshalRssFeedExt(bytes)
	if err != nil {
		panic(err)
	}
	xmlFeed = *feedPtr

	if !reflect.DeepEqual(testRssFeedExtXML, xmlFeed) {
		diffs := pretty.Diff(testRssFeedExtXML, xmlFeed)
		t.Log(pretty.Println(diffs))
		t.Error("object was not unmarshalled correctly")
	}
}

func TestRssExtRoundTrip(t *testing.T) {
	// Marshal the test object
	bytes, err := xml.MarshalIndent(testRssFeedExtXML, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Unmarshal it back
	roundTripFeed, err := UnmarshalRssFeedExt(bytes)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Compare
	if !reflect.DeepEqual(testRssFeedExtXML, *roundTripFeed) {
		diffs := pretty.Diff(testRssFeedExtXML, *roundTripFeed)
		t.Log(pretty.Println(diffs))
		t.Error("Roundtrip failed: objects are different")
	}
}

func TestRssExtCustomExtensions(t *testing.T) {
	ns := "http://www.example.com/custom"

	findExt := func(exts []Extension, space, local string) *Extension {
		for i := range exts {
			if exts[i].XMLName.Space == space && exts[i].XMLName.Local == local {
				return &exts[i]
			}
		}
		return nil
	}

	// Check Channel
	group := findExt(testRssFeedExtXML.Channel.Extensions, ns, "group")
	if group == nil {
		t.Fatal("Channel custom:group not found")
	}

	elem := findExt(group.Children, ns, "elem")
	if elem == nil {
		t.Fatal("Channel custom:elem not found inside group")
	}

	if elem.Value != "channel content" {
		t.Errorf("Channel custom:elem value mismatch. Got %q, want %q", elem.Value, "channel content")
	}

	// Check Item
	if len(testRssFeedExtXML.Channel.Items) == 0 {
		t.Fatal("No items in channel")
	}
	item := testRssFeedExtXML.Channel.Items[0]

	group = findExt(item.Extensions, ns, "group")
	if group == nil {
		t.Fatal("Item custom:group not found")
	}

	// Check attribute
	foundAttr := false
	for _, attr := range group.Attrs {
		if attr.Name.Local == "name" && attr.Value == "item name" {
			foundAttr = true
			break
		}
	}
	if !foundAttr {
		t.Error("Item custom:group attribute 'name' with value 'item name' not found")
	}

	elem = findExt(group.Children, ns, "elem")
	if elem == nil {
		t.Fatal("Item custom:elem not found inside group")
	}

	if elem.Value != "item content" {
		t.Errorf("Item custom:elem value mismatch. Got %q, want %q", elem.Value, "item content")
	}
}
