package payloads

import "github.com/Hann0/music-api-clone/internal/models"

type TrackDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Rank     int    `json:"rank"`
	Type     string `json:"type"`

	Album  *AlbumDTO  `json:"album"`
	Artist *ArtistDTO `json:"artist"`
}

func NewTrackDTO(track *models.Track) *TrackDTO {
	return &TrackDTO{
		ID:       track.ID,
		Title:    track.Title,
		Duration: track.Duration,
		Rank:     track.Rank,
		Type:     "track",
	}
}

func (t *TrackDTO) SetAlbum(album *AlbumDTO) {
	t.Album = album
}

func (t *TrackDTO) SetArtist(artist *ArtistDTO) {
	t.Artist = artist
}
