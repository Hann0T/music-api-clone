package payloads

import "github.com/Hann0/music-api-clone/internal/models"

type ArtistDTO struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Picture      string `json:"picture"`
	Fans         int    `json:"fans"`
	Type         string `json:"type"`
	NumberAlbums int    `json:"nb_albums"`
}

func NewArtistDTO(artist *models.Artist) *ArtistDTO {
	return &ArtistDTO{
		ID:      artist.ID,
		Name:    artist.Name,
		Picture: artist.Picture,
		Fans:    artist.Fans,
		Type:    "artist",
	}
}

func (a *ArtistDTO) SetNumberAlbums(number int) {
	a.NumberAlbums = number
}
