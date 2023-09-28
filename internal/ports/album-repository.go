package ports

import "github.com/Hann0/music-api-clone/internal/models"

type AlbumRepository interface {
	ReadAlbum(id int) (*models.Album, error)
	ReadAlbums() ([]*models.Album, error)
}
