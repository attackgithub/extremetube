package extremetube

type ExtremetubeEmbedCode struct {
	Embed ExtremetubeCode `json:"embed,omitempty"`
}

type ExtremetubeCode struct {
	Code string `json:"code,omitempty"`
}
