package names

import "github.com/ilmaruk/cooltures"

type englishGenerator struct {
	rand cooltures.Randomiser
}

func (e englishGenerator) FullName(g Gender) string {
	return e.FirstName(g) + " " + e.LastName()
}

func (e englishGenerator) FirstName(g Gender) string {
	g = pickGender(g, e.rand)
	return pickName(engFirstNames[g], e.rand)
}

func (e englishGenerator) LastName() string {
	return pickName(engLastNames, e.rand)
}

var engFirstNames = map[Gender][]string{
	GenderFemale: {"Evelyn"},
	GenderMale:   {"Aiden"},
}

var engLastNames = []string{"Turpin"}
