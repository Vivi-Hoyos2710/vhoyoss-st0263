package client

import "github.com/Vivi-Hoyos2710/vhoyoss-st0263/P1/config"

type ServiceClient struct {
	BaseURL string
}

func NewServiceClient(configs config.Config) ServiceInterface {
	return &ServiceClient{BaseURL: configs.IPCentralServer + ":" + configs.PortCentralServer}

}

func (c *ServiceClient) Login() error {
	//TODO implement me
	panic("implement me")
}

func (c *ServiceClient) Logout() error {
	//TODO implement me
	panic("implement me")
}

func (c *ServiceClient) Query() error {
	//TODO implement me
	panic("implement me")
}

func (c *ServiceClient) SendIndex() error {
	//TODO implement me
	panic("implement me")
}
