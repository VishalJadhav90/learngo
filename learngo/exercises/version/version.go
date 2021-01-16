package version

import "runtime"

// Version function returns go version installed
func Version() string {
	return runtime.Version()
}
