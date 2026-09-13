package feeds

import (
	"encoding/json"
	"testing"
	"time"
)

func TestJSONAttachmentDurationValue(t *testing.T) {
	attachment := JSONAttachment{Duration: 1500 * time.Millisecond}
	for _, value := range []interface{}{attachment, &attachment, map[string]JSONAttachment{"audio": attachment}} {
		data, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		var wire map[string]interface{}
		if err := json.Unmarshal(data, &wire); err != nil {
			t.Fatal(err)
		}
		if nested, ok := wire["audio"]; ok {
			wire = nested.(map[string]interface{})
		}
		if got := wire["duration_in_seconds"]; got != 1.5 {
			t.Errorf("duration = %v, want 1.5 (%T)", got, value)
		}
	}
}
