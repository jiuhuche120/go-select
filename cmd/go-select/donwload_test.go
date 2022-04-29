package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvailableVersion(t *testing.T) {
	versions, err := AvailableVersion()
	fmt.Println(versions)
	assert.Nil(t, err)
}
