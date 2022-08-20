//go:build linux && apparmor
// +build linux,apparmor

package apparmor

import (
	"reflect"
	"runtime"
	"testing"
)

func TestNewAppArmor(t *testing.T) {
	previousFunc := goOS
	defer func() {
		goOS = previousFunc
	}()

	expectedCurrent := func() aa {
		if runtime.GOOS == "linux" {
			return &AppArmor{}
		}
		return &unsupported{}
	}()

	tests := []struct {
		name   string
		os     string
		wanted aa
	}{
		{
			// run this test case first, as following runs
			// will change the default value of goOS
			name:   "current OS",
			wanted: expectedCurrent,
		},
		{
			name:   "linux",
			os:     "linux",
			wanted: &AppArmor{},
		},
		{
			name:   "darwin",
			os:     "darwin",
			wanted: &unsupported{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.os != "" {
				goOS = func() string { return tt.os }
			}

			got := NewAppArmor()
			if reflect.TypeOf(got) != reflect.TypeOf(tt.wanted) {
				t.Errorf("wanted %T got %T", tt.wanted, got)
			}
		})
	}
}
