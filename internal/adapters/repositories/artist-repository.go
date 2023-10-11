package repositories

import (
	"github.com/Hann0/music-api-clone/internal/models"
	"github.com/jmoiron/sqlx"
)

type ArtistPostgres struct {
	DB *sqlx.DB
}

func NewArtistPostgresRepository(DB *sqlx.DB) *ArtistPostgres {
	return &ArtistPostgres{DB}
}

func (r *ArtistPostgres) ReadArtist(id int) (*models.Artist, error) {
	artist := models.Artist{}

	err := r.DB.Get(&artist, "select * from artists where id=$1", id)
	if err != nil {
		return &artist, err
	}

	return &artist, nil
}

func (r *ArtistPostgres) ReadArtists() ([]*models.Artist, error) {
	artists := []*models.Artist{}

	err := r.DB.Select(&artists, "select * from artists")
	if err != nil {
		return artists, err
	}

	return artists, nil
}
