package repositories

import (
	"errors"

	"github.com/Hann0/music-api-clone/internal/models"
)

type InMemoryAlbumRepository struct {
	Albums []*models.Album
}

func NewInMemoryAlbumRepository() *InMemoryAlbumRepository {
	return &InMemoryAlbumRepository{
		Albums: []*models.Album{
			models.NewAlbum(1, "tyranny", "123", "asdf", 3),
			models.NewAlbum(2, "virtue", "123", "asdf", 3),
		},
	}
}

func (r *InMemoryAlbumRepository) ReadAlbum(id int) (*models.Album, error) {
	for _, album := range r.Albums {
		if album.ID == id {
			return album, nil
		}
	}

	return &models.Album{}, errors.New("Album not found")
}

func (r *InMemoryAlbumRepository) ReadAlbums() ([]*models.Album, error) {
	return r.Albums, nil
}

type InMemoryArtistRepository struct {
	Artists []*models.Artist
}

func NewInMemoryArtistRepository() *InMemoryArtistRepository {
	return &InMemoryArtistRepository{}
}

func (r *InMemoryArtistRepository) ReadArtist(id int) (*models.Artist, error) {
	for _, artist := range r.Artists {
		if artist.ID == id {
			return artist, nil
		}
	}

	return &models.Artist{}, errors.New("Album not found")
}

func (r *InMemoryArtistRepository) ReadArtists() ([]*models.Artist, error) {
	return r.Artists, nil
}
