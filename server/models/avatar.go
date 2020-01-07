package models

// AvatarFormat is the info needed to create a correctly formatted avatar. Currently
// this is the same struct as the config but later on this might change.
type AvatarFormat struct {
	Size    uint8    `json:"size"`
	Palette []string `json:"palette"`
}
