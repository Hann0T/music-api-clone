package models

type Track struct {
	ID       int
	Title    string
	Duration int
	Rank     int
	AlbumID  int
}

func NewTrack(id int, title string, duration int, rank int, albumID int) *Track {
	return &Track{id, title, duration, rank, albumID}
}
