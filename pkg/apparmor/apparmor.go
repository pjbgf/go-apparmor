package apparmor

import "github.com/go-logr/logr"

type AppArmor struct {
	logger *logr.Logger
}

func NewAppArmor() *AppArmor {
	return &AppArmor{}
}

func (a *AppArmor) WithLogger(logger *logr.Logger) *AppArmor {
	a.logger = logger
	return a
}
