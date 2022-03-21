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
		return englishGenerator{rand: r}
	default:
		return simpleGenerator{rand: r}
	}
}

type simpleGenerator struct {
	rand cooltures.Randomiser
}

func (s simpleGenerator) FirstName(g Gender) string {
	switch g {
	case GenderFemale:
		return "Jane"
	case GenderMale:
		return "Joe"
	default:
		names := []string{"Jane", "Joe"}
		return names[s.rand.Intn(len(names))]
	}
}

func (s simpleGenerator) LastName() string {
	return "Bloggs"
}
