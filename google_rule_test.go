package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleUsername(t *testing.T) {
	rule := GoogleRule{}
	assert.Equal(t, "t", rule.ProcessUsername("t+.est"))
}

func TestGoogleDomain(t *testing.T) {
	rule := GoogleRule{}
	assert.Equal(t, "gmail.com", rule.ProcessDomain("googlemail.com"))
	assert.Equal(t, "gmail.com", rule.ProcessDomain("gmail.com"))
}
