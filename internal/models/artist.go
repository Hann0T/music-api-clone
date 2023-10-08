package models

type Artist struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Picture string `db:"picture"`
	Fans    int    `db:"fans"`
}

func NewArtist(id int, name string, picture string, fans int) *Artist {
	return &Artist{id, name, picture, fans}
}
