package names_test

import (
	"fmt"
	"github.com/ilmaruk/cooltures/names"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullName(t *testing.T) {
	tt := []struct {
		gender   names.Gender
		expected string
	}{
		{
			gender:   names.GenderFemale,
			expected: "Jane Bloggs",
		},
		{
			gender:   names.GenderMale,
			expected: "Joe Bloggs",
		},
	}

	n := names.New()
	for _, tc := range tt {
		t.Run(fmt.Sprintf("Gender %s", tc.gender), func(t *testing.T) {
			o := names.Options{
				Gender: tc.gender,
			}
			assert.Equal(t, tc.expected, n.FullName(o))
		})
	}
}
