package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound        = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)

// RepositoryAuth is a repository Interface
type RepositoryAuth interface {
	SavePeer(user Peer) error
	GetAll() (*map[string]Peer, error)
	UpdatePeer(user Peer) error
	GetPeer(username string) (Peer, error)
	AllAvailablePeers(excludedPeerr string) []string
	QueryAvailableList(excludedPeerr string, list []string) []string
}
type defaultMapRepo struct {
	peerRegisterTable *map[string]Peer
}

func NewDefaultRepo(newPeerRegisterTable map[string]Peer) RepositoryAuth {
	return &defaultMapRepo{
		peerRegisterTable: &newPeerRegisterTable}
}

func (d *defaultMapRepo) SavePeer(user Peer) error {
	table := *d.peerRegisterTable
	table[user.Username] = user

	return nil

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
func (d *defaultMapRepo) AllAvailablePeers(excludedPeerr string) []string {
	orderList := make([]string, 0)
	for _, v := range *d.peerRegisterTable {
		if v.State == "up" && v.Username != excludedPeerr {
			orderList = append(orderList, v.Username)
		}

	}
	return orderList
}
func (d *defaultMapRepo) QueryAvailableList(exludedPeerUser string, list []string) []string {
	availableList := make([]string, 0)

	for _, v := range list {
		if v != exludedPeerUser {
			peer, err := d.GetPeer(v)
			if err == nil && peer.State == "up" {
				availableList = append(availableList, v)
			}
		}
	}

	return availableList
}
func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
func passwordIsValid(hashedPassword, enteredPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
	return err == nil
}
