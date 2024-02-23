package web

type AuthLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	MyURL    string `json:"my_url"`
}

type AuthLoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

func makeURL() {

}
