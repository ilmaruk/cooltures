package names_test

import (
	"github.com/ilmaruk/cooltures/names"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullName(t *testing.T) {
	n := names.New()
	o := names.Options{}
	assert.Equal(t, "first last", n.FullName(o))
}
