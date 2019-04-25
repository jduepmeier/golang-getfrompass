// Package getfrompass provides a helper method to get passwords
// from the passwordmanager pass (https://www.passwordstore.org).
package getfrompass

import (
	"os"
	"os/exec"
	"strings"
)

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
		return pass, newPassExecNotFoundError(os.Getenv("PATH"))
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
		return pass, newPassIsEmptyError(key)
	}

	return pass, nil
}
