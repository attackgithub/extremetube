package extremetube

type ExtremetubeVideo struct {
	ID           float64 `json:"id,omitempty"`
	Duration     string  `json:"duration,omitempty"`
	Views        float64 `json:"views,omitempty"`
	Rating       string  `json:"rating,omitempty"`
	Ratings      string  `json:"ratings,omitempty"`
	Title        string  `json:"title,omitempty"`
	URL          string  `json:"url,omitempty"`
	Thumb        string  `json:"thumb,omitempty"`
	DefaultThumb string  `json:"default_thumb,omitempty"`
	PublishDate  string  `json:"publish_date,omitempty"`
}
