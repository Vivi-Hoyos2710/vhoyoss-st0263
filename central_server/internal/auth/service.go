package auth

import (
	"math/rand"
	"errors"
	"fmt"
)
var (
	ErrNoPeersAvailable = errors.New("no peers available to assign")
)
type Service interface {
	Login(authUser Peer) (string, error)
	Logout(authUser PeerLogOut) error
	GetUser(username string) (Peer, error)
	AssignPeer(username string, file string) (location string, err error)
	SelectRandomPeer(users []string) (Peer, error)
}

type ServiceClient struct {
	repository RepositoryAuth
	peerCount int
}

func NewServiceClient(repo RepositoryAuth) *ServiceClient {
	return &ServiceClient{repository: repo}

}
func (s *ServiceClient) Login(authUser Peer) (string, error) {

	return s.repository.SavePeer(authUser)
}

func (s *ServiceClient) Logout(authUser PeerLogOut) error {
	currentUser, err := s.repository.GetPeer(authUser.Username)
	if err != nil {
		return err
	}
	currentUser.State = "down"
	err = s.repository.UpdatePeer(currentUser)
	return err
}

func (s *ServiceClient) GetUser(username string) (Peer, error) {
	return s.repository.GetPeer(username)
}
func (s *ServiceClient) AssignPeer(username string,file string) (location string, err error){
	excludedPeerr,_:=s.repository.GetPeer(username)
	list:= s.repository.PeerOrderList(excludedPeerr)
	if len(list) == 0 {
        return location,ErrNoPeersAvailable
    }
	c:=s.peerCount%len(list)
	candidatePeerUsername:= list[c]
	peer,_:=s.repository.GetPeer(candidatePeerUsername)
	
	s.peerCount++
	fmt.Println("peerCount: ", s.peerCount,peer.Username)
	location=peer.UserURL
	err=nil
	return 
	
	
}
func (s *ServiceClient) SelectRandomPeer(users []string) (Peer, error) {
	if len(users) == 0 {
		return Peer{},ErrNoPeersAvailable
	}
	// Generate a random index within the slice boundaries
	randomIndex := rand.Intn(len(users))
  
	selectedUser := users[randomIndex]
	
	peer, err := s.GetUser(selectedUser)
	if err != nil {
	  return Peer{}, err
	}
	
	return peer, nil
	
	
	
}
