package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := New("error text")

	assert.EqualError(t, err, "error text")
}

func TestNoProfileExists(t *testing.T) {
	assert.EqualError(t, NoProfileExists, "no profile exists")
}

func TestProfileNotSet(t *testing.T) {
	assert.EqualError(t, ProfileNotSet, "user profile not set")
}
