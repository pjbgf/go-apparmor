module github.com/pjbgf/go-apparmor/example/code

go 1.18

replace github.com/pjbgf/go-apparmor => ../..

require (
	github.com/bombsimon/logrusr/v2 v2.0.1
	github.com/pjbgf/go-apparmor v0.0.5
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
)
