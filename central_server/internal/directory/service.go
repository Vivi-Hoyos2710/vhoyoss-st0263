package directory

type ServiceDirectory interface {
	Query() error
	SendIndex() error
}

type ServiceDefaultDir struct {
	repository Repository
}

func NewServiceClient(repo Repository) *ServiceClient {
	return &ServiceClient{repository: repo}

}

func (s ServiceClient) Query() error {
	//TODO implement me
	panic("implement me")
}

func (s ServiceClient) SendIndex() error {
	//TODO implement me
	panic("implement me")
}
