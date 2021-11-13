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
	MAX_MOVES = 10
)

type Alien struct {
	Name        string
	InputChan   chan int
	CurrentCity *City
	trapped     bool
	dead        bool
	mu          sync.Mutex
}

func (a *Alien) String() string {
	return fmt.Sprintf("%v", a.Name)
}

func (a *Alien) Wander(wg *sync.WaitGroup) {
	fmt.Printf("Alien[%v] has started to Wander\n", a.Name)
	fmt.Printf("Alien[%v] => is in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, 0)

	for i := 1; i <= MAX_MOVES; i++ {

		a.CurrentCity.mu.Lock()
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		//Move to next posible City
		nextCity, direction := a.CurrentCity.RandomNeighbour()

		if nextCity != nil {

			if nextCity.Invader != nil && nextCity.Invader.Name != a.Name {
				nextCity.mu.Lock()
				nextCity.Invader.mu.Lock()
				//Get ready to fight the other baddie
				fmt.Printf("Alien[%v] => MOVED %v '%v' from '%v' === ITER:[%v]\n", a.Name, direction, nextCity.Name, a.CurrentCity.Name, i)
				fmt.Printf("*** Alien[%v] VS Alien[%v] in '%v' ***\n", a.Name, nextCity.Invader.Name, nextCity.Name)

				//Kill both aliens..
				a.dead = true
				fmt.Printf("Alien[%v] => is **DEAD** in '%v' === ITER:[%v]\n", a.Name, nextCity.Name, i)
				nextCity.Invader.dead = true

				// nextCity.Invader.dieCmdCh <- true
				a.CurrentCity.mu.Unlock()

				//DELETE CITY FROM WORLD
				nextCity.mu.Unlock()
				nextCity.Invader.mu.Unlock()

				break
			}

			a.CurrentCity.mu.Unlock()
			fmt.Printf("Alien[%v] => MOVED %v '%v' from '%v' === ITER:[%v]\n", a.Name, direction, nextCity.Name, a.CurrentCity.Name, i)
			a.CurrentCity = nextCity

		} else {
			a.trapped = true
			fmt.Printf("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
			a.CurrentCity.mu.Unlock()
			break
		}

		if a.dead {
			fmt.Printf("Alien[%v] => is **DEAD** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
			break
		}
		if a.trapped {
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
