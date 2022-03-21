package names

import "github.com/ilmaruk/cooltures"

type Culture string

const CultureAny Culture = ""

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

// pickGender either returns the chosen gender or it
// picks a random one, if GenderAny was chosen.
func pickGender(g Gender, r cooltures.Randomiser) Gender {
	if g == GenderAny {
		return []Gender{GenderFemale, GenderMale}[r.Intn(2)]
	}
	return g
}
