// Package gogitversion provides an embeddable version in an application. When
// no version is embedded it tries to find one using `git describe` and pulling
// it out of the found tags.
//
// The version gets embedded using `-ldflags` and setting `github.com/prognosai/gogitversion.version`.
// For example:
//
//   go install -ldflags "-X github.com/prognosai/gogitversion.version=0.0.1" ./...
//
// When no version is embedded git is used to try and extract the version from
// the repository tags with the command:
//
//   git describe --tags --dirty
//
// If it fails for any reason the version 'unknown' is returned.
//
package gogitversion

import (
	"os/exec"
	"strings"
)

var version string

// gitDescribeVersion trys to use `git describe` to return the version. If git
// isn't available or git describe fails, returns `unknown`
func gitDescribeVersion() string {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		return "unknown"
	}

	/* #nosec */
	out, err := exec.Command(gitPath, "describe", "--tags", "--dirty").Output()
	if err != nil {
		return "unknown"
	}

	return strings.TrimSpace(string(out))
}

// Get the version that has been set with `-ldflags "-X ..."` or try and use git describe to find it.` Returns 'unknown'
// if the version can't be determined.
func Get() string {
	if version != "" {
		return version
	}

	return gitDescribeVersion()
}
