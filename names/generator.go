package names

import "github.com/ilmaruk/cooltures"

type Generator interface {
	FirstName(g Gender) string
	LastName() string
}

const (
	CultureEnglish Culture = "en"
)

func getGenerator(c Culture, r cooltures.Randomiser) Generator {
	switch c {
	case CultureEnglish:
		return simpleGenerator{
			firstNames: engFirstNames,
			lastNames:  engLastNames,
			rand:       r,
		}
	default:
		return simpleGenerator{
			firstNames: map[Gender][]string{
				GenderFemale: {"Jane"},
				GenderMale:   {"Joe"},
			},
			lastNames: []string{"Bloggs"},
			rand:      r,
		}
	}
}

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
