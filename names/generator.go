package names

import "github.com/ilmaruk/cooltures"

type Generator interface {
	FirstName(g Gender) string
	LastName() string
}

func getGenerator(c Culture, r cooltures.Randomiser) Generator {
	return simpleGenerator{rand: r}
}

type simpleGenerator struct {
	rand cooltures.Randomiser
}

func (s simpleGenerator) FirstName(g Gender) string {
	return "first"
}

func (s simpleGenerator) LastName() string {
	return "last"
}
