package names

import (
	"github.com/ilmaruk/cooltures"
)

type Generator interface {
	FirstName(g Gender) string
	LastName() string
	SetRandomiser(r cooltures.Randomiser)
}

const (
	CultureEnglish Culture = "en"
	CultureFrench  Culture = "fr"
	CultureGerman  Culture = "de"
	CultureItalian Culture = "it"
	CulturePolish  Culture = "pl"
	CultureSpanish Culture = "es"
)

var cultureGenerators = map[Culture]Generator{
	CultureEnglish: &simpleGenerator{
		firstNames: engFirstNames,
		lastNames:  engLastNames,
	},
	CultureFrench: &simpleGenerator{
		firstNames: fraFirstNames,
		lastNames:  fraLastNames,
	},
	CultureGerman: &simpleGenerator{
		firstNames: gerFirstNames,
		lastNames:  gerLastNames,
	},
	CultureItalian: &simpleGenerator{
		firstNames: itaFirstNames,
		lastNames:  itaLastNames,
	},
	CulturePolish: &simpleGenerator{
		firstNames: polFirstNames,
		lastNames:  polLastNames,
	},
	CultureSpanish: &spanishGenerator{},
}

func getGenerator(c Culture, r cooltures.Randomiser) (Culture, Generator) {
	c = pickCulture(c, r)
	gen := cultureGenerators[c]
	gen.SetRandomiser(r)
	return c, gen
}

// simpleGenerator generates single first and last names from
// given slices of names. Double first and last names should be
// included in the provided slices.
type simpleGenerator struct {
	firstNames map[Gender][]string
	lastNames  []string
	rand       cooltures.Randomiser
}

func (s simpleGenerator) FullName(g Gender) string {
	return s.FirstName(g) + " " + s.LastName()
}

func (s simpleGenerator) FirstName(g Gender) string {
	return pickName(s.firstNames[g], s.rand)
}

func (s simpleGenerator) LastName() string {
	return pickName(s.lastNames, s.rand)
}

func (s *simpleGenerator) SetRandomiser(r cooltures.Randomiser) {
	s.rand = r
}
