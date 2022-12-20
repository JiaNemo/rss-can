package rule

import (
	"encoding/json"
	"strings"

	"github.com/soulteary/RSS-Can/internal/define"
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
	return ApplyDefaults(config), nil
}

func ApplyDefaults(config define.JavaScriptConfig) define.JavaScriptConfig {
	modeInRuls := strings.ToLower(config.Mode)
	if !(modeInRuls == define.PARSE_MODE_SSR || modeInRuls == define.PARSE_MODE_CSR || modeInRuls == define.PARSE_MODE_MIX) {
		config.Mode = define.DEFAULT_PARSE_MODE
	}
	return config
}

func GetWebsiteDataWithConfig(config define.JavaScriptConfig) (result define.BodyParsed) {
	switch strings.ToLower(config.Mode) {
	case define.PARSE_MODE_SSR:
		return parser.GetDataAndConfigBySSR(config)
	case define.PARSE_MODE_CSR:
		// TODO read config
		const container = define.DEFAULT_HEADLESS_CHROME
		const proxy = define.DEFAULT_PROXY_ADDRESS
		return parser.GetDataAndConfigByCSR(config, container, proxy)
	case define.PARSE_MODE_MIX:
		const container = define.DEFAULT_HEADLESS_CHROME
		const proxy = define.DEFAULT_PROXY_ADDRESS
		return parser.GetDataAndConfigByMix(config, container, proxy)
	}

	// TODO handle remote mode(api) ...
	return result
}
