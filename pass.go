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

// KeyNotInStoreError is a Error that is returned if the key is not in
// the password store.
type KeyNotInStoreError struct {
	key string
}

// Error returns the message which key is not in the password store.
func (err KeyNotInStoreError) Error() string {
	return fmt.Sprintf("%s is not in the password store", err.key)
}

// newKeyNotInStoreError is a helper method to create a KeyNotInStoreError
func newKeyNotInStoreError(key string) error {
	return KeyNotInStoreError{
		key: key,
	}
}

// GetFromPass returns the password from the
// command 'pass' (see https://www.passwordstore.org/).
//
// To use this method the command 'pass' must be installed
// on the host. If the command is not installed the method
// will return the error ErrPassExecNotFound.
//
// If the given key is not in the password store a
// KeyNotInStoreError is returned.
func GetFromPass(key string) (string, error) {
	var pass string

	path, err := exec.LookPath("pass")
	if err != nil {
		return pass, ErrPassExecNotFound
	}

	out, err := exec.Command(path, "show", key).Output()
	if err != nil {
		return pass, newKeyNotInStoreError(key)
	}

	splits := strings.Split(string(out), "\n")
	if len(splits) > 0 {
		pass = strings.TrimRight(splits[0], "\r")
	} else {
		return pass, ErrEmptyPassword
	}

	return pass, nil
}
