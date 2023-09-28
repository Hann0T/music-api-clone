package models

type Album struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	UPC      string `json:"upc"`
	Cover    string `json:"cover"`
	ArtistID int    `json:"artist_id"`
}

func NewAlbum(id int, title string, upc string, cover string, artistID int) *Album {
	return &Album{id, title, upc, cover, artistID}
}
