package names

import (
	"fmt"
	"github.com/ilmaruk/cooltures"
	"math/rand"
	"time"
)

type Culture string

const (
	CultureAny     Culture = ""
	CultureEnglish Culture = "eng"
	CultureItalian Culture = "ita"
	CultureSpanish Culture = "spa"
)

type GeneratorBuilder func(r cooltures.Randomiser) Generator

// Here goes the map of supported cultures and their builders
var generatorBuilders = map[Culture]GeneratorBuilder{
	CultureEnglish: NewEnglish,
	CultureItalian: NewItalian,
	CultureSpanish: NewSpanish,
}

type Gender string

const (
	GenderAny    Gender = ""
	GenderFemale Gender = "female"
	GenderMale   Gender = "male"
)

type Options struct {
	Culture Culture
	Gender  Gender
}

// Generator is the interface to be implemented by culture specific generators.
type Generator interface {
	Name(g Gender) string
	FirstName(g Gender) string
	LastName() string
}

type Names struct {
	rand cooltures.Randomiser
}

func New(r cooltures.Randomiser) Names {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return Names{
		rand: r,
	}
}

func (g Names) Name(o Options) (name string, culture Culture, gender Gender, err error) {
	// Selects the culture
	culture, err = selectCulture(g.rand, o.Culture)
	if err != nil {
		return
	}

	// Sets the gender
	if o.Gender != GenderAny {
		gender = o.Gender
	} else {
		gender = []Gender{GenderFemale, GenderMale}[g.rand.Intn(2)]
	}

	// Sets the generator
	generatorBuilder, _ := generatorBuilders[culture]
	generator := generatorBuilder(g.rand)

	name = generator.Name(gender)
	err = nil

	return
}

func (g Names) FirstName(o Options) (name string, culture Culture, gender Gender, err error) {
	// Selects the culture
	culture, err = selectCulture(g.rand, o.Culture)
	if err != nil {
		return
	}

	// Sets the gender
	if o.Gender != GenderAny {
		gender = o.Gender
	} else {
		gender = []Gender{GenderFemale, GenderMale}[g.rand.Intn(2)]
	}

	// Sets the generator
	generatorBuilder, _ := generatorBuilders[culture]
	generator := generatorBuilder(g.rand)

	name = generator.FirstName(gender)
	err = nil

	return
}

func (g Names) LastName(o Options) (name string, culture Culture, gender Gender, err error) {
	// Selects the culture
	culture, err = selectCulture(g.rand, o.Culture)
	if err != nil {
		return
	}

	// Sets the generator
	generatorBuilder, _ := generatorBuilders[culture]
	generator := generatorBuilder(g.rand)

	name = generator.LastName()
	err = nil

	return
}

func singleName(r cooltures.Randomiser, names []string) string {
	return names[r.Intn(len(names))]
}

func doubleName(r cooltures.Randomiser, names []string, allowDuplicate bool) string {
	first := singleName(r, names)
	second := singleName(r, names)
	for first == second && !allowDuplicate {
		second = singleName(r, names)
	}

	return first + " " + second
}

func getSources(g Gender, names map[Gender][]string) []string {
	sources := make([]string, 0)
	if g == GenderAny {
		sources = append(sources, names[GenderFemale]...)
		sources = append(sources, names[GenderMale]...)
	} else {
		sources = append(sources, names[g]...)
	}
	return sources
}

func selectCulture(r cooltures.Randomiser, c Culture) (Culture, error) {
	if c != CultureAny {
		// Validates the culture (valid if there's a relevant generator)
		if _, ok := generatorBuilders[c]; !ok {
			return c, fmt.Errorf("invalid culture %s", c)
		}
		return c, nil
	} else {
		// Selects a random culture
		cultures := make([]Culture, 0, len(generatorBuilders))
		for key := range generatorBuilders {
			cultures = append(cultures, key)
		}
		return cultures[r.Intn(len(cultures))], nil
	}
}
