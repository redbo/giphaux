package templates

import (
	"html/template"
)

//go:generate go-bindata -pkg templates .

// LoadTemplates parses all templates from generated bindata and returns them as a suite.
func LoadTemplates() (*template.Template, error) {
	funcMap := map[string]interface{}{
		"contains": func(needle string, haystack []string) bool {
			for _, s := range haystack {
				if needle == s {
					return true
				}
			}
			return false
		},
	}
	root := template.New("/").Funcs(funcMap)
	for _, src := range AssetNames() {
		if body, err := Asset(src); err != nil {
			return nil, err
		} else if _, err := root.New(src).Parse(string(body)); err != nil {
			return nil, err
		}
	}
	return root, nil
}
