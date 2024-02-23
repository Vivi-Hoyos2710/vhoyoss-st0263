package auth

type Service interface {
	Login(authUser Peer) (string, error)
	Logout() error
}

type ServiceClient struct {
	repository RepositoryAuth
}

func NewServiceClient(repo RepositoryAuth) *ServiceClient {
	return &ServiceClient{repository: repo}

}
func (s *ServiceClient) Login(authUser Peer) (string, error) {
	return s.repository.SavePeer(authUser)
}

func (s ServiceClient) Logout() error {
	return s.repository.DeletePeer()
}
