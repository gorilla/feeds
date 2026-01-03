package feeds

import (
	"encoding/xml"
	"errors"
	"regexp"
)

var iso639Regex = regexp.MustCompile(`^[a-z]{2,3}(-[a-zA-Z0-9-]+)?$`)

type ISO639Code string

func (c ISO639Code) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if c == "" {
		return nil
	}
	if !iso639Regex.MatchString(string(c)) {
		return errors.New("invalid ISO 639 code")
	}
	return e.EncodeElement(string(c), start)
}
