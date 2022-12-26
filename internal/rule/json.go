package rule

import (
	"encoding/json"
	"strings"

	"github.com/soulteary/RSS-Can/internal/define"
	"github.com/soulteary/RSS-Can/internal/fn"
	"github.com/soulteary/RSS-Can/internal/parser"
)

func ParseConfigFromJSON(str string, ruleFile string) (define.JavaScriptConfig, error) {
	var config define.JavaScriptConfig
	err := json.Unmarshal([]byte(str), &config)
	if err != nil {
		return config, err
	}

	if ruleFile != "" {
		config.File = ruleFile
	}

	modeInRule := strings.ToLower(config.Mode)
	if !fn.IsStrInArray([]string{define.PARSE_MODE_SSR, define.PARSE_MODE_CSR, define.PARSE_MODE_MIX}, modeInRule) {
		config.Mode = define.DEFAULT_PARSE_MODE
	}
	return config, nil
}

func GetWebsiteDataWithConfig(config define.JavaScriptConfig) (result define.BodyParsed) {
	switch strings.ToLower(config.Mode) {
	case define.PARSE_MODE_SSR:
		return parser.GetDataAndConfigBySSR(config)
	case define.PARSE_MODE_CSR:
		// TODO check headless addr is valid
		// TODO set proxy by config
		container := define.HEADLESS_SERVER
		proxy := define.PROXY_SERVER
		if config.Headless != "" && strings.Contains(config.Headless, ":") {
			container = config.Headless
		}
		return parser.GetDataAndConfigByCSR(config, container, proxy)
	case define.PARSE_MODE_MIX:
		// TODO check headless addr is valid
		// TODO set proxy by config
		container := define.HEADLESS_SERVER
		proxy := define.PROXY_SERVER
		if config.Headless != "" && strings.Contains(config.Headless, ":") {
			container = config.Headless
		}
		return parser.GetDataAndConfigByMix(config, container, proxy)
	}

	// TODO handle remote mode(api) ...
	return result
}
