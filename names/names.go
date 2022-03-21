package names

import (
	"github.com/ilmaruk/cooltures"
	"math/rand"
	"time"
)

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

func (n Names) FullName(o Options) string {
	return n.FirstName(o) + " " + n.LastName(o)
}

func (n Names) FirstName(o Options) string {
	g := getGenerator(o.Culture, n.rand)
	return g.FirstName(o.Gender)
}

func (n Names) LastName(o Options) string {
	g := getGenerator(o.Culture, n.rand)
	return g.LastName()
}
