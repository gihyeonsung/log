package application

type AuthnService interface {
	Login(key string) (bool, error)
}

type EnvVarAuthnService struct {
	envVar string
}

func (s *EnvVarAuthnService) Login(key string) (bool, error) {
	return s.envVar == key, nil
}
