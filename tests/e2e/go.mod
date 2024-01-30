module github.com/pjbgf/go-apparmor/tests/e2e

go 1.21

replace github.com/pjbgf/go-apparmor => ../..

require (
	github.com/bombsimon/logrusr/v2 v2.0.1
	github.com/pjbgf/go-apparmor v0.1.1
	github.com/sirupsen/logrus v1.9.0
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	golang.org/x/sys v0.16.0 // indirect
)
