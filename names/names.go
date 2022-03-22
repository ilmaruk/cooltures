package names

import (
	"github.com/ilmaruk/cooltures"
	"math/rand"
	"time"
)

type Name struct {
	Name    string
	Culture Culture
	Gender  Gender
}

type Names struct {
	rand cooltures.Randomiser
}

func NewWithRandomiser(r cooltures.Randomiser) Names {
	return Names{
		rand: r,
	}
}

func New() Names {
	return NewWithRandomiser(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func (n Names) FullName(o Options) Name {
	// The first name also sets culture and gender
	first := n.FirstName(o)
	o.Culture = first.Culture
	o.Gender = first.Gender

	last := n.LastName(o)

	return Name{
		Name:    first.Name + " " + last.Name,
		Culture: o.Culture,
		Gender:  o.Gender,
	}
}

func (n Names) FirstName(o Options) Name {
	c, gen := getGenerator(o.Culture, n.rand)
	g := pickGender(o.Gender, n.rand)

	return Name{
		Name:    gen.FirstName(g),
		Culture: c,
		Gender:  g,
	}
}

func (n Names) LastName(o Options) Name {
	c, gen := getGenerator(o.Culture, n.rand)

	return Name{
		Name:    gen.LastName(),
		Culture: c,
		Gender:  o.Gender,
	}
}

func pickName(src []string, r cooltures.Randomiser) string {
	return src[r.Intn(len(src))]
}
