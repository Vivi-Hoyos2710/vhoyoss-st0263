package client

type ServiceInterface interface {
	Login() error
	Logout() error
	Query() error
	SendIndex() error
}
