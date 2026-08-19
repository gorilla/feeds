package feeds

import (
	"regexp"
	"testing"
)

var uuidV4Regex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

func TestUUID(t *testing.T) {
	for i := 0; i < 50; i++ {
		s := NewUUID()
		s2 := NewUUID()
		if len(s) != 16 {
			t.Errorf("Expecting len of 16, got %d\n", len(s))
		}
		if len(s.String()) != 36 {
			t.Errorf("Expecting uuid hex string len of 36, got %d\n", len(s.String()))
		}
		if *s == *s2 {
			t.Errorf("Expecting different UUID values to be different, but they are the same.\n")
		}
		// Version 4 check (0x40 in high nibble of byte 6)
		if (s[6] & 0xf0) != 0x40 {
			t.Errorf("Expecting version 4 (0x40), got %#x", s[6]&0xf0)
		}
		// Variant check (0x80 in high 2 bits of byte 8)
		if (s[8] & 0xc0) != 0x80 {
			t.Errorf("Expecting RFC 4122 variant (0x80), got %#x", s[8]&0xc0)
		}
		// Hex string format check
		if !uuidV4Regex.MatchString(s.String()) {
			t.Errorf("UUID string %q does not match RFC 4122 v4 pattern", s.String())
		}
	}
}
