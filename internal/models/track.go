package models

type Track struct {
	ID       int    `db:"id"`
	Title    string `db:"title"`
	Duration int    `db:"duration"`
	Rank     int    `db:"rank"`
	AlbumID  int    `db:"album_id"`
}

func NewTrack(id int, title string, duration int, rank int, albumID int) *Track {
	return &Track{id, title, duration, rank, albumID}
}
