package feeds

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestJSONFeedEmptyItems(t *testing.T) {
	for _, items := range [][]*JSONItem{nil, {}} {
		feed := JSONFeed{Version: jsonFeedVersion, Title: "Empty feed", Items: items}
		data, err := json.Marshal(feed)
		if err != nil {
			t.Fatal(err)
		}
		var got map[string]json.RawMessage
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatal(err)
		}
		if string(got["items"]) != "[]" {
			t.Errorf("items = %s, want []", got["items"])
		}
	}
}

func TestFeedEmptyJSONItems(t *testing.T) {
	feed := &Feed{Title: "Empty feed"}
	text, err := feed.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if err := feed.WriteJSON(&buf); err != nil {
		t.Fatal(err)
	}
	for _, data := range []string{text, buf.String()} {
		var got map[string]json.RawMessage
		if err := json.Unmarshal([]byte(data), &got); err != nil {
			t.Fatal(err)
		}
		if string(got["items"]) != "[]" {
			t.Errorf("items = %s, want []", got["items"])
		}
	}
	if feed.Items != nil {
		t.Fatal("serialization mutated source items")
	}
}
