package directory

type ServiceDirectory interface {
	Query() error
	SendIndex() error
}

type ServiceDefaultDir struct {
	repository DirRepository
}

func NewServiceClient(repo DirRepository) *ServiceDefaultDir {
	return &ServiceDefaultDir{repository: repo}

}

func (s ServiceDefaultDir) Query() error {
	//TODO implement me
	panic("implement me")
}

func (s ServiceDefaultDir) SendIndex() error {
	//TODO implement me
	panic("implement me")
}
