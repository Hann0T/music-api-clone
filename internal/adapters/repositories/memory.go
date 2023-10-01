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
			models.NewAlbum(1, "Tyranny", "123", "img", 1),
			models.NewAlbum(2, "Virtue", "124", "img", 1),
			models.NewAlbum(3, "Korn", "125", "img", 2),
			models.NewAlbum(4, "Life is peachy", "126", "img", 2),
			models.NewAlbum(5, "Follow the leader", "127", "img", 2),
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
	return &InMemoryArtistRepository{
        Artists: []*models.Artist{
            models.NewArtist(1, "The voidz", "img", 10000),
            models.NewArtist(2, "Korn", "img", 120409),
        },
    }
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
