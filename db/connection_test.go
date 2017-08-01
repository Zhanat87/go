package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := Connection.Ping()

	assert.NotNil(t, Connection)
	assert.Nil(t, err)
}
