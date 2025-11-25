package spotify

import (
	"music/internal/image"
)

type SearchResult struct {
	Artists Search[ArtistObject]
	Tracks  Search[TrackObject]
}

type Search[O any] struct {
	// A link to the Web API endpoint returning
	// the full result of the request
	Href string
	// The maximum number of items in the response
	// (as set in the query or by default).
	Limit int
	// URL to the next page of items.
	// ( null if none)
	Next string
	// The offset of the items returned
	// (as set in the query or by default)
	Offset int
	// URL to the previous page of items.
	// ( null if none)
	Previous string
	// The total number of items available to return.
	Total int
	// Array of Object
	Items []O
}

type ArtistObject struct {
	// A list of the genres the artist is associated with.
	// If not yet classified, the array is empty.
	Genres []string
	// A link to the Web API endpoint providing full details of the artist.
	Href string
	// The Spotify ID for the artist.
	Id string
	// The name of the artist.
	Name string
	// The popularity of the artist. The value will be between 0 and 100,
	// with 100 being the most popular. The artist's popularity is calculated
	// from the popularity of all the artist's tracks.
	Popularity int
	// The Spotify URI for the artist.
	Uri string
}

type TrackObject struct {
	// The album which the track appears.
	Album AlbumObject
	// The artists who performed on the track.
	Artists []ArtistObject
	// The disc number.
	DiscNumber int `json:"disc_number"`
	// The track length in milliseconds.
	DurationMs int `json:"duration_ms"`
	// Whether or not the track has explicit lyrics.
	Explicit bool
	// A link to the Web API endpoint providing full details of the track.
	Href string
	// The Spotify ID for the track.
	Id string
	// The name of the track.
	Name string
	// The popularity of the track.
	// The value will be between 0 and 100, with 100 being the most popular.
	Popularity int
	// The number of the track. If an album has several discs,
	// the track number is the number on the specified disc.
	TrackNumber int `json:"track_number"`
	// The Spotify URI for the track.
	Uri string
}

type AlbumObject struct {
	// The type of the album.
	AlbumType string `json:"album_type"`
	// The number of tracks in the album.
	TotalTracks int `json:"total_tracks"`
	// A link to the Web API endpoint providing full details of the album.
	Href string
	// The Spotify ID for the album.
	Id string
	// The cover art for the album in various sizes, widest first.
	Images []ImageObject
	// The name of the album.
	Name string
	// The date the album was first released.
	ReleaseDate string `json:"release_date"`
	// The Spotify URI for the album.
	Uri string
	// The artists of the album.
	Artists []ArtistObject
}

type ImageObject struct {
	// The source URL of the image.
	Url string
	// The image height in pixels.
	Height int
	// The image width in pixels.
	Width int
}

// =============================================================================
// Implement the list.Item interface for TrackObject
// =============================================================================

func (t TrackObject) FilterValue() string {
	return t.Artists[0].Name
}

func (t TrackObject) Title() string {
	return t.Name // + " : " + t.Id
}

func (t TrackObject) Description() string {
	img, err := image.GetImage(t.Album.Images[0].Url)
	if err != nil {
		panic("foo")
	}
	return image.DrawImage(img, 22, 22)
	//return t.Artists[0].Name
}
