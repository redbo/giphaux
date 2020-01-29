package shared

import (
	"encoding/base64"
	"testing"
)

func TestNormalizeTag(t *testing.T) {
	tests := []struct {
		valid      bool
		test       string
		normalized string
	}{
		{true, "a test", "a-test"},
		{true, "a-test", "a-test"},
		{true, " a test ", "a-test"},
		{true, " a#test ", "atest"},
		{true, "i can't even ", "i-cant-even"},
		{false, "", ""},
		{false, " \n", ""},
	}
	for _, testcase := range tests {
		if tag, err := NormalizeTag(testcase.test); (err != nil && testcase.valid) || (err == nil && !testcase.valid) {
			t.Logf("NormalizeTag(%q) -> %v, %v", testcase.test, tag, err)
			t.Fail()
		}
	}
}

func TestNormalizeUsername(t *testing.T) {
	tests := []struct {
		valid      bool
		test       string
		normalized string
	}{
		{true, "mike", "mike"},
		{true, "mike79", "mike79"},
		{true, "mike-79", "mike-79"},
		{true, "MIKE", "mike"},
		{false, "extremelylongusername", ""},
		{false, "mi", ""},
		{false, "@mike", ""},
	}
	for _, testcase := range tests {
		if username, err := NormalizeUsername(testcase.test); (err != nil && testcase.valid) || (err == nil && !testcase.valid) {
			t.Logf("NormalizeUsername(%q) -> %v, %v", testcase.test, username, err)
			t.Fail()
		}
	}
}

func TestNormalizeRating(t *testing.T) {
	tests := []struct {
		valid      bool
		test       string
		normalized string
	}{
		{true, "G", "g"},
		{true, "g", "g"},
		{true, "PG-13", "pg-13"},
		{true, "R", "r"},
		{false, "nc-17", ""},
		{false, "", ""},
	}
	for _, testcase := range tests {
		if username, err := NormalizeRating(testcase.test); (err != nil && testcase.valid) || (err == nil && !testcase.valid) {
			t.Logf("NormalizeRating(%q) -> %v, %v", testcase.test, username, err)
			t.Fail()
		}
	}
}

func TestGIFID(t *testing.T) {
	tests := []struct {
		valid      bool
		test       string
		normalized string
	}{
		{true, "abcdefghijklmnopqrstuv", "abcdefghijklmnopqrstuv"},
		{true, " 1234567890abc ", "1234567890abc"},
		{false, "", ""},
		{false, "abc 123", ""},
		{false, "1234567890abc*", ""},
	}
	for _, testcase := range tests {
		if gifid, err := NormalizeGIFID(testcase.test); (err != nil && testcase.valid) || (err == nil && !testcase.valid) {
			t.Logf("NormalizeGIFID(%q) -> %q, %v", testcase.test, gifid, err)
			t.Fail()
		}
	}
}

func TestGIFInfo(t *testing.T) {
	gifdata, err := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7")
	if err != nil {
		t.Fail()
	}
	width, height, size, frames, err := GIFInfo(gifdata)
	if err != nil {
		t.Fail()
	}
	if err != nil || width != 1 || height != 1 || frames != 1 || size != 42 {
		t.Logf("GIFInfo([stuff]) -> %v, %v, %v, %v, %v", width, height, size, frames, err)
		t.Fail()
	}
	_, _, _, _, err = GIFInfo([]byte("junk"))
	if err == nil {
		t.Logf("GIFInfo([junk]) did not return an error")
		t.Fail()
	}
}
