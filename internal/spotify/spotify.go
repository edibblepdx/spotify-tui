package spotify

type SearchResult struct {
	Artists Artists
}

type Artists struct {
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
	// Array of ArtistObject
	Items []ArtistObject
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
