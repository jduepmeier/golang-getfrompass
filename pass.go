// Package getfrompass provides a helper method to get passwords
// from the passwordmanager pass (https://www.passwordstore.org).
package getfrompass

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var (
	ErrEmptyPassword    error = errors.New("no password found")
	ErrPassExecNotFound error = errors.New("could not find pass executable")
)

// KeyNotInStoreError is an Error that is returned if the key is not in
// the password store.
type KeyNotInStoreError struct {
	key string
}

// Error returns the message which key is not in the password store.
func (err KeyNotInStoreError) Error() string {
	return fmt.Sprintf("%s is not in the password store", err.key)
}

// PassExitError is an Error that is returned if the pass command has exited
// with an error.
type PassExitError struct {
	Message string
	Err     error
}

// Error retuns the message that pass has exited with an error.
func (err PassExitError) Error() string {
	return fmt.Sprintf("pass has exited with the following message: %s", err.Message)
}

// newKeyNotInStoreError is a helper method to create a KeyNotInStoreError
func newKeyNotInStoreError(key string) error {
	return KeyNotInStoreError{
		key: key,
	}
}

func newPassExitError(err *exec.ExitError) error {
	return PassExitError{
		Message: string(err.Stderr),
		Err:     err,
	}
}

// GetFromPass returns the password from the
// command 'pass' (see https://www.passwordstore.org/).
//
// To use this method the command 'pass' must be installed
// on the host. If the command is not installed the method
// will return the error ErrPassExecNotFound.
//
// If the given key is not in the password store an
// KeyNotInStoreError is returned.
//
// If the pass command has exited with an other error an
// PassExitError is returned.
func GetFromPass(key string) (string, error) {
	var pass string

	// check if the pass executable is in path.
	path, err := exec.LookPath("pass")
	if err != nil {
		return pass, ErrPassExecNotFound
	}

	// the command to execute is 'pass show <key>'
	out, err := exec.Command(path, "show", key).Output()
	if err != nil {
		// if the command fails the key is not in the password store.
		if exitError, ok := err.(*exec.ExitError); ok {
			switch exitError.ExitCode() {
			case 1:
				return pass, newKeyNotInStoreError(key)
			default:
				return pass, newPassExitError(exitError)
			}
		}
	}

	splits := strings.Split(string(out), "\n")
	if len(splits) > 0 {
		pass = strings.TrimRight(splits[0], "\r")
	} else {
		return pass, ErrEmptyPassword
	}

	return pass, nil
}
