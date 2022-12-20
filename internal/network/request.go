package network

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/soulteary/RSS-Can/internal/cacher"
	"github.com/soulteary/RSS-Can/internal/define"
	"github.com/soulteary/RSS-Can/internal/fn"
	"github.com/soulteary/RSS-Can/internal/logger"
)

func Get(url string, userAgent string) (code define.ErrorCode, status string, response *http.Response) {
	client := &http.Client{Timeout: define.GLOBAL_REQ_TIMEOUT}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		code = define.ERROR_CODE_INIT_NETWORK_FAILED
		status = fmt.Sprintf("%s: %s", define.ERROR_STATUS_INIT_NETWORK_FAILED, fmt.Errorf("%w", err))
		return code, status, response
	}

	req.Header.Set("User-Agent", userAgent)

	response, err = client.Do(req)
	if err != nil {
		code = define.ERROR_CODE_NETWORK
		status = fmt.Sprintf("%s: %s", define.ERROR_STATUS_NETWORK, fmt.Errorf("%w", err))
		return code, status, response
	}

	code = define.ERROR_CODE_NULL
	status = define.ERROR_STATUS_NULL
	return code, status, response
}

func GetRemoteDocument(url string, charset string) define.RemoteBodySanitized {
	var code define.ErrorCode
	var status string
	var now = time.Now()

	cache, err := cacher.Get(url)
	if err == nil && cache != "" {
		logger.Instance.Debugln("Get remote document from cache")
		code = define.ERROR_CODE_NULL
		status = define.ERROR_STATUS_NULL
		return define.MixupRemoteBodySanitized(code, status, now, cache)
	}

	code, status, res := Get(url, define.GLOBAL_USER_AGENT)
	if code != define.ERROR_CODE_NULL {
		return define.MixupRemoteBodySanitized(code, status, now, "")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		code = define.ERROR_CODE_API_NOT_READY
		status = fmt.Sprintf("%s: %d %s", define.ERROR_STATUS_API_NOT_READY, res.StatusCode, res.Status)
		return define.MixupRemoteBodySanitized(code, status, now, "")
	}

	bodyParsed, err := fn.DecodeHTMLBody(res.Body, charset)
	if err != nil {
		code = define.ERROR_CODE_DECODE_CAHRSET_FAILED
		status = fmt.Sprintf("%s: %s", define.ERROR_STATUS_DECODE_CAHRSET_FAILED, fmt.Errorf("%w", err))
		return define.MixupRemoteBodySanitized(code, status, now, "")
	}

	code = define.ERROR_CODE_NULL
	status = define.ERROR_STATUS_NULL
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(bodyParsed)

	err = cacher.Set(url, buffer.String())
	if err != nil {
		logger.Instance.Warn("Unable to use cache")
	} else {
		// TODO set with rule config
		// 10mins
		cacher.Expire(url, 10*60*time.Second)
	}
	return define.MixupRemoteBodySanitized(code, status, now, buffer.String())
}

func GetRemoteDocumentAsMarkdown(url string, selector string, charset string) string {
	doc := GetRemoteDocument(url, charset)
	if doc.Body == "" {
		return ""
	}

	document, err := goquery.NewDocumentFromReader(strings.NewReader(doc.Body))
	if err != nil {
		return ""
	}

	// default selector use whole document body
	if selector == "" {
		selector = "body"
	}
	html, err := document.Find(selector).Html()
	if err != nil {
		return ""
	}

	md, err := fn.Html2Md(html)
	if err != nil {
		return ""
	}

	return md
}
