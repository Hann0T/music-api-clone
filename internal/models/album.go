package models

type Album struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	UPC      string `db:"upc"`
	Cover    string `db:"cover"`
	ArtistID int    `db:"artist_id"`
}

func NewAlbum(id int, title string, upc string, cover string, artistID int) *Album {
	return &Album{id, title, upc, cover, artistID}
}
