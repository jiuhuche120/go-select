package main

import (
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	err := Initialize()
	assert.Nil(t, err)
	dir, err := homedir.Expand(defaultPath)
	assert.Nil(t, err)
	assert.Equal(t, true, Exist(dir))
}
