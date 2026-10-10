package feeds

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJSONAttachmentUnmarshalDuration(t *testing.T) {
	for _, tt := range []struct {
		name string
		data string
		want time.Duration
	}{
		{"zero", `{"duration_in_seconds":0}`, 0},
		{"fractional", `{"duration_in_seconds":1.5}`, 1500 * time.Millisecond},
		{"omitted", `{"title":"new title"}`, 5 * time.Second},
		{"null", `{"duration_in_seconds":null}`, 5 * time.Second},
		{"negative", `{"duration_in_seconds":-1}`, 5 * time.Second},
	} {
		t.Run(tt.name, func(t *testing.T) {
			a := JSONAttachment{Url: "https://example.com/audio", Duration: 5 * time.Second}
			if err := json.Unmarshal([]byte(tt.data), &a); err != nil {
				t.Fatal(err)
			}
			if a.Duration != tt.want {
				t.Errorf("Duration = %s, want %s", a.Duration, tt.want)
			}
			if a.Url != "https://example.com/audio" {
				t.Errorf("omitted URL was changed: %q", a.Url)
			}
		})
	}
}
