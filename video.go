package vimeo

import "net/url"
import "encoding/json"

// Video Type Option
const (
	TypePost      = "POST"
	TypeStreaming = "streaming"
	TypePull      = "pull"
)

// Request of [POST] /videos
type VideoOptions struct {
	Type        string `json:"type,omitempty"`
	RedirectURL string `json:"RedirectURL,omitempty"`
	Link        string `json:"Link,omitempty"`
}

func (o *VideoOptions) GetJSONBytes() ([]byte, error) {
	return json.Marshal(o)
}

func (o *VideoOptions) GetValues() url.Values {
	return nil
}

// Response of [GET] /me/videos/{video_id}
type Video struct {
	App           interface{} `json:"app"`
	ContentRating []string    `json:"content_rating"`
	CreatedTime   string      `json:"created_time"`
	Description   interface{} `json:"description"`
	Download      []struct {
		CreatedTime string  `json:"created_time"`
		Expires     string  `json:"expires"`
		Fps         float64 `json:"fps"`
		Height      int     `json:"height"`
		Link        string  `json:"link"`
		Md5         string  `json:"md5"`
		Quality     string  `json:"quality"`
		Size        int     `json:"size"`
		Type        string  `json:"type"`
		Width       int     `json:"width"`
	} `json:"download"`
	Duration int `json:"duration"`
	Embed    struct {
		Buttons struct {
			Embed      bool `json:"embed"`
			Fullscreen bool `json:"fullscreen"`
			Hd         bool `json:"hd"`
			Like       bool `json:"like"`
			Scaling    bool `json:"scaling"`
			Share      bool `json:"share"`
			Watchlater bool `json:"watchlater"`
		} `json:"buttons"`
		Color string `json:"color"`
		HTML  string `json:"html"`
		Logos struct {
			Custom struct {
				Active bool        `json:"active"`
				Link   interface{} `json:"link"`
				Sticky bool        `json:"sticky"`
			} `json:"custom"`
			Vimeo bool `json:"vimeo"`
		} `json:"logos"`
		Playbar bool `json:"playbar"`
		Title   struct {
			Name     string `json:"name"`
			Owner    string `json:"owner"`
			Portrait string `json:"portrait"`
		} `json:"title"`
		URI    interface{} `json:"uri"`
		Volume bool        `json:"volume"`
	} `json:"embed"`
	EmbedPresets interface{} `json:"embed_presets"`
	Files        []struct {
		CreatedTime string  `json:"created_time"`
		Fps         float64 `json:"fps"`
		Height      int     `json:"height"`
		Link        string  `json:"link"`
		LinkSecure  string  `json:"link_secure"`
		Md5         string  `json:"md5"`
		Quality     string  `json:"quality"`
		Size        int     `json:"size"`
		Type        string  `json:"type"`
		Width       int     `json:"width"`
	} `json:"files"`
	Height   int         `json:"height"`
	Language interface{} `json:"language"`
	License  interface{} `json:"license"`
	Link     string      `json:"link"`
	Metadata struct {
		Connections struct {
			Comments struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"comments"`
			Credits interface{} `json:"credits"`
			Likes   struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"likes"`
			Pictures struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"pictures"`
			Related    interface{} `json:"related"`
			Texttracks struct {
				Options []string `json:"options"`
				Total   int      `json:"total"`
				URI     string   `json:"uri"`
			} `json:"texttracks"`
		} `json:"connections"`
		Interactions struct {
			Watchlater struct {
				Added     bool        `json:"added"`
				AddedTime interface{} `json:"added_time"`
				URI       string      `json:"uri"`
			} `json:"watchlater"`
		} `json:"interactions"`
	} `json:"metadata"`
	ModifiedTime string `json:"modified_time"`
	Name         string `json:"name"`
	Pictures     struct {
		Active      bool   `json:"active"`
		ResourceKey string `json:"resource_key"`
		Sizes       []struct {
			Height             int    `json:"height"`
			Link               string `json:"link"`
			LinkWithPlayButton string `json:"link_with_play_button"`
			Width              int    `json:"width"`
		} `json:"sizes"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"pictures"`
	Privacy struct {
		Add      bool   `json:"add"`
		Comments string `json:"comments"`
		Download bool   `json:"download"`
		Embed    string `json:"embed"`
		View     string `json:"view"`
	} `json:"privacy"`
	ReleaseTime string `json:"release_time"`
	ResourceKey string `json:"resource_key"`
	ReviewLink  string `json:"review_link"`
	Stats       struct {
		Plays int `json:"plays"`
	} `json:"stats"`
	Status string        `json:"status"`
	Tags   []interface{} `json:"tags"`
	URI    string        `json:"uri"`
	User   struct {
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
	Width int `json:"width"`
}

// My video list
type VideoList struct {
	List
	Data []Video `json:"data"`
}
