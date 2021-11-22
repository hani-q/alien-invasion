package structs

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

const ALIEN_TAG = "ALIEN"

type Status string

//Cardinal Directions Enum
const (
	HATCHED   Status = "HATCHED"
	EXHASUTED Status = "EXHASUTED"
	DEAD      Status = "DEAD"
	TRAPPED   Status = "TRAPPED"
)

type Alien struct {
	Name         string
	moveCount    int
	PersonalChan chan string //incoming channel
	CurrCityName string
	prevCityName string //Name of the previosuly viisted City
	status       Status //When no roads lead out of a city this is set

}

func (a *Alien) Hatch(maxMoves int, world *World, queenChan chan<- AlienLanguage) {
	if maxMoves == 0 {
		msg := fmt.Sprintf("numIterations (%v) cannot be 0", maxMoves)
		_ = fmt.Errorf(msg)
		panic(msg)
	}

	//Counter to track an Alien that is stuck between 2 cities
	a.status = HATCHED
	log.Infof("%v: [%v] in '%v' has Hatched", ALIEN_TAG, a.Name, a.CurrCityName)

	go func() {
		//Tick means alien should move
		//Random so that each alien moves differently form the other
		ticker := time.NewTicker(time.Duration(rand.Intn(500)) * time.Millisecond)
		for {
			select {
			case cmd := <-a.PersonalChan:
				log.Infof("%v: [%v] commanded to {%v}", ALIEN_TAG, a.Name, cmd)
				switch cmd {
				case "die":
					a.status = DEAD
					queenChan <- AlienLanguage{AlienName: a.Name, AlienStatus: a.status}
					return
				}

			case <-ticker.C:
				a.moveCount = a.moveCount + 1
				if a.moveCount == maxMoves {
					a.status = EXHASUTED
					log.Infof("%v: [%v] => is **EXHASUTED** in '%v', moves(%v)", ALIEN_TAG, a.Name, a.CurrCityName, a.moveCount)
					queenChan <- AlienLanguage{AlienName: a.Name, AlienStatus: a.status}
					return
				}

				//Move to next posible City
				//Get a random direction

				currCityPtr := world.GetCity(a.CurrCityName)
				if currCityPtr == nil {
					//City has been Destoryed
					return
				}
				currCityPtr.mu.Lock()

				//Check if Alien is Stuck moving Back and Forth between a City
				//Can happen if Alien is between 2 cities that are only linked to each other

				//This will return a nil if all roads destroyed OR
				//if Available Cities are full (and battle is ongoing in those cities)
				//means we are trapped

				nextCityPtr, nextDirection := currCityPtr.RandomNeighbour()

				//If Next City is Available lets move there
				//Take lock of Current City & next City
				//to avoid anyone else moving In & to avoid anyone else moving onm from nextCity

				if nextCityPtr == nil {

					//Update Alien to TRAPPED status
					//and end this Blood lust
					a.status = TRAPPED
					log.Infof("%v: [%v] => is **TRAPPED** in '%v', moves(%v)", ALIEN_TAG, a.Name, a.CurrCityName, a.moveCount)
					queenChan <- AlienLanguage{AlienName: a.Name, AlienStatus: a.status}
					return
				} else {
					nextCityPtr.mu.Lock()

					//We have 2 scenarios, next city has 0 occupants
					//OR it has 1 occupants

					nextAlien := nextCityPtr.Occupant

					if nextAlien == nil {
						//0 Occupants await us
						//Peacefull transfer to new city

						//Track the previous city.
						a.prevCityName = a.CurrCityName

						//Update next cities names
						a.CurrCityName = nextCityPtr.Name

						//Update the current cities occupants
						currCityPtr.Occupant = nil

						//Add urself to next Cities occupants
						nextCityPtr.Occupant = a

						log.Infof("%v [%v] => **MOVED** %v:%v from '%v', moves(%v)",
							ALIEN_TAG, a.Name, nextDirection, nextCityPtr.Name, currCityPtr.Name, a.moveCount)

						nextCityPtr.mu.Unlock()
						currCityPtr.mu.Unlock()

					} else {

						//This will end in a Bloody fight between the 2 that can end both of them and the next city
						//Remove all references of the Alien from the current city
						a.CurrCityName = ""

						//Update the current cities occupants
						currCityPtr.Occupant = nil

						//Add Alien to next cities occupants,Shouldnt make a difference
						//Add urself to next Cities occupants
						nextCityPtr.Occupant = nil

						log.Infof("%v [%v] => **FIGHTING** %v after moving %v:%v from '%v', moves(%v)",
							ALIEN_TAG, a.Name, nextAlien.Name, nextDirection, nextCityPtr.Name, currCityPtr.Name, a.moveCount)
						nextAlien.PersonalChan <- "die"

						fmt.Printf("%v has been destroyed by %v and %v\n", nextCityPtr.Name, a.Name, nextAlien.Name)
						a.status = DEAD
						log.Infof("%v [%v] & [%v] are **DEAD** and took '%v' with them", ALIEN_TAG, a.Name, nextAlien.Name, nextCityPtr.Name)

						//Initial Delete of City from XWorld
						world.DeleteCity(nextCityPtr.Name)
						queenChan <- AlienLanguage{AlienName: a.Name, AlienStatus: a.status}
						nextCityPtr.mu.Unlock()
						currCityPtr.mu.Unlock()
						return
					}
				}
			}

		}
	}()

}
