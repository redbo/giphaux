package shared

import "io"

// DataStore is the controller API for interacting with backend app data.
// Sqlite is the only implementation, but I've tried to keep the interface
// non-leaky and abstract.  Note that this never works.
type DataStore interface {
	RandomID() string

	NewUser(username, password string) (*User, error)
	GetUser(username string) (*User, error)
	GetUserByKey(key string) (*User, error)
	GetUserByCookie(cookie string) (*User, error)
	UserFrontpage(username string) (*FrontPageData, error)
	Frontpage() (*FrontPageData, error)

	AddGIF(username string, caption string, tags []string, cats []string, sourceURL string,
		rating string, width, height, size, frames int, filedata []byte) (*GIF, error)
	RemoveGIF(username string, gifid string) error

	Search(query string, limit int, offset int, rating string) ([]*GIF, int, error)
	Trending(limit int, off int, rating string) ([]*GIF, int, error)
	RandomSearch(q string, weirdness int) (*GIF, error)
	RandomByTag(tag string, rating string) (*GIF, error)
	GIFByID(id string) (*GIF, error)
	GIFsByID(ids []string, limit, offset int) ([]*GIF, int, error)
	GIFData(id string) ([]byte, error)

	UserGIFInfo(username string, gifid string) (*UserGIFInfo, error)

	AddCategory(username string, title string) error
	RemoveCategory(username string, category string) error
	UpdateCategories(username, gifid string, categories []string) error

	AddFavorite(username string, gifid string, categories []string) error
	RemoveFavorite(username string, gifid string) error
}

// TemplateExecuter is an object that renders templates based on name and data.
type TemplateExecuter interface {
	ExecuteTemplate(io.Writer, string, interface{}) error
}

// User represents a user in the system.
type User struct {
	Username   string
	APIKey     string
	Password   string
	Cookie     string
	Categories []string
}

// UserGIFInfo returns information for a user's relationship to an image.
type UserGIFInfo struct {
	Owner      bool
	Favorited  bool
	Categories map[string]bool
}

// FrontPageData contains information for displaying a user's home page.
type FrontPageData struct {
	Categories map[string][]*GIF
}

// GIF contains all the fields in a data response from the Giphy API
type GIF struct {
	Type             string    `json:"type"`
	ID               string    `json:"id"`
	URL              string    `json:"url"`
	BitlyGifURL      string    `json:"bitly_gif_url"`
	BitlyURL         string    `json:"bitly_url"`
	EmbedURL         string    `json:"embed_url"`
	Username         string    `json:"username"`
	Source           string    `json:"source"`
	Rating           string    `json:"rating"`
	Caption          string    `json:"caption"`
	ContentURL       string    `json:"content_url"`
	ImportDatetime   string    `json:"import_datetime"`
	TrendingDatetime string    `json:"trending_datetime"`
	Images           Images    `json:"images"`
	Analytics        Analytics `json:"analytics"`
	Tags             []string  `json:"-"`
}

// Images represents all the different types of images in the Giphy API.
type Images struct {
	FixedHeight            *Image `json:"fixed_height,omitempty"`
	FixedHeightStill       *Image `json:"fixed_height_still,omitempty"`
	FixedHeightDownsampled *Image `json:"fixed_height_downsampled,omitempty"`
	FixedWidth             *Image `json:"fixed_width,omitempty"`
	FixedWidthStill        *Image `json:"fixed_width_still,omitempty"`
	FixedWidthDownsampled  *Image `json:"fixed_width_downsampled,omitempty"`
	Downsized              *Image `json:"downsized,omitempty"`
	DownsizedStill         *Image `json:"downsized_still,omitempty"`
	Original               *Image `json:"original,omitempty"`
	OriginalStill          *Image `json:"original_still,omitempty"`
}

// Image represents an image in the giphy API.
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

// APIResponse represents a response to the user-facing http API.
type APIResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Meta       Meta        `json:"meta"`
}

// AnalyticsURL wraps a URL in the analytics API response object.
type AnalyticsURL struct {
	URL string `json:"url"`
}

// Analytics is an API response object that contains callback links.
type Analytics struct {
	OnClick AnalyticsURL `json:"onclick"`
	OnLoad  AnalyticsURL `json:"onload"`
	OnSent  AnalyticsURL `json:"onsent"`
}
