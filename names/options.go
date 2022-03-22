package names

import (
	"github.com/ilmaruk/cooltures"
)

type Culture string

type Gender string

const (
	GenderAny    Gender = ""
	GenderFemale Gender = "F"
	GenderMale   Gender = "M"
)

type Options struct {
	Culture Culture
	Gender  Gender
}

// pickCulture either returns the chosen culture, if it is supported,
// or it picks a random one, if CultureAny was chosen.
// If the selected culture is not supported a random one is chosen.
func pickCulture(c Culture, r cooltures.Randomiser) Culture {
	_, ok := cultureGenerators[c]
	if ok {
		// c is a supported culture
		return c
	}

	var cultures = make([]Culture, 0, len(cultureGenerators))
	for culture := range cultureGenerators {
		cultures = append(cultures, culture)
	}
	return cultures[r.Intn(len(cultures))]
}

// pickGender either returns the chosen gender, or it
// picks a random one, if GenderAny was chosen.
func pickGender(g Gender, r cooltures.Randomiser) Gender {
	if g == GenderAny {
		return []Gender{GenderFemale, GenderMale}[r.Intn(2)]
	}
	return g
}
