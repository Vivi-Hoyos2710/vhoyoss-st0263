package auth

import (
	"errors"
	"fmt"
)

var (
	ErrExists = errors.New("user already exists")
)

// RepositoryAuth is a repository Interface
type RepositoryAuth interface {
	SavePeer(user Peer) (string, error)
	DeletePeer() error
	GetAll() (map[int]Peer, error)
}
type defaultMapRepo struct {
	peerRegisterTable *map[int]Peer
}

func NewDefaultRepo(newPeerRegisterTable map[int]Peer) RepositoryAuth {
	return &defaultMapRepo{
		peerRegisterTable: &newPeerRegisterTable,
	}
}

func (d defaultMapRepo) SavePeer(user Peer) (string, error) {
	table := *d.peerRegisterTable
	for key, peer := range table {
		if peer.Username == user.Username {
			return "", fmt.Errorf("%w id:%d", ErrExists, key)
		}
	}
	size := len(table)
	table[size] = user
	token := "dummyToken123"
	return token, nil
}

func (d defaultMapRepo) GetAll() (map[int]Peer, error) {
	v := make(map[int]Peer)
	table := *d.peerRegisterTable
	// copy db
	for key, value := range table {
		v[key] = value
	}
	return v, nil
}

func (d defaultMapRepo) DeletePeer() error {
	//TODO implement me
	panic("implement me")
}
