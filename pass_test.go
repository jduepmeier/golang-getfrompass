package getfrompass

import (
	"errors"
	"os/exec"
	"testing"
)

func TestMissingPass(t *testing.T) {
	lookupCommand = func(file string) (string, error) {
		return "", exec.ErrNotFound
	}
	defer func() { lookupCommand = exec.LookPath }()

	_, err := GetFromPass("test-key")
	if !errors.As(err, &PassExecNotFoundError{}) {
		t.Errorf("GetFromPass should have returned PassExecNotFoundError. Got %s instead.", err)
	}
}

func testLookup(file string) (string, error) {
	return "/bin/pass", nil
}

func TestGetMissingKey(t *testing.T) {
	lookupCommand = testLookup
	defer func() { lookupCommand = exec.LookPath }()
	execCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("/bin/false")
	}
	defer func() { execCommand = exec.Command }()
	_, err := GetFromPass("missing_key")

	if _, ok := err.(PassExecNotFoundError); ok {
		t.Errorf("GetFromPass returned PassExecNotFoundError instead of nil")
	} else if _, ok := err.(KeyNotInStoreError); !ok {
		t.Errorf("Expected KeyNotInStoreError from function, got %T (%s)", err, err)
	}
}

func TestGetTestPassword(t *testing.T) {
	lookupCommand = testLookup
	defer func() { lookupCommand = exec.LookPath }()
	execCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("echo", "testPassword")
	}
	defer func() { execCommand = exec.Command }()
	pass, err := GetFromPass("test-password")

	if err != nil {
		t.Error(err)
	}

	if pass != "testPassword" {
		t.Errorf("Expected password to be 'testPassword', got '%s'", pass)
	}
}

func TestGetEmptyPassword(t *testing.T) {
	lookupCommand = testLookup
	defer func() { lookupCommand = exec.LookPath }()
	execCommand = func(name string, arg ...string) *exec.Cmd {
		return exec.Command("/bin/true")
	}
	defer func() { execCommand = exec.Command }()
	_, err := GetFromPass("test-password")

	if !errors.As(err, &PassIsEmptyError{}) {
		t.Errorf("GetFromPass should raise PassIsEmptyError. Got %s instead.", err)
	}
}
