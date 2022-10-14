package apparmor

import (
	"github.com/go-logr/logr"
)

type appArmorOpts struct {
	logger    logr.Logger
	policyDir string
}

type AppArmorOption func(*appArmorOpts)

func (a *appArmorOpts) applyOpts(opts ...AppArmorOption) {
	for _, o := range opts {
		o(a)
	}
}

// WithLogger sets a logger to be using whilst executing operations.
func WithLogger(logger logr.Logger) AppArmorOption {
	return func(o *appArmorOpts) {
		o.logger = logger
	}
}

// WithPolicyDir sets the directory to be used to manage policies.
func WithPolicyDir(dir string) AppArmorOption {
	return func(o *appArmorOpts) {
		o.policyDir = dir
	}
}
