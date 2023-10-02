package ports

import "github.com/Hann0/music-api-clone/internal/models"

type TrackRepository interface {
	ReadTrack(id int) (*models.Track, error)
	ReadTracksByAlbum(album *models.Album) ([]*models.Track, error)
	ReadTracks() ([]*models.Track, error)
}
