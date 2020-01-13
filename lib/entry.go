package pages

import (
	core "github.com/Hatch1fy/service-core"
)

func newEntry(name, key, template string, d Data) (e Entry) {
	e.Name = name
	e.Key = key
	e.Template = template
	e.Data = d
	return
}

// Entry represents a pages entry
type Entry struct {
	core.Entry

	// Key is the page's key (sanitized name)
	Key string `json:"key"`
	// Name is the chosen name of the page
	Name string `json:"name"`
	// Template is the name of the template associated with the page
	Template string `json:"template"`
	// Data represents any data associated with the page
	Data Data `json:"data"`

	// RedirectTo notifies which URL to redirect to
	RedirectTo string `json:"redirectTo,omitempty"`
}

// Validate will validate en entry
func (e *Entry) Validate() (err error) {
	if len(e.Name) == 0 {
		return ErrEmptyName
	}

	if len(e.Key) == 0 {
		return ErrEmptyKey
	}

	if len(e.Template) == 0 {
		return ErrEmptyTemplate
	}

	return
}

// GetRelationshipIDs is a service-core helper method for getting relationship IDs of an entry
func (e *Entry) GetRelationshipIDs() (ids []string) {
	ids = append(ids, e.Key)
	return
}
