package getfrompass

import (
	"testing"
)

func TestGetMissingKey(t *testing.T) {
	_, err := GetFromPass("missing_key")

	if _, ok := err.(PassExecNotFoundError); ok {
		t.Fatal("Please install pass first")
	} else if _, ok := err.(KeyNotInStoreError); !ok {
		t.Errorf("Expected KeyNotInStoreError from function, got %T (%s)", err, err)
	}
}

func TestGetTestPassword(t *testing.T) {
	pass, err := GetFromPass("test-password")

	if err != nil {
		t.Error(err)
	}

	if pass != "testPassword" {
		t.Errorf("Expected password to be 'testPassword', got '%s'", pass)
	}
}
