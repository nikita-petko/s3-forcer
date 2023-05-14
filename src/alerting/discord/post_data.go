package discord

type embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"` // It comes from a hex-encoded number
}

type postData struct {
	Embeds []*embed `json:"embed"`
}
