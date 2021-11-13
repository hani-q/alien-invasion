package structs

import (
	"alien-invasion/util"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/goombaio/namegenerator"
)

const (
	MAX_MOVES = 20
)

type Alien struct {
	Name        string
	InputChan   chan int
	CurrentCity *City
	trapped     bool
	dead        bool
}

func (a *Alien) String() string {
	return fmt.Sprintf("%v", a.Name)
}

func (a *Alien) Wander(wg *sync.WaitGroup) {
	fmt.Printf("Alien[%v] has started to Wander\n", a.Name)

	fmt.Printf("Alien[%v] => is in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, 1)
	for i := 1; i <= MAX_MOVES; i++ {

		a.CurrentCity.mu.Lock()

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		//Move to next posible City
		nextCity, direction := a.CurrentCity.RandomNeighbour()
		if nextCity != nil {

			a.CurrentCity.mu.Unlock()
			a.CurrentCity = nextCity
			fmt.Printf("Alien[%v] => MOVED %v '%v' === ITER:[%v]\n", a.Name, direction, a.CurrentCity.Name, i)
			if i == MAX_MOVES {
				fmt.Printf("Alien[%v] => is **TIRED** in %v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
			}
		} else {
			a.trapped = true
			fmt.Printf("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
			a.CurrentCity.mu.Unlock()
			break
		}
	}

	wg.Done()
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
