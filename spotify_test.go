package spotify

import (
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {

	s := New(&http.Client{})

	_, err := s.SearchArtists("foo", 0)

	if err != nil {
		t.Errorf("error not nil: ", err)
	}
}

func TestLookUp(t *testing.T) {

	s := New(&http.Client{})

	_, err := s.LookupTrack("spotify:track:6NmXV4o6bmp704aPGyTVVG")

	if err != nil {
		t.Errorf("error not nil: ", err)
	}

}
