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
	MAX_MOVES = 500
)

type Alien struct {
	Name        string
	InputChan   chan int
	CurrentCity *City
	Trapped     bool
	Dead        bool
}

func (a *Alien) String() string {
	return fmt.Sprintf("%v", a.Name)
}

func (a *Alien) Wander(wg *sync.WaitGroup) {
	fmt.Printf("Alien[%v] has started to Wander\n", a.Name)
	fmt.Printf("Alien[%v] => is in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, 0)

	for i := 1; i <= MAX_MOVES; i++ {

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		//Move to next posible City
		a.CurrentCity.mu.Lock()
		nextCity, direction := a.CurrentCity.RandomNeighbour()
		a.CurrentCity.mu.Unlock()

		if nextCity != nil {

			if nextCity.Invader != nil && nextCity.Invader.Name != a.Name {

				nextCity.mu.Lock()
				//Get ready to fight the other baddie
				fmt.Printf("Alien[%v] => MOVED %v '%v' from '%v' === ITER:[%v]\n", a.Name, direction, nextCity.Name, a.CurrentCity.Name, i)
				fmt.Printf("*** Alien[%v] VS Alien[%v] in '%v' ***\n", a.Name, nextCity.Invader.Name, nextCity.Name)

				//Kill both aliens..
				a.Dead = true
				fmt.Printf("Alien[%v] => is **DEAD** in '%v' === ITER:[%v]\n", a.Name, nextCity.Name, i)
				nextCity.Invader.Dead = true

				//DELETE CITY FROM WORLD
				XWorld.DeleteCity(nextCity.Name)
				nextCity.mu.Unlock()
				break
			}

			fmt.Printf("Alien[%v] => MOVED %v '%v' from '%v' === ITER:[%v]\n", a.Name, direction, nextCity.Name, a.CurrentCity.Name, i)
			a.CurrentCity = nextCity

		} else {
			a.Trapped = true
			fmt.Printf("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)

			break
		}

		if a.Dead {
			fmt.Printf("Alien[%v] => is **DEAD** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
			break
		}
		if a.Trapped {
			fmt.Printf("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
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
