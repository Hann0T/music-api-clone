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

func (r *InMemoryAlbumRepository) ReadAlbumsByArtist(artist *models.Artist) ([]*models.Album, error) {
	var albums []*models.Album
	for _, album := range r.Albums {
		if album.ArtistID == artist.ID {
			albums = append(albums, album)
		}
	}

	return albums, nil
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

type InMemoryTrackRepository struct {
	Tracks []*models.Track
}

func NewInMemoryTrackRepository() *InMemoryTrackRepository {
	return &InMemoryTrackRepository{
		Tracks: []*models.Track{
			models.NewTrack(1, "Take Me In Your Army", 254, 12, 1),
			models.NewTrack(2, "Human sadness", 1054, 5, 1),
			models.NewTrack(3, "QYURRYUS", 314, 2, 2),
			models.NewTrack(4, "Blind", 223, 1, 3),
			models.NewTrack(5, "Twist", 131, 1, 4),
			models.NewTrack(6, "Freak on a leash", 236, 4, 5),
		},
	}
}

func (r *InMemoryTrackRepository) ReadTrack(id int) (*models.Track, error) {
	for _, track := range r.Tracks {
		if track.AlbumID == id {
			return track, nil
		}
	}

	return &models.Track{}, errors.New("Track not found")
}

func (r *InMemoryTrackRepository) ReadTracks() ([]*models.Track, error) {
	return r.Tracks, nil
}

func (r *InMemoryTrackRepository) ReadTracksByAlbum(album *models.Album) ([]*models.Track, error) {
	var tracks []*models.Track
	for _, track := range r.Tracks {
		if track.AlbumID == album.ID {
			tracks = append(tracks, track)
		}
	}

	return tracks, nil
}
