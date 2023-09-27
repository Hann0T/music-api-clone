package models

type Artist struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Fans    int    `json:"fans"`
}

func NewArtist(name string, picture string, fans int) (Artist, error) {
	return Artist{name, picture, fans}, nil
}
