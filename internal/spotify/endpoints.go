package spotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Do get Request.
// Caller must close response body.
func get(endpoint string, token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // !
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	return resp, nil
}

// Search for artists on Spotify.
func SearchArtists(query string, token string) ([]ArtistObject, error) {
	endpoint := "https://api.spotify.com/v1/search?"

	params := url.Values{}
	params.Set("q", query)
	params.Set("type", "artist")
	params.Set("limit", "20")

	resp, err := get(endpoint+params.Encode(), token)
	if err != nil {
		return nil, fmt.Errorf("search query failed: %v", err)
	}
	defer resp.Body.Close()

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Artists.Items, nil
}

// Search for tracks on Spotify.
func SearchTracks(query string, token string) ([]TrackObject, error) {
	endpoint := "https://api.spotify.com/v1/search?"

	params := url.Values{}
	params.Set("q", query)
	params.Set("type", "track")
	params.Set("limit", "20")

	resp, err := get(endpoint+params.Encode(), token)
	if err != nil {
		return nil, fmt.Errorf("search query failed: %v", err)
	}
	defer resp.Body.Close()

	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Tracks.Items, nil
}
