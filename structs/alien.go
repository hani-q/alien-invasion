package structs

import (
	"alien-invasion/util"
	"fmt"
	"math/rand"
	"time"

	"github.com/goombaio/namegenerator"
)

const (
	MAX_MOVES = 100
)

type Alien struct {
	Name         string
	InputChan    chan int
	CurrentCity  *City
	prevCityName string
	trapped      bool
	dead         bool
}

func (a *Alien) String() string {
	return fmt.Sprintf("%v", a.Name)
}

func (a *Alien) Wander() {
	fmt.Printf("Alien[%v] has started to Wander\n", a.Name)

	fmt.Printf("Alien[%v] => is in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, 1)
	for i := 1; i <= MAX_MOVES; i++ {

		// a.CurrentCity.mu.Lock()

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		//Move to next posible City
		nextCity, direction := a.CurrentCity.RandomNeighbour(a.prevCityName)
		if nextCity != nil {
			a.prevCityName = a.CurrentCity.Name
			a.CurrentCity = nextCity
			fmt.Printf("Alien[%v] => MOVED %v '%v' === ITER:[%v]\n", a.Name, direction, a.CurrentCity.Name, i)
		} else {
			a.trapped = true
			fmt.Printf("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
			break
		}
		// defer a.CurrentCity.mu.Unlock()

	}
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
