package domain

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
