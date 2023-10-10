package repositories

import (
	"github.com/Hann0/music-api-clone/internal/models"
	"github.com/jmoiron/sqlx"
)

type AlbumPostgres struct {
	DB *sqlx.DB
}

func NewAlbumPostgresRepository(DB *sqlx.DB) *AlbumPostgres {
	return &AlbumPostgres{DB}
}

func (r *AlbumPostgres) ReadAlbum(id int) (*models.Album, error) {
	album := models.Album{}
	err := r.DB.Get(&album, "select * from albums where id=$1", id)
	if err != nil {
		return &album, err
	}
	return &album, nil
}

func (r *AlbumPostgres) ReadAlbumsByArtist(artist *models.Artist) ([]*models.Album, error) {
	albums := []*models.Album{}

	err := r.DB.Select(&albums, "select * from albums where artist_id=$1", artist.ID)

	if err != nil {
		return albums, err
	}

	return albums, nil
}

func (r *AlbumPostgres) ReadAlbums() ([]*models.Album, error) {
	albums := []*models.Album{}

	err := r.DB.Select(&albums, "select * from albums")
	if err != nil {
		return albums, err
	}

	return albums, nil
}
