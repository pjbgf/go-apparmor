package hostop

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// containerIdentifierFiles defines files that identifies the
// current context is inside a container.
var containerIdentifierFiles = []string{
	"/.dockerenv",
	"/run/.containerenv",
}

// hostIdentifierExec defines executables that are executed on pid 1
// when running at host's pid namespace.
var hostIdentifierExec = []string{
	"init",
	"systemd",
}

// InsideContainer checks whether the current process is being executed
// inside of a container.
func InsideContainer() bool {
	for _, file := range containerIdentifierFiles {
		if _, err := os.Stat(file); err == nil {
			return true
		}
	}
	return false
}

// hostPidNamespace checks whether the current process is using
// host's PID namespace.
func (m *mountHostOp) hostPidNamespace() (bool, error) {
	file, err := os.Open("/proc/1/sched")
	if err != nil {
		return false, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			m.logger.V(1).Info(fmt.Sprintf("closing file: %s", err))
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	if scanner.Scan() {
		line := scanner.Text()
		exec := strings.Split(line, " ")[0]
		for _, hostExec := range hostIdentifierExec {
			if exec == hostExec {
				return true, nil
			}
		}
	}
	return false, nil
}
