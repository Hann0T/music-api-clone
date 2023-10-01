package models

type Track struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Rank     int    `json:"rank"`
	AlbumID  int    `json:"album_id"`
}

func NewTrack(id int, title string, duration int, rank int, albumID int) *Track {
	return &Track{id, title, duration, rank, albumID}
}
