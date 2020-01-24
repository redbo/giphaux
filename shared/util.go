package shared

import (
	"fmt"
	"regexp"
	"strings"
)

var multispace = regexp.MustCompile(`\s+`)                      // used as part of tag normalization
var invalidTagChars = regexp.MustCompile(`(^-+)|[^\w\-]|(-+$)`) // matches any non-word characters and padding hyphens
var validUsername = regexp.MustCompile(`^[a-z0-9\-]{3,20}$`)    // make sure usernames are 3-20 lowercase alphanumeric
var validGIFID = regexp.MustCompile(`^[a-zA-Z0-9]{8,64}$`)      // make sure gifids are 8-64 alphanumerics

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
