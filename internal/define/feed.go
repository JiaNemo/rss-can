package define

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

type InfoItemHQY struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Link        string `json:"url"`
	Date        string `json:"publishDate"`
	Author      string `json:"author,omitempty"`
	Category    string `json:"catalogName,omitempty"`
	Description string `json:"summary,omitempty"`
	Content     string `json:"content,omitempty"`
}

const (
	ConfigHookReadLink = "readlink"
)

type ConfigHook struct {
	Action string `json:"action"`
	Object string `json:"object"`
	URL    string `json:"URL"`
}

type PROP_ID struct {
	Object string `json:"object"`
	Prop   string `json:"prop"`
}

type JavaScriptConfig struct {
	Name       string `json:"name"`
	URL        string `json:"URL"`
	OriginHost string `json:"OriginHost"`
	Host       string `json:"Host"`
	Logo       string `json:"Logo"`

	Mode         string  `json:"Mode"`
	File         string  //private field
	Charset      string  `json:"Charset"`
	Expire       int     `json:"Expire"`
	Headless     string  `json:"Headless"`
	Timeout      int     `json:"Timeout"`
	Proxy        string  `json:"Proxy"`
	DisableCache bool    `json:"DisableCache"`
	IdByRegexp   string  `json:"IdByRegexp"`
	IdByProp     PROP_ID `json:"IdByProp"`
	Cookies      string  `json:"Cookies"`

	ListContainer string     `json:"ListContainer"`
	Title         string     `json:"Title"`
	Author        string     `json:"Author"`
	Link          string     `json:"Link"`
	DateTime      string     `json:"DateTime"`
	DateTimeHook  ConfigHook `json:"DateTimeHook"`
	CategoryHook  ConfigHook `json:"CategoryHook"`
	Description   string     `json:"Description"`

	DescriptionHook ConfigHook `json:"DescriptionHook"`
	Content         string     `json:"Content"`

	ContentHook ConfigHook `json:"ContentHook"`
	Pager       string     `json:"Pager"`

	PagerLimit int `json:"PagerLimit"`

	TenantId        string `json:"TenantId"`
	Category        string `json:"Category"`
	CataLogId       string `json:"CataLogId"`
	PageSize        string `json:"PageSize"`
	MaxPageCount    string `json:"MaxPageCount"`
	CopyRight       string `json:"CopyRight"`
	SubTitle        string `json:"SubTitle"`
	MainDescription string `json:"MainDescription"`
}
