package names_test

import (
	"fmt"
	"github.com/ilmaruk/cooltures/names"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullName(t *testing.T) {
	tt := []struct {
		culture  names.Culture
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
		{
			culture:  names.CultureEnglish,
			gender:   names.GenderFemale,
			expected: "Evelyn Turpin",
		},
		{
			culture:  names.CultureEnglish,
			gender:   names.GenderMale,
			expected: "Aiden Turpin",
		},
	}

	n := names.New()
	for _, tc := range tt {
		t.Run(fmt.Sprintf("Culture %s / Gender %s", tc.culture, tc.gender), func(t *testing.T) {
			o := names.Options{
				Culture: tc.culture,
				Gender:  tc.gender,
			}
			assert.Equal(t, tc.expected, n.FullName(o))
		})
	}
}
