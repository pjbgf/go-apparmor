// Package main is an example of how go-apparmor can be used to
// manage AppArmor profiles.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bombsimon/logrusr/v2"
	"github.com/pjbgf/go-apparmor/pkg/apparmor"
	"github.com/pjbgf/go-apparmor/pkg/hostop"
	"github.com/sirupsen/logrus"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s /some-root-folder/profile-path\n", os.Args[0])
		os.Exit(-1)
	}

	profilePath := os.Args[1]
	profileName := strings.Split(filepath.Base(profilePath), ".")[0]

	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)
	logger := logrusr.New(log,
		logrusr.WithReportCaller(),
	).WithCallDepth(0)

	calls := func() error {
		aa := apparmor.NewAppArmor(logger)

		fmt.Println("Delete Policy...")
		if err := aa.DeletePolicy(profileName); err != nil {
			fmt.Println("ERROR: delete policy: %w", err)
		}

		fmt.Println("Load Policy...")
		if err := aa.LoadPolicy(profilePath); err != nil {
			return fmt.Errorf("load policy: %w", err)
		}
		return nil
	}

	mount := hostop.NewMountHostOp(logger)
	if err := mount.Do(calls); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
