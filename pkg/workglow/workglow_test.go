package workglow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkGlowParse(t *testing.T) {
	assert := assert.New(t)

	_, err := ParseWorkGlow("testdata/test_workglow.yaml")
	assert.Nil(err)
}
