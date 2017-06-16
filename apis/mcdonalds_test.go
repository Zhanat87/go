package apis

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// MONGODB_DSN=localhost:27017 go test github.com/Zhanat87/go/apis
func TestMcdonaldsItems(t *testing.T) {
	res := getMcdonaldsItems()

	assert.Equal(t, 90, len(res))
}
