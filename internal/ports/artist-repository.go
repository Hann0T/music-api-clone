package ports

import "github.com/Hann0/music-api-clone/internal/models"

type ArtistRepository interface {
	ReadArtist(id int) (*models.Artist, error)
	ReadArtists() ([]*models.Artist, error)
}
