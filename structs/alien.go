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

type AlienCommChan struct {
	cmd       string
	alienName string
}

type Alien struct {
	Name        string
	InputChan   chan int
	CurrentCity *City
	trapped     bool
	dead        bool
	exhausted   bool
}

func (a *Alien) String() string {
	return fmt.Sprintf("%v", a.Name)
}

func (a *Alien) Wander(wg *sync.WaitGroup) {
	fmt.Printf("Alien[%v] has started to Wander\n", a.Name)
	fmt.Printf("Alien[%v] => is in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, 0)

	select {
	default:
		for i := 1; i <= MAX_MOVES; i++ {
			if a.dead {
				fmt.Printf("Alien[%v] => is **DEAD** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
				break
			}

			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			a.CurrentCity.mu.Lock()

			//Move to next posible City
			nextCity, direction := a.CurrentCity.RandomNeighbour()
			if nextCity != nil {

				if nextCity.Invader != nil && nextCity.Invader.Name != a.Name {
					//Get ready to fight the other baddie
					fmt.Printf("Alien[%v] => MOVED %v '%v' from '%v' === ITER:[%v]\n", a.Name, direction, nextCity.Name, a.CurrentCity.Name, i)
					fmt.Printf("*** Alien[%v] VS Alien[%v] in '%v' ***\n", a.Name, nextCity.Invader.Name, nextCity.Name)
					fmt.Printf("Alien[%v] => is **DEAD** in '%v' === ITER:[%v]\n", a.Name, nextCity.Name, i)
					//Kill both aliens..
					a.dead = true
					nextCity.Invader.dead = true
					a.CurrentCity.mu.Unlock()
					break
				}

				a.CurrentCity.mu.Unlock()
				fmt.Printf("Alien[%v] => MOVED %v '%v' from '%v' === ITER:[%v]\n", a.Name, direction, nextCity.Name, a.CurrentCity.Name, i)
				a.CurrentCity = nextCity

				if i == MAX_MOVES {
					fmt.Printf("Alien[%v] => is **TIRED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
					a.exhausted = true
				}
			} else {
				a.trapped = true
				fmt.Printf("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
				a.CurrentCity.mu.Unlock()
				break
			}
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
