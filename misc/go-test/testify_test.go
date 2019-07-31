package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add(t *testing.T) {
	assert.Equal(t, add(4, 9), 10)
}
