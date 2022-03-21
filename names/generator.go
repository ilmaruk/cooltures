package names

import (
	"fmt"
	"github.com/ilmaruk/cooltures"
)

type Generator interface {
	FirstName(g Gender) string
	LastName() string
	SetRandomiser(r cooltures.Randomiser)
}

const (
	CultureEnglish Culture = "en"
	CultureItalian Culture = "it"
)

var cultureGenerators = map[Culture]Generator{
	CultureEnglish: &simpleGenerator{
		firstNames: engFirstNames,
		lastNames:  engLastNames,
	},
	CultureItalian: &simpleGenerator{
		firstNames: itaFirstNames,
		lastNames:  itaLastNames,
	},
}

func getGenerator(c Culture, r cooltures.Randomiser) (Generator, error) {
	if c == CultureAny {
		var generators = make([]Generator, 0, len(cultureGenerators))
		for _, g := range cultureGenerators {
			generators = append(generators, g)
		}
		g := generators[r.Intn(len(generators))]
		g.SetRandomiser(r)
		return g, nil
	}

	g, ok := cultureGenerators[c]
	if !ok {
		return nil, fmt.Errorf("culture not supported: %s", c)
	}
	g.SetRandomiser(r)
	return g, nil
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
	g = pickGender(g, s.rand)
	return pickName(s.firstNames[g], s.rand)
}

func (s simpleGenerator) LastName() string {
	return pickName(s.lastNames, s.rand)
}

func (s *simpleGenerator) SetRandomiser(r cooltures.Randomiser) {
	s.rand = r
}
