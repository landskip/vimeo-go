package vimeo

// Response of [GET]/me
type Me struct {
	Account       string      `json:"account"`
	Bio           interface{} `json:"bio"`
	ContentFilter []string    `json:"content_filter"`
	CreatedTime   string      `json:"created_time"`
	Link          string      `json:"link"`
	Location      interface{} `json:"location"`
	Metadata      struct {
		Connections struct {
			Activities struct {
				Options []string `json:"options"`
				URI     string   `json:"uri"`
			} `json:"activities"`
			Albums struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"albums"`
			Appearances struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"appearances"`
			Categories struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"categories"`
			Channels struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"channels"`
			Feed struct {
				Options []string `json:"options"`
				URI     string   `json:"uri"`
			} `json:"feed"`
			Followers struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"followers"`
			Following struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"following"`
			Groups struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"groups"`
			Likes struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"likes"`
			ModeratedChannels struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"moderated_channels"`
			Pictures struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"pictures"`
			Portfolios struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"portfolios"`
			Shared struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"shared"`
			Videos struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"videos"`
			WatchedVideos struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"watched_videos"`
			Watchlater struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"watchlater"`
		} `json:"connections"`
	} `json:"metadata"`
	Name        string      `json:"name"`
	Pictures    interface{} `json:"pictures"`
	Preferences struct {
		Videos struct {
			Privacy string `json:"privacy"`
		} `json:"videos"`
	} `json:"preferences"`
	ResourceKey string        `json:"resource_key"`
	URI         string        `json:"uri"`
	Websites    []interface{} `json:"websites"`
}
