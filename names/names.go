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

func (n Names) FullName(o Options) (string, error) {
	first, err := n.FirstName(o)
	if err != nil {
		return "", err
	}
	last, err := n.LastName(o)
	if err != nil {
		return "", err
	}
	return first + " " + last, nil
}

func (n Names) FirstName(o Options) (string, error) {
	g, err := getGenerator(o.Culture, n.rand)
	if err != nil {
		return "", err
	}
	return g.FirstName(o.Gender), nil
}

func (n Names) LastName(o Options) (string, error) {
	g, err := getGenerator(o.Culture, n.rand)
	if err != nil {
		return "", err
	}
	return g.LastName(), nil
}

func pickName(src []string, r cooltures.Randomiser) string {
	return src[r.Intn(len(src))]
}
