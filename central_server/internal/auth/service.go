package auth

type Service interface {
	Login(authUser Peer) (string, error)
	Logout(authUser PeerLogOut) error
	GetUser(username string) (Peer, error)
}

type ServiceClient struct {
	repository RepositoryAuth
}

func NewServiceClient(repo RepositoryAuth) *ServiceClient {
	return &ServiceClient{repository: repo}

}
func (s ServiceClient) Login(authUser Peer) (string, error) {

	return s.repository.SavePeer(authUser)
}

func (s ServiceClient) Logout(authUser PeerLogOut) error {
	currentUser, err := s.repository.GetUser(authUser.Username)
	if err != nil {
		return err
	}
	currentUser.State = "down"
	err = s.repository.UpdatePeer(currentUser)
	return err
}

func (s ServiceClient) GetUser(username string) (Peer, error) {
	return s.repository.GetUser(username)
}
