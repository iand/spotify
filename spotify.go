/*
  This is free and unencumbered software released into the public domain. For more
  information, see <http://unlicense.org/> or the accompanying UNLICENSE file.
*/

// Client for the Spotify API
package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
}

type SearchTracksResponse struct {
	Info   SearchInfo `json:"info"`
	Tracks []Track    `json:"tracks"`
}

type SearchAlbumsResponse struct {
	Info   SearchInfo `json:"info"`
	Albums []Album    `json:"albums"`
}

type SearchArtistsResponse struct {
	Info    SearchInfo `json:"info"`
	Artists []Artist   `json:"artists"`
}

type SearchInfo struct {
	TotalResults int    `json:"num_results"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	Query        string `json:"query"`
	Type         string `json:"type"`
	Page         int    `json:"page"`
}

type Track struct {
	Name        string       `json:"name"`
	Popularity  string       `json:"popularity,omitempty"`
	Length      float64      `json:"length,omitempty"`
	URI         string       `json:"href,omitempty"`
	TrackNumber string       `json:"track-number,omitempty"`
	DiscNumber  string       `json:"disc-number,omitempty"`
	Album       Album        `json:"album"`
	Artists     []Artist     `json:"artists"`
	Available   bool         `json:"available"`
	ExternalIDs []ExternalID `json:"external-ids"`
}

type Album struct {
	Name         string       `json:"name"`
	Released     string       `json:"released,omitempty"`
	Length       float64      `json:"length,omitempty"`
	URI          string       `json:"href,omitempty"`
	Availability Availability `json:"availability,omitempty"`
	Artists      []Artist     `json:"artists"`
	Artist       string       `json:"artist"`
	ExternalIDs  []ExternalID `json:"external-ids"`
}

type Availability struct {
	Territories string `json:"territories"`
}

type ExternalID struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Artist struct {
	Name       string `json:"name"`
	Popularity string `json:"popularity,omitempty"`
	URI        string `json:"href,omitempty"`
}

func New() *Client {
	return &Client{}
}

func (client *Client) SearchTracks(srch string, page int) (*SearchTracksResponse, error) {
	url := fmt.Sprintf("http://ws.spotify.com/search/1/track.json?q=%s&page=%d", url.QueryEscape(srch), page)
	var data SearchTracksResponse

	err := client.doSearch(url, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *Client) SearchAlbums(srch string, page int) (*SearchAlbumsResponse, error) {
	url := fmt.Sprintf("http://ws.spotify.com/search/1/album.json?q=%s&page=%d", url.QueryEscape(srch), page)
	var data SearchAlbumsResponse

	err := client.doSearch(url, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *Client) SearchArtists(srch string, page int) (*SearchArtistsResponse, error) {
	url := fmt.Sprintf("http://ws.spotify.com/search/1/artist.json?q=%s&page=%d", url.QueryEscape(srch), page)
	var data SearchArtistsResponse

	err := client.doSearch(url, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (client *Client) doSearch(url string, data interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("Search failed with http error: %s", err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&data); err != nil {
		return fmt.Errorf("Search failed to parse JSON response: %s", err.Error())
	}

	return nil
}

func (client *Client) LookupTrack(uri string) (*Track, error) {
	var data struct {
		Track Track `json:"track"`
	}

	err := client.doLookup(uri, &data)
	if err != nil {
		return nil, err
	}
	return &data.Track, nil
}

func (client *Client) LookupAlbum(uri string) (*Album, error) {
	var data struct {
		Album Album `json:"album"`
	}

	err := client.doLookup(uri, &data)
	if err != nil {
		return nil, err
	}
	return &data.Album, nil
}

func (client *Client) LookupArtist(uri string) (*Artist, error) {
	var data struct {
		Artist Artist `json:"artist"`
	}

	err := client.doLookup(uri, &data)
	if err != nil {
		return nil, err
	}
	return &data.Artist, nil
}

func (client *Client) doLookup(uri string, data interface{}) error {
	url := fmt.Sprintf("http://ws.spotify.com/lookup/1/.json?uri=%s", url.QueryEscape(uri))
	resp, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("Lookup failed with http error: %s", err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&data); err != nil {
		return fmt.Errorf("Lookup failed to parse JSON response: %s", err.Error())
	}

	return nil

}
