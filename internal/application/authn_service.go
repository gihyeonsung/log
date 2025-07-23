package application

type AuthnService interface {
	Login(key string) (bool, error)
}

type EnvVarAuthnService struct {
	envVar string
}

var _ AuthnService = (*EnvVarAuthnService)(nil)

func NewEnvVarAuthnService(envVar string) *EnvVarAuthnService {
	return &EnvVarAuthnService{envVar: envVar}
}

func (s *EnvVarAuthnService) Login(key string) (bool, error) {
	return s.envVar == key, nil
}
