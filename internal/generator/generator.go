package generator

import (
	"time"

	"github.com/gorilla/feeds"
	"github.com/soulteary/RSS-Can/internal/define"
	"github.com/soulteary/RSS-Can/internal/jssdk"
	"github.com/soulteary/RSS-Can/internal/logger"
)

func GenerateFeedsByType(config define.JavaScriptConfig, data define.BodyParsed, rssType string) string {
	now := time.Now()

	rssFeed := &feeds.Feed{
		Title:   config.File,
		Created: now,
	}

	if config.Host != "" {
		rssFeed.Link = &feeds.Link{Href: config.Host}
	}

	if config.Logo != "" {
		rssFeed.Image = &feeds.Image{
			Url:    config.Logo,
			Title:  config.Name,
			Link:   config.Host,
			Width:  200,
			Height: 200,
		}
	}

	if config.Name != "" {
		rssFeed.Title = config.Name
	}

	if config.CopyRight != "" {
		rssFeed.Copyright = config.CopyRight
	}

	if config.SubTitle != "" {
		rssFeed.Subtitle = config.SubTitle
	}

	if config.MainDescription != "" {
		rssFeed.Description = config.MainDescription
	}

	logger.Instance.Infof("Frank len data body : %d", len(data.Body))

	for _, data := range data.Body {
		feedItem := feeds.Item{
			Title:       data.Title,
			Author:      &feeds.Author{Name: data.Author},
			Description: data.Description,
			Link:        &feeds.Link{Href: data.Link},
		}

		if data.Date != "" {
			timeUnix, err := jssdk.ConvertStrToUnix(data.Date)
			if err == nil {
				feedItem.Created = timeUnix
			}
		}

		if data.ID != "" {
			feedItem.Id = data.ID
		}

		if data.Content != "" {
			feedItem.Content = data.Content
		}

		rssFeed.Items = append(rssFeed.Items, &feedItem)
	}

	var rss string
	var err error

	switch rssType {
	case define.FEED_TYPE_RSS:
		rss, err = rssFeed.ToRss()
	case define.FEED_TYPE_ATOM:
		rss, err = rssFeed.ToAtom()
	case define.FEED_TYPE_JSON:
		rss, err = rssFeed.ToJSON()
	default:
		rss = ""
	}

	if err != nil {
		logger.Instance.Errorf("Generate feed failed: %v", err)
		return ""
	}

	logger.Instance.Infof("Frank data count : %d", len(rssFeed.Items))

	return rss
}
