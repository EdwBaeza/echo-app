package domain

import "fmt"

// User model
type User struct {
	ID    string `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}

type UserPage struct {
	Page
	Data []*User `json:"data"`
}

func (page *UserPage) SetLinks() {
	page.Links = &Links{
		Self: fmt.Sprintf("https://example.com/users/%d", page.PageNumber),
		Next: fmt.Sprintf("https://example.com/users/%d", page.PageNumber),
		Prev: fmt.Sprintf("https://example.com/users/%d", page.PageNumber),
	}
}

func (page *UserPage) GetNextLink() *string {

	return nil
}
