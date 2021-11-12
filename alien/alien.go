package alien

import (
	"alien-invasion/util"
	"fmt"
	"time"

	"github.com/goombaio/namegenerator"
)

type Alien struct {
	Name            string
	CurrentCityName string
	prevCityName    string
	num_moves       int
}

func (a *Alien) String() string {
	return fmt.Sprintf("%v", a.Name)
}

func SpawnAliens(alienCount int) []*Alien {

	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	var result []*Alien
	for i := 0; i < alienCount; i++ {
		result = append(result, &Alien{Name: util.Capitalise(nameGenerator.Generate())})
	}

	return result
}
