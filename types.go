package main

// GIF contains all the fields in a data response from the Giphy API
type GIF struct {
	Type             string `json:"type"`
	ID               string `json:"id"`
	URL              string `json:"url"`
	BitlyGifURL      string `json:"bitly_gif_url"`
	BitlyURL         string `json:"bitly_url"`
	EmbedURL         string `json:"embed_url"`
	Username         string `json:"username"`
	Source           string `json:"source"`
	Rating           string `json:"rating"`
	Caption          string `json:"caption"`
	ContentURL       string `json:"content_url"`
	ImportDatetime   string `json:"import_datetime"`
	TrendingDatetime string `json:"trending_datetime"`
	Images           Images `json:"images"`
}

// Images represents all the different types of images
type Images struct {
	FixedHeight            Image `json:"fixed_height"`
	FixedHeightStill       Image `json:"fixed_height_still"`
	FixedHeightDownsampled Image `json:"fixed_height_downsampled"`
	FixedWidth             Image `json:"fixed_width"`
	FixedWidthStill        Image `json:"fixed_width_still"`
	FixedWidthDownsampled  Image `json:"fixed_width_downsampled"`
	Downsized              Image `json:"downsized"`
	DownsizedStill         Image `json:"downsized_still"`
	Original               Image `json:"original"`
	OriginalStill          Image `json:"original_still"`
}

// Image represents an image
type Image struct {
	URL    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Size   string `json:"size,omitempty"`
	Frames string `json:"frames,omitempty"`
	Mp4    string `json:"mp4,omitempty"`
}

// Pagination represents the pagination section in a Giphy API response
type Pagination struct {
	TotalCount int `json:"total_count"`
	Count      int `json:"count"`
	Offset     int `json:"offset"`
}

// Meta represents the meta section in a Giphy API response
type Meta struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// RandomID represents the random ID section of a Giphy API response.
type RandomID struct {
	ID string `json:"random_id"`
}
