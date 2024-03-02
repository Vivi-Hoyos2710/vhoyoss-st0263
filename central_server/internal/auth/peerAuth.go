package auth

import "github.com/golang-jwt/jwt"

type Peer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserURL  string `json:"user_url"`
	State    string `json:"-"`
	TokenSesion string `json:"-"`
}

type PeerClaims struct {
	Peer
	jwt.StandardClaims
}