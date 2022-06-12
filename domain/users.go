package domain

type User struct {
	Login string
	Hash  []byte
}

type UserCredentials struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}
