package bot

import (
	"testing"
)

var (
	cfgTest = &WeChatBotConfig{
		Default: &WebChatBotConfigItem{
			HookAddr:     defaultHook,
			AtMobileList: []string{"jack", "jim", "13012345678"},
			AtNameList:   []string{defaultAtName, "jack ma", "麻花腾"},
			Switch:       true,
		},
		Debug: &WebChatBotConfigItem{
			HookAddr:     defaultHook,
			AtMobileList: nil,
			Switch:       true,
		},
		Info: &WebChatBotConfigItem{
			HookAddr:     defaultHook,
			AtMobileList: nil,
			Switch:       true,
		},
		Warn: &WebChatBotConfigItem{
			HookAddr:     defaultHook,
			AtMobileList: nil,
			Switch:       true,
		},
		Error: &WebChatBotConfigItem{
			HookAddr:     defaultHook,
			AtMobileList: nil,
			Switch:       true,
		},
		Task: &WebChatBotConfigItem{
			HookAddr:     defaultHook,
			AtMobileList: nil,
			Switch:       true,
		},
	}
)

func TestWechatBot_Alarm(t *testing.T) {
	tpl := `
hello from golang
<font color="info">绿色</font>
<font color="comment">灰色</font>
<font color="warning">橙红色</font>
**加粗**
# 标题一
## 标题二
### 标题三
[这是一个链接](http://work.weixin.qq.com/api/doc)
> 引用文字
`

	msg := struct {
		Mid string
		Age int64
		Num float64
	}{
		"2333",
		23,
		23.23,
	}

	t.Log(tpl, msg)
	bot := NewWeChatBot(cfgTest)
	bot.Debug(true, "hello from golang by %v", "jack")
	bot.Error(false, tpl)
	bot.Error(true, "")
	bot.Error(false, "> 消息同步: \npayMid=%v,\n upMid=%v,\n elecNum=%v,\n msg=%+v", "123", "456", 20, msg)
	bot.Default(true, "hello")
}

func TestMessageMeta_FromText(t *testing.T) {
	msg := new(MessageMeta)
	msg.FromText("hello", nil, nil)
	t.Log(msg.Text.MentionedList, msg.Text.MentionedMobileList)
}
