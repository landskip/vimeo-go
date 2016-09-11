package vimeo

// Own client
// https://developer.vimeo.com/api/endpoints/me#/ondemand/pages
type MeClient struct{}

// Me
// TODO: add Metadata
type User struct {
	URI         string      `json:"uri"`
	Name        string      `json:"name"`
	Link        string      `json:"link"`
	Location    string      `json:"location"`
	Bio         string      `json:"bio"`
	CreatedTime string      `json:"created_time"`
	Account     string      `json:"account"`
	Pictures    string      `json:"pictures"`
	Websites    []string    `json:"websites"`
	Preferences Preferences `json:"preferences"`
}

// Preferences
type Preferences struct {
	Videos Privacy `json:"videos"`
}

// Privacy
type Privacy struct {
	Privacy string `json:"privacy"`
}

// View one user
func (c *MeClient) ShowInfomation() (*User, error) {
	user := User{}
	err := query("GET", "/me", nil, &user)
	return &user, err
}
