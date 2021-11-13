package structs

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/hani-q/alien-invasion/util"
	log "github.com/sirupsen/logrus"
)

const (
	MAX_MOVES = 10000 //Reduce to a lower value for testing
)

type Alien struct {
	Name        string
	CurrentCity *City //Current City the Alien is occupying
	Trapped     bool  //When no roads lead out of a city this is set
	Dead        bool
}

func (a *Alien) String() string {
	return fmt.Sprintf("%v", a.Name)
}

//Allows the Alien to wander and randomly a cardinal direction leading out of his currentCity
//occupied city. ALien will move MAX_MOVES until he is has run out of Moved OR Trapped OR Dead
func (a *Alien) Wander(wg *sync.WaitGroup) {

	log.Infof("Alien[%v] has started to Wander\n", a.Name)

	//Iteration Zero, Log the current City of the Alien
	log.Infof("Alien[%v] => is in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, 0)

	for i := 1; i <= MAX_MOVES; i++ {

		//Sleep a Random amount to make the logging/output more readable
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		//Move to next posible City
		//Get a random direction
		nextCity, direction := a.CurrentCity.RandomNeighbour()

		if nextCity != nil {
			if nextCity.Invader != nil && nextCity.Invader.Name != a.Name {
				//Lock the currently occipied City to Prevent another Alien moving in
				a.CurrentCity.mu.Lock()
				nextCity.mu.Lock()
				//Get ready to fight the other baddie
				log.Infof("Alien[%v] => MOVED %v '%v' from '%v' === ITER:[%v]\n", a.Name, direction, nextCity.Name, a.CurrentCity.Name, i)
				log.Infof("*** Alien[%v] VS Alien[%v] in '%v' ***\n", a.Name, nextCity.Invader.Name, nextCity.Name)

				//Kill both aliens..
				a.Dead = true
				log.Infof("Alien[%v] => is **DEAD** in '%v' === ITER:[%v]\n", a.Name, nextCity.Name, i)
				nextCity.Invader.Dead = true

				//DELETE CITY FROM WORLD
				XWorld.DeleteCity(nextCity.Name)

				fmt.Printf("%v has been destroyed by alien %v and alien %v\n", nextCity.Name, a.Name, nextCity.Invader)

				nextCity.mu.Unlock()
				a.CurrentCity.mu.Unlock()
				break
			}

			log.Infof("Alien[%v] => MOVED %v '%v' from '%v' === ITER:[%v]\n", a.Name, direction, nextCity.Name, a.CurrentCity.Name, i)
			a.CurrentCity = nextCity

		} else {
			a.CurrentCity.mu.Lock()
			a.Trapped = true
			log.Infof("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
			a.CurrentCity.mu.Unlock()
			break
		}

		if a.Dead {
			log.Infof("Alien[%v] => is **DEAD** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
			break
		}
		if a.Trapped {
			log.Infof("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrentCity.Name, i)
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
