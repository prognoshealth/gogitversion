package gogitversion

import (
	"testing"
)

func TestGet(t *testing.T) {
	expected := "9999999.9999999.1"

	version = expected
	defer func() { version = "" }()

	actual := Get()

	if expected != actual {
		t.Errorf("'%s' expected but got '%s'", expected, actual)
	}
}

func TestGet_GitUnknown(t *testing.T) {
	// TODO: horrible test, needs some refactoring to allow for better testing.
	expected := "unknown"
	actual := Get()

	if expected != actual {
		t.Errorf("'%s' expected but got '%s'", expected, actual)
	}
}
