package authentication

type Auth interface {
	register()
	login()
	logout()
}
