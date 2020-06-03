package text

import (
	"github.com/microcosm-cc/bluemonday"
)

// 严格过滤模式:
func AntiXssStrict(text string) string {
	p := bluemonday.StrictPolicy() // 严格模式
	html := p.Sanitize(text)       // 消毒
	return html
}

// 非严格模式: 可部分显示 H5/JS
func AntiXssLite(text string) string {
	p := bluemonday.UGCPolicy() // 非严格模式
	html := p.Sanitize(text)    // 消毒
	return html
}
