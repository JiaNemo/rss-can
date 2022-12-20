package parser

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/soulteary/RSS-Can/internal/define"
	"github.com/soulteary/RSS-Can/internal/logger"
	"github.com/soulteary/RSS-Can/internal/network"
)

func ParsePageByGoQuery(data define.RemoteBodySanitized, callback func(document *goquery.Document) []define.InfoItem) define.BodyParsed {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(data.Body))

	if err != nil {
		code := define.ERROR_CODE_PARSE_CONTENT_FAILED
		status := fmt.Sprintf("%s: %s", define.ERROR_STATUS_PARSE_CONTENT_FAILED, fmt.Errorf("%w", err))
		return define.MixupBodyParsed(code, status, data.Date, nil)
	}

	code := define.ERROR_CODE_NULL
	status := define.ERROR_STATUS_NULL
	items := callback(document)
	return define.MixupBodyParsed(code, status, data.Date, items)
}

func ParsePagerByGoQuery(data define.RemoteBodySanitized, callback func(document *goquery.Document) []string) (result []string) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(data.Body))
	if err != nil {
		return result
	}
	result = callback(document)
	return result
}

func jsBridge(field string, method string, s *goquery.Selection) string {
	// TODO use allowlist
	if strings.Contains(field, ".") || strings.Contains(field, "#") || strings.Contains(field, "p") || strings.Contains(field, "a") || strings.Contains(field, "li") {
		// extract information by attributes
		find := strings.ToLower(method)
		if find == "text" {
			return strings.TrimSpace(s.Find(field).Text())
		} else if find == "html" {
			html, err := s.Find(field).Html()
			if err != nil {
				return ""
			}
			return html
		} else if find == "href" || strings.HasPrefix(find, "data-") {
			prop, exists := s.Find(field).Attr(method)
			if !exists {
				return ""
			}
			return strings.TrimSpace(prop)
		}
	}

	// if not a selector, fallback the original content
	return field
}

func GetDataAndConfigBySSR(config define.JavaScriptConfig) (result define.BodyParsed) {
	doc := network.GetRemoteDocument(config.URL, config.Charset, config.Expire, config.DisableCache)
	if doc.Body == "" {
		return result
	}
	return ParseDataAndConfigBySSR(config, doc, "")
}

func fixLink(link string, baseUrl string) (string, error) {
	if !(strings.HasPrefix("http://", link) || strings.HasPrefix("https://", link)) {
		base, err := url.Parse(baseUrl)
		if err != nil {
			logger.Instance.Infof("Parsing link failed %s", link)
			return "", err
		}

		ref, err := url.Parse(link)
		if err != nil {
			logger.Instance.Infof("Parsing ref link failed %s", link)
			return "", err
		}
		u := base.ResolveReference(ref)
		return u.String(), nil
	}
	return link, nil
}

func GetPager(config define.JavaScriptConfig, document *goquery.Document) (links []string) {
	if config.Pager != "" {
		document.Find(config.Pager).Each(func(i int, s *goquery.Selection) {
			rawLink, exist := s.Attr("href")
			if exist {
				link, err := fixLink(rawLink, config.URL)
				if err == nil {
					links = append(links, link)
				}
			}
		})
		return links
	}
	return links
}

func taskLinkInPager(pageLinks []string, link string) bool {
	for _, item := range pageLinks {
		if item == link {
			return true
		}
	}
	return false
}

func ParseDataAndConfigBySSR(config define.JavaScriptConfig, userDoc define.RemoteBodySanitized, userHtml string) (result define.BodyParsed) {
	var doc define.RemoteBodySanitized
	if userHtml != "" {
		doc.Code = define.ERROR_CODE_NULL
		doc.Status = define.ERROR_STATUS_NULL
		doc.Body = userHtml
		doc.Date = time.Now()
	} else {
		if userDoc.Code == define.ERROR_CODE_NULL {
			doc = userDoc
		}
	}

	var items []define.InfoItem

	if config.Pager != "" {
		pageLinks := ParsePagerByGoQuery(doc, func(document *goquery.Document) []string {
			return GetPager(config, document)
		})

		if len(pageLinks) > config.PagerLimit {
			if taskLinkInPager(pageLinks, config.URL) {
				pageLinks = pageLinks[:config.PagerLimit]
			} else {
				if config.PagerLimit > 2 {
					pageLinks = pageLinks[:config.PagerLimit-1]
				}
			}
		}

		for _, link := range pageLinks {
			newConfig := config
			newConfig.URL = link
			newConfig.Pager = ""
			ret := GetDataAndConfigBySSR(newConfig)
			items = append(items, ret.Body...)
		}
		code := define.ERROR_CODE_NULL
		status := define.ERROR_STATUS_NULL
		return define.MixupBodyParsed(code, status, time.Now(), items)
	}

	return ParsePageByGoQuery(doc, func(document *goquery.Document) (items []define.InfoItem) {
		document.Find(config.ListContainer).Each(func(i int, s *goquery.Selection) {
			var item define.InfoItem
			// title must exist in the config
			if config.Title != "" {
				title := jsBridge(config.Title, "text", s)
				item.Title = title

				if config.Author != "" {
					author := jsBridge(config.Author, "text", s)
					item.Author = author
				}

				if config.DateTime != "" {
					time := jsBridge(config.DateTime, "text", s)
					item.Date = time
				}

				if config.Category != "" {
					category := jsBridge(config.Category, "text", s)
					item.Category = category
				}

				if config.Description != "" {
					description := jsBridge(config.Description, "text", s)
					item.Description = description
				}

				if config.Link != "" {
					rawLink := jsBridge(config.Link, "href", s)
					link, err := fixLink(rawLink, config.URL)
					if err == nil {
						item.Link = link
					}

					// todo prepare regexp
					if config.IdByRegexp != "" {
						re := regexp.MustCompile("`" + config.IdByRegexp + "`")
						match := re.FindAllStringSubmatch(rawLink, -1)
						fmt.Println(match)
						if len(match) == 1 {
							item.ID = match[0][1]
						}
					}
				}

				// TODO bind hook action
				if config.ContentBefore.Action != "" {
					contentBefore := network.GetRemoteDocumentAsMarkdown(item.Link, config.ContentBefore.Object, config.Charset, config.Expire, config.DisableCache)
					item.Content = contentBefore
				}

				if item.Content == "" {
					if config.Content != "" {
						item.Content = config.Content
					}
				}

				items = append(items, item)
			}
		})
		return items
	})
}
