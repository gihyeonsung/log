package application

type AuthnService interface {
	Login(key string) (bool, error)
}
