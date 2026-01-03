package feeds

import (
	"encoding/xml"
	"testing"
)

type LanguageTestStruct struct {
	XMLName  xml.Name   `xml:"test"`
	Language ISO639Code `xml:"language"`
}

func TestISO639CodeMarshalXML(t *testing.T) {
	tests := []struct {
		name    string
		code    ISO639Code
		wantXML string
		wantErr bool
	}{
		{
			name:    "valid code en",
			code:    "en",
			wantXML: "<test><language>en</language></test>",
			wantErr: false,
		},
		{
			name:    "valid code de",
			code:    "de",
			wantXML: "<test><language>de</language></test>",
			wantErr: false,
		},
		{
			name:    "valid code de-AT",
			code:    "de-AT",
			wantXML: "<test><language>de-AT</language></test>",
			wantErr: false,
		},
		{
			name:    "valid code en-us",
			code:    "en-us",
			wantXML: "<test><language>en-us</language></test>",
			wantErr: false,
		},
		{
			name:    "valid code eng",
			code:    "eng",
			wantXML: "<test><language>eng</language></test>",
			wantErr: false,
		},
		{
			name:    "invalid code case",
			code:    "EN",
			wantXML: "",
			wantErr: true,
		},
		{
			name:    "invalid code numeric",
			code:    "12",
			wantXML: "",
			wantErr: true,
		},
		{
			name:    "empty code",
			code:    "",
			wantXML: "<test></test>",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := LanguageTestStruct{Language: tt.code}
			bytes, err := xml.Marshal(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if string(bytes) != tt.wantXML {
					t.Errorf("MarshalXML() got = %s, want %s", string(bytes), tt.wantXML)
				}
			}
		})
	}
}
