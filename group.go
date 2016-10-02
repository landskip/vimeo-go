package vimeo

// Response of Group
type Group struct {
	CreatedTime string      `json:"created_time"`
	Description string      `json:"description"`
	Header      interface{} `json:"header"`
	Link        string      `json:"link"`
	Metadata    struct {
		Connections struct {
			Users struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"users"`
			Videos struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"videos"`
		} `json:"connections"`
		Interactions struct {
			Join struct {
				Added     bool        `json:"added"`
				AddedTime string      `json:"added_time"`
				Title     interface{} `json:"title"`
				Type      string      `json:"type"`
				URI       string      `json:"uri"`
			} `json:"join"`
		} `json:"interactions"`
	} `json:"metadata"`
	ModifiedTime string      `json:"modified_time"`
	Name         string      `json:"name"`
	Pictures     interface{} `json:"pictures"`
	Privacy      struct {
		Comment string `json:"comment"`
		Forums  string `json:"forums"`
		Invite  string `json:"invite"`
		Join    string `json:"join"`
		Videos  string `json:"videos"`
		View    string `json:"view"`
	} `json:"privacy"`
	ResourceKey string `json:"resource_key"`
	URI         string `json:"uri"`
	User        struct {
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
	} `json:"user"`
}

// My group list
type GroupList struct {
	List
	Data []Group `json:"data"`
}
