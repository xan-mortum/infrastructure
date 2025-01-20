package service

type AuthClient interface {
	SignIn() ([]byte, error)
	CheckToken(string) ([]byte, error)
}

type ResourceClient interface {
	GetBooks() ([]byte, error)
	GetUsers() ([]byte, error)
}
