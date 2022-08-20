// Package main is an example of how go-apparmor can be used to
// manage AppArmor profiles.
package main

import (
	"fmt"
	"os"

	"github.com/bombsimon/logrusr/v2"
	"github.com/pjbgf/go-apparmor/pkg/apparmor"
	"github.com/pjbgf/go-apparmor/pkg/hostop"
	"github.com/sirupsen/logrus"
)

const (
	validPolicyName    = "go-apparmor-testprofile"
	validPolicyContent = `#include <tunables/global>
profile go-apparmor-testprofile flags=(attach_disconnected, mediate_deleted) {
  #include <abstractions/base>

  deny network,
}
`
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)
	logger := logrusr.New(log,
		logrusr.WithReportCaller(),
	).WithCallDepth(0)

	logger.Info("Start E2E apparmor test")

	calls := func() error {
		f, err := os.CreateTemp("", "test-policy-*.aa")
		if err != nil {
			return fmt.Errorf("failed to create policy file: %v", err)
		}
		defer os.Remove(f.Name())

		_, err = f.WriteString(validPolicyContent)
		if err != nil {
			return fmt.Errorf("failed to write policy file: %v", err)
		}

		aa := apparmor.NewAppArmor().WithLogger(logger)
		enabled, err := aa.Enabled()
		if err != nil {
			return fmt.Errorf("failed to check AppArmor status: %v", err)
		}
		if !enabled {
			fmt.Printf("AppArmor is not enabled\n")
		}

		logger.Info("loading policy", "policy-name", validPolicyName)
		if err := aa.LoadPolicy(f.Name()); err != nil {
			return fmt.Errorf("load policy: %w", err)
		}

		loaded, err := aa.PolicyLoaded(validPolicyName)
		if err != nil {
			return fmt.Errorf("failed to load AppArmor policy: %v", err)
		}
		logger.Info("policy status", "policy-name", validPolicyName, "loaded", loaded)

		logger.Info("deleting policy", "policy-name", validPolicyName)
		if err := aa.DeletePolicy(validPolicyName); err != nil {
			return fmt.Errorf("failed to delete policy: %w", err)
		}

		loaded, err = aa.PolicyLoaded(validPolicyName)
		if err != nil {
			return fmt.Errorf("failed to load AppArmor policy: %v", err)
		}
		logger.Info("policy status", "policy-name", validPolicyName, "loaded", loaded)
		if loaded {
			return fmt.Errorf("policy was not removed", "policy-name", validPolicyName)
		}

		return nil
	}

	mount := hostop.NewMountHostOp().WithLogger(logger)
	if err := mount.Do(calls); err != nil {
		logger.Error(err, "E2E apparmor test failed")
		os.Exit(1)
	}
}
