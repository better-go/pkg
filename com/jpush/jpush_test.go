package jpush

import "testing"

func TestJPush_SendMessage(t *testing.T) {
	appKey := "92f0f39efe6c5a600b10b878" // test-key  // 92f0f39efe6c5a600b10b878
	appSecret := "748d6b8277a5398790ce2325" // test-secret // 748d6b8277a5398790ce2325

	// client:
	push := NewJPush(appKey, appSecret)

	// do push:
	err := push.SendMessage(push.SetPayload(
		push.SetPlatform(Android),
		push.SetAudience(false, []string{}, []string{}, []string{}, []string{}, ),
		push.SetNotice("alert test1", "android test1", "", ""),
		push.SetMessage("test title1", "test content1"),
	))

	t.Logf("push result: %v", err)
}
