package shared

import (
	"bytes"
	"fmt"
	"image/gif"
	"regexp"
	"strings"
)

var ratings = []string{"g", "pg", "pg-13", "r"}                 // Ratings are the supported ratings for images.  Basically follows movie ratings.
var multispace = regexp.MustCompile(`\s+`)                      // used as part of tag normalization
var invalidTagChars = regexp.MustCompile(`(^-+)|[^\w\-]|(-+$)`) // matches any non-word characters and padding hyphens
var validUsername = regexp.MustCompile(`^[a-z0-9\-]{3,20}$`)    // make sure usernames are 3-20 lowercase alphanumeric
var validGIFID = regexp.MustCompile(`^[a-zA-Z0-9]{8,64}$`)      // make sure gifids are 8-64 alphanumerics

// NormalizeRating makes sure the rating is valid and formats it correctly.
func NormalizeRating(rating string) (string, error) {
	rating = strings.TrimSpace(strings.ToLower(rating))
	for _, r := range ratings {
		if rating == r {
			return r, nil
		}
	}
	return "", fmt.Errorf("Unknown rating")
}

// NormalizeTag takes a string and makes it into a canonical tag, which means all lowercase, trimmed and any spaces replaced with hyphens.
func NormalizeTag(tag string) (string, error) {
	tag = strings.ToLower(tag)
	tag = multispace.ReplaceAllString(tag, "-")
	tag = invalidTagChars.ReplaceAllString(tag, "")
	if tag == "" {
		return "", fmt.Errorf("Invalid tag name")
	}
	return tag, nil
}

// NormalizeUsername returns the username lowercased and trimmed.  returns an error if the username isn't valid.
func NormalizeUsername(username string) (string, error) {
	username = strings.ToLower(strings.TrimSpace(username))
	if !validUsername.MatchString(username) {
		return "", fmt.Errorf("Invalid username")
	}
	return username, nil
}

// NormalizeGIFID returns the gifid with spaces trimmed.  returns an error if the gifid isn't valid.
func NormalizeGIFID(id string) (string, error) {
	id = strings.TrimSpace(id)
	if !validGIFID.MatchString(id) {
		return "", fmt.Errorf("Invalid gif id")
	}
	return id, nil
}

// GIFInfo parses the GIF data and returns its width, height, size, and number of frames.
func GIFInfo(filedata []byte) (int, int, int, int, error) {
	// dear future: could we just parse the header instead of decoding the entire gif into memory?
	img, err := gif.DecodeAll(bytes.NewBuffer(filedata))
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("Error parsing gif")
	}
	return img.Config.Width, img.Config.Height, int(len(filedata)), len(img.Image), nil
}
