package repositories

import (
	"github.com/Hann0/music-api-clone/internal/models"
	"github.com/jmoiron/sqlx"
)

type TrackPostgres struct {
	DB *sqlx.DB
}

func NewTrackPostgresRepository(DB *sqlx.DB) *TrackPostgres {
	return &TrackPostgres{DB}
}

func (r *TrackPostgres) ReadTrack(id int) (*models.Track, error) {
	track := models.Track{}

	err := r.DB.Get(&track, "select * from tracks where id=$1", id)

	if err != nil {
		return &track, err
	}

	return &track, nil
}

func (r *TrackPostgres) ReadTracks() ([]*models.Track, error) {
	tracks := []*models.Track{}

	err := r.DB.Select(&tracks, "select * from tracks")

	if err != nil {
		return tracks, err
	}

	return tracks, nil
}

func (r *TrackPostgres) ReadTracksByAlbum(album *models.Album) ([]*models.Track, error) {
	tracks := []*models.Track{}

	err := r.DB.Select(&tracks, "select * from tracks where album_id=$1", album.ID)

	if err != nil {
        return tracks, err
    }

	return tracks, nil
}
