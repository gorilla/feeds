package feeds

// https://help.apple.com/itc/podcasts_connect/#/itcb54353390

type ITunesFeed struct {
	Category []*ITunesCategory `xml:"itunes:category,omitempty"` // The show category information.
	Explicit bool              `xml:"itunes:explicit"`           // The episode parental advisory information.
	Type     ITunesFeedType    `xml:"itunes:type,omitempty"`     // The type of show.
	Title    string            `xml:"itunes:title,omitempty"`    // The show title specific for Apple Podcasts.
	Image    *ITunesImage      `xml:"itunes:image,omitempty"`    // The episode artwork.
}

type ITunesImage struct {
	Href string `xml:"href,attr,omitempty"` // URL to image.
}

type ITunesCategory struct {
	Category string `xml:"text,attr,omitempty"` // Category name.
}

type ITunesItem struct {
	Duration    string            `xml:"itunes:duration,omitempty"`    // The duration of an episode.
	Author      string            `xml:"itunes:author,omitempty"`      // The group responsible for creating the show.
	EpisodeType ITunesEpisodeType `xml:"itunes:episodeType,omitempty"` // The episode type.
}

type ITunesEpisodeType string

const (
	ITunesEpisodeTypeFull    ITunesEpisodeType = "Full"
	ITunesEpisodeTypeTrailer ITunesEpisodeType = "Trailer"
	ITunesEpisodeTypeBonus   ITunesEpisodeType = "Bonus"
)

type ITunesFeedType string

const (
	ITunesFeedTypeEpisodic ITunesFeedType = "Episodic"
	ITunesFeedTypeSerial   ITunesFeedType = "Serial"
)
