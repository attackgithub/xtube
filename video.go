package xtube

type XtubeVideo struct {
	ID           string  `json:"video_id,omitempty"`
	Duration     string  `json:"duration,omitempty"`
	Views        float64 `json:"views,omitempty"`
	Rating       string  `json:"rating,omitempty"`
	Ratings      string  `json:"ratings,omitempty"`
	Title        string  `json:"title,omitempty"`
	URL          string  `json:"url,omitempty"`
	EmbedCode    string  `json:"embedCode,omitempty"`
	Description  string  `json:"description,omitempty"`
	DefaultThumb string  `json:"default_thumb,omitempty"`
	Thumb        string  `json:"thumb,omitempty"`
	PublishDate  string  `json:"publish_date,omitempty"`
	Thumbs       []XtubeThumb
}

type XtubeThumb struct {
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
	Src    string  `json:"src,omitempty"`
}

type XtubeTag struct {
	Name string `json:"tag_name,omitempty"`
}
