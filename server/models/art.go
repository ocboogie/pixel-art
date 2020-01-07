package models

// ArtFormat is the info needed to create a correctly formatted art for a post.
// Currently this is the same struct as the config but later on this might
// change.
type ArtFormat struct {
	Size   uint16 `json:"size"`
	Colors uint8  `json:"colors"`
}
