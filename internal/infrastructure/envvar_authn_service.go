package infrastructure

import "github.com/gihyeonsung/log/internal/application"

type EnvVarAuthnService struct {
	envVar string
}

var _ application.AuthnService = (*EnvVarAuthnService)(nil)

func NewEnvVarAuthnService(envVar string) *EnvVarAuthnService {
	return &EnvVarAuthnService{envVar: envVar}
}

func (s *EnvVarAuthnService) Login(key string) (bool, error) {
	return s.envVar == key, nil
}
