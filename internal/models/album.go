package models

type Album struct {
	Title    string `json:"title"`
	UPC      string `json:"upc"`
	Cover    string `json:"cover"`
	ArtistID int    `json:"artist_id"`
}

func NewAlbum(title string, upc string, cover string, artistID int) (Album, error) {
	return Album{title, upc, cover, artistID}, nil
}
