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

type KeyNotInStoreError struct {
	key string
}

func (err KeyNotInStoreError) Error() string {
	return fmt.Sprintf("%s is not in the password store", err.key)
}

func newKeyNotInStoreError(key string) error {
	return KeyNotInStoreError{
		key: key,
	}
}

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
