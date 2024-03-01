package auth

import (
	"errors"
)

var (
	ErrExists   = errors.New("user is already logged")
	ErrNotFound = errors.New("user not found")
)

// RepositoryAuth is a repository Interface
type RepositoryAuth interface {
	SavePeer(user Peer) (string, error)
	GetAll() (*map[string]Peer, error)
	UpdatePeer(user Peer) error
	GetPeer(username string) (Peer, error)
	PeerOrderList (excludedPeerr Peer)([]string)

}
type defaultMapRepo struct {
	peerRegisterTable *map[string]Peer
}

func NewDefaultRepo(newPeerRegisterTable map[string]Peer) RepositoryAuth {
	return &defaultMapRepo{
		peerRegisterTable: &newPeerRegisterTable	}
}

func (d *defaultMapRepo) SavePeer(user Peer) (string, error) {
	table := *d.peerRegisterTable
	currentUser, err := d.GetPeer(user.Username)
	if err == nil && currentUser.State == "up" {
		return "dummyToken123", ErrExists
	}
	user.State = "up"
	table[user.Username] = user
	token := "dummyToken123"
	return token, nil

}

func (d *defaultMapRepo) GetAll() (*map[string]Peer, error) {
	v := make(map[string]Peer)
	table := *d.peerRegisterTable
	for _, user := range table {
		v[user.Username] = table[user.Username]

	}
	return &v, nil
}

func (d *defaultMapRepo) UpdatePeer(user Peer) error {
	table := *d.peerRegisterTable
	table[user.Username] = user
	return nil
}

func (d *defaultMapRepo) GetPeer(username string) (Peer, error) {
	table := *d.peerRegisterTable
	if _, ok := table[username]; ok {
		return table[username], nil
	}
	return Peer{}, ErrNotFound
}
func (d *defaultMapRepo) PeerOrderList (excludedPeerr Peer)([]string){
	orderList:= make([]string,0)
	for _, v := range *d.peerRegisterTable {
		if v.State == "up" && v.Username != excludedPeerr.Username{
			orderList = append(orderList,v.Username)
		}
		
	}
	return orderList
}
