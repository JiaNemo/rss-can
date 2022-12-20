package define

import "time"

const (
	FEED_TYPE_RSS  = "rss"
	FEED_TYPE_ATOM = "atom"
	FEED_TYPE_JSON = "json"
)

const (
	FEED_MIME_TYPE_RSS     = "application/rss+xml"
	FEED_MIME_TYPE_ATOM    = "application/atom+xml"
	FEED_MIME_TYPE_JSON    = "application/feed+json"
	FEED_MIME_TYPE_DEFAULT = "text/plain"
)

type InfoItem struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Date        string `json:"date"`
	Author      string `json:"author,omitempty"`
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
}

type JavaScriptConfig struct {
	URL          string        `json:"URL"`
	Mode         string        `json:"Mode"`
	File         string        //private field
	Charset      string        `json:"Charset"`
	Expire       time.Duration `json:"Expire"`
	Headless     string        `json:"Headless"`
	DisableCache bool          `json:"DisableCache"`
	IdByRegexp   string        `json:"IdByRegexp"`

	ListContainer string `json:"ListContainer"`
	Title         string `json:"Title"`
	Author        string `json:"Author"`
	Category      string `json:"Category"`
	DateTime      string `json:"DateTime"`
	Description   string `json:"Description"`
	Link          string `json:"Link"`

	Content       string `json:"Content"`
	ContentBefore struct {
		Action string `json:"action"`
		Object string `json:"object"`
		URL    string `json:"URL"`
	} `json:"ContentBefore"`

	Pager      string `json:"Pager"`
	PagerLimit int    `json:"PagerLimit"`
}
