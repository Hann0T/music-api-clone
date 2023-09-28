package models

type Artist struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Fans    int    `json:"fans"`
}

func NewArtist(id int, name string, picture string, fans int) *Artist {
	return &Artist{id, name, picture, fans}
}
