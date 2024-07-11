package parser

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/soulteary/RSS-Can/internal/define"
	"github.com/soulteary/RSS-Can/internal/network"
	"math"
	"strconv"
	"strings"
	"time"
)

func GetDataAndConfigBySSRHQY(config define.JavaScriptConfig) (result define.BodyParsed) {
	doc := network.GetHQYDocument(config.URL, config.Charset, config.Expire, config.DisableCache, config, 1)
	if doc.Body == "" {
		return result
	}
	return ParseDataAndConfigBySSRHQY(config, doc, "")
}

func ParseDataAndConfigBySSRHQY(config define.JavaScriptConfig, userDoc define.RemoteBodySanitized, userHtml string) (result define.BodyParsed) {
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
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var items []define.InfoItem
	pages := 1
	sum := json.Get([]byte(doc.Body), "data").Get("totalRecords").ToInt()
	pageSize := json.Get([]byte(doc.Body), "data").Get("pageSize").ToInt()
	pageCount := int(math.Ceil(float64(sum) / float64(pageSize)))
	pages = pageCount

	if config.Pager != "" {
		if pageCount > config.PagerLimit {
			pages = config.PagerLimit
		}
	}

	maxPageCount, _ := strconv.Atoi(config.MaxPageCount)

	if pages >= maxPageCount {
		pages = maxPageCount
	}

	for i := 1; i <= pages; i++ {
		ret := network.GetHQYDocument(config.URL, config.Charset, config.Expire, config.DisableCache, config, i)
		var arr []define.InfoItemHQY
		err := json.Unmarshal([]byte(json.Get([]byte(ret.Body), "data").Get("pageRecords").ToString()), &arr)
		if err == nil {
			for _, hqy := range arr {
				var i define.InfoItem
				i.ID = strconv.Itoa(hqy.ID)
				i.Title = hqy.Title
				i.Link = strings.ReplaceAll(hqy.Link, config.OriginHost, config.Host)
				i.Date = hqy.Date
				i.Author = hqy.Author
				i.Category = hqy.Category
				i.Description = hqy.Description
				i.Content = hqy.Content
				items = append(items, i)
			}
		}
	}

	code := define.ERROR_CODE_NULL
	status := define.ERROR_STATUS_NULL
	return define.MixupBodyParsed(code, status, time.Now(), items)
}
