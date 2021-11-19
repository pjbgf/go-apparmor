module github.com/pjbgf/go-apparmor/example/code

go 1.17

require (
	github.com/bombsimon/logrusr/v2 v2.0.1
	github.com/pjbgf/go-apparmor v0.0.4
	github.com/sirupsen/logrus v1.8.1
)

replace github.com/pjbgf/go-apparmor => ../..

require (
	github.com/go-logr/logr v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20211107104306-e0b2ad06fe42 // indirect
)
