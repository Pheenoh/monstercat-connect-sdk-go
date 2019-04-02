package monstercat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Response struct {
	Results []Track `json:"results"`
	Sort    []sort  `json:"sort"`
}

type sort struct {
	ReleaseDate int64 `json:"releaseDate"`
}

type Track struct {
	Artists              string        `json:"renderedArtists"`
	CatalogID            string        `json:"catalogId"`
	CoverArt             string        `json:"coverArt"`
	CoverURL             string        `json:"coverurl"`
	Downloadable         bool          `json:"downloadable"`
	FreeDownloadForUsers bool          `json:"freeDownloadForUsers"`
	GRID                 string        `json:"grid"`
	ID                   string        `json:"_id"`
	ImageHashSum         string        `json:"imageHashSum"`
	InEarlyAccess        bool          `json:"inEarlyAccess"`
	Label                string        `json:"label"`
	PlaylistID           string        `json:"label"`
	PreReleaeDate        time.Time     `json:"preReleaseDate"`
	PrimaryGenre         string        `json:"primaryGenre"`
	ReleaseDate          time.Time     `json:"releaseDate"`
	SecondaryGenre       string        `json:"secondaryGenre"`
	ShowOnWebsite        bool          `json:"showOnWebsite"`
	Streamable           bool          `json:"streamable"`
	Tags                 []string      `json:"tags"`
	ThumbHashes          ThumbnailHash `json:"thumbHashes"`
	Title                string        `json:"title"`
	Type                 string        `json:"tyoe"`
	UPC                  string        `json:"upc"`
	URLs                 []PlatformURL `json:"urls"`
}

type PlatformURL struct {
	Original string `json:"original"`
	Short    string `json:"short"`
	Platform string `json:"platform"`
}

type ThumbnailHash struct {
	Hash32   string `json:"32"`
	Hash64   string `json:"64"`
	Hash128  string `json:"128"`
	Hash256  string `json:"256"`
	Hash512  string `json:"512"`
	Hash1024 string `json:"1024"`
}

func (c *Client) GetTracks() ([]Track, error) {
	path := fmt.Sprintf("/api/catalog/track")
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

	// finish
}
