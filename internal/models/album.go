package models

type Album struct {
	ID       int
	Title    string
	UPC      string
	Cover    string
	ArtistID int
}

func NewAlbum(id int, title string, upc string, cover string, artistID int) *Album {
	return &Album{id, title, upc, cover, artistID}
}
