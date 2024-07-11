package rule

import (
	"encoding/json"
	"strings"

	"github.com/soulteary/RSS-Can/internal/define"
	"github.com/soulteary/RSS-Can/internal/fn"
	"github.com/soulteary/RSS-Can/internal/parser"
)

func ParseConfigFromJSON(str string, rule define.RuleCache) (define.JavaScriptConfig, error) {
	var config define.JavaScriptConfig
	err := json.Unmarshal([]byte(str), &config)
	if err != nil {
		return config, err
	}

	if rule.File != "" {
		config.File = rule.File
	}

	modeInRule := strings.ToLower(config.Mode)
	if !fn.IsStrInArray([]string{define.PARSE_MODE_SSR, define.PARSE_MODE_CSR, define.PARSE_MODE_MIX, define.PARSE_MODE_SSR_HQY}, modeInRule) {
		config.Mode = define.DEFAULT_PARSE_MODE
	}
	return config, nil
}

func GetWebsiteDataWithConfig(config define.JavaScriptConfig) (result define.BodyParsed) {
	switch strings.ToLower(config.Mode) {
	case define.PARSE_MODE_SSR:
		return parser.GetDataAndConfigBySSR(config)
	case define.PARSE_MODE_CSR:
		container := define.HEADLESS_SERVER
		if fn.IsVaildHeadlessAddr(config.Headless) {
			container = config.Headless
		}
		proxy := define.PROXY_SERVER
		if fn.IsVaildProxyAddr(config.Proxy) {
			proxy = config.Proxy
		}
		return parser.GetDataAndConfigByCSR(config, container, proxy)
	case define.PARSE_MODE_MIX:
		container := define.HEADLESS_SERVER
		if fn.IsVaildHeadlessAddr(config.Headless) {
			container = config.Headless
		}
		proxy := define.PROXY_SERVER
		if fn.IsVaildProxyAddr(config.Proxy) {
			proxy = config.Proxy
		}
		return parser.GetDataAndConfigByMix(config, container, proxy)
	case define.PARSE_MODE_SSR_HQY:
		return parser.GetDataAndConfigBySSRHQY(config)
	}

	// TODO handle remote mode(api) ...
	return result
}
