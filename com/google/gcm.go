package google

/*
- 说明:
	Firebase Cloud Messaging(FCM) = Google Cloud Messaging(GCM)
	google 推送服务

- docs:
	- https://firebase.google.com/docs?hl=zh-cn
	- https://firebase.google.com/docs/libraries/
	- https://firebase.google.com/docs/cloud-messaging?hl=zh-cn
	- https://firebase.google.com/docs/cloud-messaging/android/client
	- etc:
		- https://studygolang.com/articles/19998
		- https://medium.com/@mikru168/ios-google-notification-firebase-cloud-message-c2849117be08

- libs:
	- https://github.com/firebase/
	- https://github.com/firebase/firebase-admin-go
		- https://godoc.org/firebase.google.com/go
	- https://github.com/googlearchive/go-gcm
		- https://github.com/kikinteractive/go-gcm
		- 官方 fork
	- https://github.com/googollee/go-gcm
	- https://github.com/SherClockHolmes/webpush-go

- app client biz steps:
		- 初次启动您的应用时，FCM SDK 会为客户端应用实例生成一个注册令牌。
		- 如果您希望指定单一目标设备或者创建设备组，则需要通过继承 FirebaseMessagingService 并重写 onNewToken 来获取此令牌。
*/

var (
	// google push:
	gcmUrl = "https://fcm.googleapis.com/fcm/send"
)

// google push 服务:
type Push struct {
}
