package feeds

import (
	"encoding/json"
	"testing"
)

func TestJSONBannerImage(t *testing.T) {
	const wire = `{"id":"1","banner_image":"https://example.com/banner.png"}`
	var item JSONItem
	if err := json.Unmarshal([]byte(wire), &item); err != nil {
		t.Fatal(err)
	}
	if item.BannerImage != "https://example.com/banner.png" {
		t.Fatalf("banner image = %q", item.BannerImage)
	}
	got, err := json.Marshal(item)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != wire {
		t.Fatalf("encoded item = %s, want %s", got, wire)
	}
}

func TestJSONAttachmentSize(t *testing.T) {
	const wire = `{"url":"https://example.com/audio.mp3","mime_type":"audio/mpeg","size_in_bytes":12345}`
	var attachment JSONAttachment
	if err := json.Unmarshal([]byte(wire), &attachment); err != nil {
		t.Fatal(err)
	}
	if attachment.Size != 12345 {
		t.Fatalf("attachment size = %d", attachment.Size)
	}
	got, err := json.Marshal(&attachment)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != wire {
		t.Fatalf("encoded attachment = %s, want %s", got, wire)
	}
}
