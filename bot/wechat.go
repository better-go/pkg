package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const (
	defaultAtAll    = "@all"
	defaultAtName   = "jack"
	defaultHook     = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=a-real-wechat-key"
	httpContentType = "application/json"
	mdCodeTpl       = `{{.Lang}}
{{.Code}}
`
	mdCodeLineTpl = `{{.Code}}`
)

type WeChatBot struct {
	conf *WeChatBotConfig
}

func NewWeChatBot(conf *WeChatBotConfig) (bot *WeChatBot) {
	cfgItem := &WebChatBotConfigItem{
		HookAddr:     defaultHook,
		AtMobileList: []string{},
		AtNameList: []string{
			defaultAtName,
		},
		Switch: false, // keep bot silent
	}

	if conf == nil {
		bot = &WeChatBot{
			conf: &WeChatBotConfig{
				Default: cfgItem,
				Debug:   cfgItem,
				Info:    cfgItem,
				Warn:    cfgItem,
				Error:   cfgItem,
				Task:    cfgItem,
			},
		}
		return
	}

	// not nil:
	bot = &WeChatBot{
		conf: conf,
	}

	if bot.conf.Default == nil {
		bot.conf.Default = cfgItem
	}
	if bot.conf.Debug == nil {
		bot.conf.Debug = cfgItem
	}
	if bot.conf.Info == nil {
		bot.conf.Info = cfgItem
	}
	if bot.conf.Warn == nil {
		bot.conf.Warn = cfgItem
	}
	if bot.conf.Error == nil {
		bot.conf.Error = cfgItem
	}
	if bot.conf.Task == nil {
		bot.conf.Task = cfgItem
	}
	return
}

type WeChatBotConfig struct {
	Default *WebChatBotConfigItem
	Debug   *WebChatBotConfigItem
	Info    *WebChatBotConfigItem
	Warn    *WebChatBotConfigItem
	Error   *WebChatBotConfigItem // error bot
	Task    *WebChatBotConfigItem // 定时任务 bot
}

type WebChatBotConfigItem struct {
	HookAddr     string
	AtNameList   []string // @用户名列表
	AtMobileList []string // @指定用户手机号列表
	Switch       bool     // 启用开关
}

// 消息:
type MessageMeta struct {
	MsgType  string           `json:"msgtype"`
	Text     *TextContent     `json:"text,omitempty"`
	Markdown *MarkdownContent `json:"markdown,omitempty"`
	Image    *ImageContent    `json:"image,omitempty"`
	News     *NewsContent     `json:"news,omitempty"`
}

// FromText 文本类型
func (m *MessageMeta) FromText(content string, atMobileList []string, atNameList []string) {
	if len(atMobileList) == 0 && len(atNameList) == 0 {
		atMobileList = append(atMobileList, defaultAtAll)
	}
	m.MsgType = "text"
	m.Text = &TextContent{
		Content:             content,
		MentionedMobileList: atMobileList,
		MentionedList:       atNameList,
	}
	return
}

// FromMarkdown markdown类型
func (m *MessageMeta) FromMarkdown(content string) {
	m.MsgType = "markdown"
	m.Markdown = &MarkdownContent{
		Content: content,
	}
	return
}

func (m *MessageMeta) ToBytes() (payload []byte) {
	payload, _ = json.Marshal(m)
	return
}

// text 消息内容:
type TextContent struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`        // at 用户名
	MentionedMobileList []string `json:"mentioned_mobile_list"` // 手机号
}

// markdown:
type MarkdownContent struct {
	Content string `json:"content"`
}

// 图片:
type ImageContent struct {
	Base64 string `json:"base64"`
	MD5    string `json:"md5"`
}

// 新闻:
type NewsContent struct {
	Articles []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		PicURL      string `json:"picurl"`
	} `json:"articles"`
}

// 代码类型:
type CodeContent struct {
	Lang string
	Code string
}

func (c *CodeContent) ToString(lang string, code string, templateFmt string) (html string) {
	buf := new(bytes.Buffer)
	cc := &CodeContent{
		Lang: lang,
		Code: code,
	}

	tf := templateFmt
	if tf == "" {
		tf = mdCodeLineTpl
	}

	tpl := template.New("mdCode")
	tpl, _ = tpl.Parse(tf)
	_ = tpl.Execute(buf, cc)
	html = buf.String()
	return
}

// Info 普通
func (b *WeChatBot) Info(atMe bool, format string, args ...interface{}) (err error) {
	return b.Alarm(b.conf.Warn, atMe, format, args...)
}

func (b *WeChatBot) Warn(atMe bool, format string, args ...interface{}) (err error) {
	return b.Alarm(b.conf.Warn, atMe, format, args...)
}

// Debug debug 级别: 后续加开关,控制 prod 不报
func (b *WeChatBot) Debug(atMe bool, format string, args ...interface{}) (err error) {
	return b.Alarm(b.conf.Debug, atMe, format, args...)
}

// Error error 级别:
func (b *WeChatBot) Error(atMe bool, format string, args ...interface{}) (err error) {
	return b.Alarm(b.conf.Error, atMe, format, args...)
}

// Task 任务类:
func (b *WeChatBot) Task(atMe bool, format string, args ...interface{}) (err error) {
	return b.Alarm(b.conf.Task, atMe, format, args...)
}

func (b *WeChatBot) Default(atMe bool, format string, args ...interface{}) (err error) {
	return b.Alarm(b.conf.Default, atMe, format, args...)
}

func (b *WeChatBot) Alarm(conf *WebChatBotConfigItem, atMe bool, format string, args ...interface{}) (err error) {
	if !conf.Switch {
		return
	}

	atMobileList := make([]string, 0, 0)
	atNameList := make([]string, 0, 0)
	if atMe {
		atMobileList = conf.AtMobileList
		atNameList = conf.AtNameList
	}
	content := fmt.Sprintf(format, args...)
	return b.alarm(conf.HookAddr, content, nil, atMobileList, atNameList)
}

func (b *WeChatBot) alarm(url string, content string, msgPack []byte, atMobileList []string, atNameList []string) (err error) {
	msg := new(MessageMeta)
	if len(atMobileList) != 0 || len(atNameList) != 0 {
		msg.FromText(content, atMobileList, atNameList)
	} else {
		msg.FromMarkdown(content)
	}

	// fmt:
	payload := msgPack
	if len(payload) == 0 {
		payload = msg.ToBytes()
	}

	// post:
	_, err = b.send(url, payload)
	return
}

func (b *WeChatBot) send(url string, payload []byte) (resp *http.Response, err error) {
	return http.Post(url, httpContentType, bytes.NewBuffer(payload))
}
