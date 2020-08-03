package monstercat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type Response struct {
	Related []Release `json:"related"`
	Release Release   `json:"release"`
	Tracks  []Track   `json:"tracks"`
}

type Release struct {
	ArtistsTitle   string    `json:"artistsTitle"`
	CatalogID      string    `json:"catalogId"`
	Downloadable   bool      `json:"downloadable"`
	GenrePrimary   string    `json:"genrePrimary"`
	GenreSecondary string    `json:"genreSecondary"`
	ID             string    `json:"id"`
	InEarlyAccess  bool      `json:"inEarlyAccess"`
	Links          []string  `json:"links"`
	ReleaseDate    time.Time `json:"releaseDate"`
	Streamable     bool      `json:"streamable"`
	Title          string    `json:"title"`
	Type           string    `json:"type"`
	Version        string    `json:"version"`
}

type Track struct {
	Artists         []Artist  `json:"artists"`
	ArtistsTitle    string    `json:"artistsTitle"`
	BPM             int64     `json:"bpm"`
	CreatorFriendly bool      `json:""`
	DebutDate       time.Time `json:"debutDate"`
	Downloadable    bool      `json:"downloadable"`
	Duration        int64     `json:"duration"`
	Explicit        bool      `json:"explicit"`
	GenrePrimary    string    `json:"genrePrimary"`
	GenreSecondary  string    `json:"genreSecondary"`
	ID              string    `json:"id"`
	InEarlyAccess   bool      `json:"inEarlyAccess"`
	ISRC            bool      `json:"isrc"`
	Release         Release   `json:"release"`
	Streamable      bool      `json:"streamable"`
	Tags            []string  `json:"tags"`
	Title           string    `json:"title"`
	TrackNumber     int64     `json:"tracknumber"`
	Version         string    `json:"version"`
}

func (c *Client) GetTracks(CatalogID string) ([]Track, error) {
	path := fmt.Sprintf("/v2/catalog/release/%s", CatalogID)
	req, err := c.NewRequest("GET", path)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}
	return response.Tracks, err
}
