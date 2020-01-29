package sqlite

import (
	"time"
)

// Models for the database.

// User represents a user of the site.
type User struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  *time.Time
	Username   string `gorm:"unique_index;type:TEXT"`
	Password   string `gorm:"type:TEXT"`
	APIKey     string `gorm:"index;type:TEXT"`
	Cookie     string `gorm:"index;type:TEXT"`
	Categories []Category
}

// GIF represents a gif in the system.
type GIF struct {
	ID               uint `gorm:"primary_key"`
	CreatedAt        *time.Time
	Type             string     `gorm:"type:TEXT"`
	APIID            string     `gorm:"unique_index;type:TEXT"`
	URL              string     `gorm:"type:TEXT"`
	BitlyGifURL      string     `gorm:"type:TEXT"`
	BitlyURL         string     `gorm:"type:TEXT"`
	EmbedURL         string     `gorm:"type:TEXT"`
	Source           string     `gorm:"type:TEXT"`
	Rating           string     `gorm:"type:TEXT"`
	Caption          string     `gorm:"type:TEXT"`
	ContentURL       string     `gorm:"type:TEXT"`
	ImportDatetime   *time.Time `gorm:"type:DATETIME"`
	TrendingDatetime *time.Time `gorm:"type:DATETIME"`
	UserID           uint
	User             User
	Width            int
	Height           int
	Size             int
	Frames           int
	Tags             []Tag
}

// GIFData holds the actual bytes of a GIF.
// Denormalized to reduce i/o on things like gif searches, since these records can get big.
type GIFData struct {
	ID        uint `gorm:"primary_key"`
	GIFID     uint
	CreatedAt *time.Time
	Data      []byte `gorm:"type:BLOB"`
}

// Tag represents a tag on an image.
type Tag struct {
	ID    uint   `gorm:"primary_key"`
	GIFID uint   `gorm:"unique_index:idx_unique_gif_tag"`
	Tag   string `gorm:"type:TEXT;unique_index:idx_unique_gif_tag"`
}

// Category represents a user's categories for images, such as "funny" or "cats".
type Category struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt *time.Time
	Title     string `gorm:"type:TEXT;unique_index:idx_unique_category"`
	UserID    uint   `gorm:"unique_index:idx_unique_category"`
}

// Favorite represents a GIF that's been saved to the user's profile.
type Favorite struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt *time.Time
	UserID    uint `gorm:"unique_index:idx_unique_favorite"`
	GIFID     uint `gorm:"unique_index:idx_unique_favorite"`
}

// CategorizedFavorite represents a GIF that's been saved to the user's profile and categorized.
type CategorizedFavorite struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  *time.Time
	FavoriteID uint `gorm:"unique_index:idx_unique_categorized_favorite"`
	CategoryID uint `gorm:"unique_index:idx_unique_categorized_favorite"`
}
