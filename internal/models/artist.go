package models

type Artist struct {
	ID      int
	Name    string
	Picture string
	Fans    int
}

func NewArtist(id int, name string, picture string, fans int) *Artist {
	return &Artist{id, name, picture, fans}
}
