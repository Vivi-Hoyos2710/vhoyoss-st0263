package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"os"
	"time"
)

var (
	ErrNoPeersAvailable = errors.New("no peers available to assign")
	secretKey           = []byte(os.Getenv("SECRET"))
)

type Service interface {
	Login(authUser Peer) (string, error)
	Logout(authUser Peer) error
	GetUser(username string) (Peer, error)
	AssignPeer(excludedPeerr Peer) (location string, err error)
	SelectRandomPeer(users []string, excludedUser Peer) (Peer, error)
}

type ServiceClient struct {
	repository  RepositoryAuth
	peerCount   int
	validTokens map[string]Peer
}

func NewServiceClient(repo RepositoryAuth) *ServiceClient {
	return &ServiceClient{repository: repo}

}
func (s *ServiceClient) Login(authUser Peer) (string, error) {

	//Check if user is already logged
	currentUser, err := s.repository.GetPeer(authUser.Username)
	if err == nil && currentUser.State == "up" {
		_, err := s.ValidatePeerToken(currentUser.TokenSesion)
		if passwordIsValid(currentUser.Password, authUser.Password) && err == nil {
			return currentUser.TokenSesion, nil
		}
		if !passwordIsValid(currentUser.Password, authUser.Password) {
			return "", ErrInvalidPassword
		}
		if err != nil {
			token, err := GenerateToken(authUser)
			if err != nil {
				return "", err
			}
			authUser.TokenSesion = token
			err = s.repository.SavePeer(authUser)
			if err != nil {
				return "", err
			}
		}
	}
	//New session
	authUser.State = "up"
	//Encrypt password
	encryptedPassword, err := encryptPassword(authUser.Password)
	if err != nil {
		return "", err
	}
	authUser.Password = encryptedPassword
	//Generate token
	token, err := GenerateToken(authUser)
	if err != nil {
		return "", err
	}
	authUser.TokenSesion = token
	err = s.repository.SavePeer(authUser)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *ServiceClient) Logout(authUser Peer) error {

	authUser.State = "down"
	authUser.TokenSesion = ""
	err := s.repository.UpdatePeer(authUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceClient) GetUser(username string) (Peer, error) {
	return s.repository.GetPeer(username)
}
func (s *ServiceClient) AssignPeer(excludedPeerr Peer) (location string, err error) {

	list := s.repository.PeerOrderList(excludedPeerr)
	if len(list) == 0 {
		return location, ErrNoPeersAvailable
	}
	c := s.peerCount % len(list)
	candidatePeerUsername := list[c]
	peer, _ := s.repository.GetPeer(candidatePeerUsername)

	s.peerCount++
	fmt.Println("Peer: ", excludedPeerr.Username, "Candidate: ", peer.Username, "into the socket: ", peer.UserURL)
	location = peer.UserURL
	err = nil
	return

}
func (s *ServiceClient) SelectRandomPeer(users []string, excludedUser Peer) (Peer, error) {
	if len(users) ==0 {
		return Peer{}, ErrNoPeersAvailable
	}
	if len(users) == 1 && users[0] == excludedUser.Username{
		return Peer{}, ErrNoPeersAvailable
	}
	// Loop until it finds a random user that is not the excluded user
	var selectedUser string
	for {
		randomIndex := rand.Intn(len(users))
		if users[randomIndex] != excludedUser.Username {
			selectedUser = users[randomIndex]
			break
		}
	}

	peer, err := s.GetUser(selectedUser)
	if err != nil {
		return Peer{}, err
	}

	return peer, nil

}
func GenerateToken(peer Peer) (string, error) {
	claims := PeerClaims{
		Peer: peer,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), //Expira en 2 horas
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (s *ServiceClient) ValidatePeerToken(tokenString string) (*PeerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &PeerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*PeerClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	user, err := s.GetUser(claims.Peer.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid token claims")
	}
	if user.State == "down" {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
