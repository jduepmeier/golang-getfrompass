package getfrompass

import (
	"fmt"
	"os/exec"
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

// Error returns the message that pass has exited with an error.
func (err PassExitError) Error() string {
	return fmt.Sprintf("pass has exited with the following message: %s", err.Message)
}

// PassExecNotFoundError is an Error that is returned if the pass executable is not
// in the given path.
type PassExecNotFoundError struct {
	Path string
}

// Error returns the message that the pass executable is not in the path.
func (err PassExecNotFoundError) Error() string {
	return "could not find pass executable in path %s"
}

// newKeyNotInStoreError is a helper method to create a KeyNotInStoreError.
func newKeyNotInStoreError(key string) error {
	return KeyNotInStoreError{
		key: key,
	}
}

// newPassExitError is a helper method to return an PassExitError.
func newPassExitError(err *exec.ExitError) error {
	return PassExitError{
		Message: string(err.Stderr),
		Err:     err,
	}
}

// newPassExecNotError is a helper method to return an PassExecNotFoundError.
func newPassExecNotFoundError(path string) error {
	return PassExecNotFoundError{
		Path: path,
	}
}

// PassIsEmptyError is an Error that is returned if the password returned from pass
// is empty.
type PassIsEmptyError struct {
	key string
}

func (err PassIsEmptyError) Error() string {
	return fmt.Sprintf("password for %s is empty", err.key)
}

// newPassIsEmptyError is a helper method to return an PassIsEmptyError.
func newPassIsEmptyError(key string) error {
	return PassIsEmptyError{
		key: key,
	}
}
