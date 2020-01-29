package shared

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

	AddGIF(username string, caption string, tags []string, cats []string, sourceURL string, rating string, width, height int) (*GIF, error)
	RemoveGIF(username string, gifid string) error
	UpdateCategories(username, gifid string, categories []string) error

	Search(query string, limit int, offset int, rating string) ([]*GIF, int, error)
	Trending(limit int, off int, rating string) ([]*GIF, int, error)
	RandomSearch(q string, weirdness int) (*GIF, error)
	RandomByTag(tag string, rating string) (*GIF, error)
	GIFByID(id string) (*GIF, error)
	GIFsByID(ids []string) ([]*GIF, error)

	UserGIFInfo(username string, gifid string) (*UserGIFInfo, error)

	AddCategory(username string, title string) error
	RemoveCategory(username string, category string) error

	AddFavorite(username string, gifid string, categories []string) error
	RemoveFavorite(username string, gifid string) error
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
	Type             string   `json:"type"`
	ID               string   `json:"id"`
	URL              string   `json:"url"`
	BitlyGifURL      string   `json:"bitly_gif_url"`
	BitlyURL         string   `json:"bitly_url"`
	EmbedURL         string   `json:"embed_url"`
	Username         string   `json:"username"`
	Source           string   `json:"source"`
	Rating           string   `json:"rating"`
	Caption          string   `json:"caption"`
	ContentURL       string   `json:"content_url"`
	ImportDatetime   string   `json:"import_datetime"`
	TrendingDatetime string   `json:"trending_datetime"`
	Images           Images   `json:"images"`
	Tags             []string `json:"-"`
}

// Images represents all the different types of images in the Giphy API.
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
