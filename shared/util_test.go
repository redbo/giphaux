package shared

import "testing"

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
			t.Logf("ValidateUsername(%q) -> %v, %v", testcase.test, tag, err)
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
			t.Logf("ValidateUsername(%q) -> %v, %v", testcase.test, username, err)
			t.Fail()
		}
	}
}
