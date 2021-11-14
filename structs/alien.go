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

type Alien struct {
	Name         string
	CurrCityName string
	prevCityName string //Name of the previosuly viisted City
	Trapped      bool   //When no roads lead out of a city this is set
}

func (a *Alien) String() string {
	return fmt.Sprintf("%v: Address=%v, Trapped=%v", a.Name, a.CurrCityName, a.Trapped)
}

func (a *Alien) Wander(wg *sync.WaitGroup, maxMoves int) {
	//Counter to track an Alien that is stuck between 2 cities
	//After counter is 0. Alien wont move anymore

	log.Infof("Alien[%v] in '%v' has started to Wander\n", a.Name, a.CurrCityName)
	for i := 1; i <= maxMoves; i++ {

		//Sleep a Random amount to mimic each alien moving to a new city
		//faster or slower
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		//Check if City exist in World Map
		if _, ok := XWorld[a.CurrCityName]; !ok {
			//City has been Destoryed.. Alien is also Dead
			break
		}

		//Move to next posible City
		//Get a random direction
		currCityPtr := XWorld[a.CurrCityName]

		//Check if Alien is Stuck moving Back and Forth between a City
		//Can happen if Alien is between 2 cities that are only linked to each other

		//This will return a nil if all roads destroyed OR
		//if Available Cities are full (and battle is ongoing in those cities)
		//means we are trapped
		nextCityPtr, nextDirection := currCityPtr.RandomNeighbour()

		if nextCityPtr == nil {
			//Update Alien to TRAPPED status
			//and end this Blood lust
			a.Trapped = true
			log.Infof("Alien[%v] => is **TRAPPED** in '%v' === ITER:[%v]\n", a.Name, a.CurrCityName, i)
			break
		}

		//If Next City is Available lets move there
		//Take lock of Current City & next City
		//to avoid anyone else moving In & to avoid anyone else moving onm from nextCity
		currCityPtr.mu.Lock()
		nextCityPtr.mu.Lock()

		//We have 2 scenarios, next city has 0 occupants
		//OR it has 1 occupants

		occuantCount := nextCityPtr.CountOccupants()

		if occuantCount == 0 {
			//0 Occupants await us
			//Peacefull transfer

			//Track the previous city.
			a.prevCityName = a.CurrCityName

			//Update next cities names
			a.CurrCityName = nextCityPtr.Name

			//Update the current cities occupants
			currCityPtr.RemoveOccupant(a.Name)

			//Add urself to next Cities occupants
			nextCityPtr.AddOccupant(a)

			log.Infof("Alien[%v] => **MOVED** %v:%v from '%v' === ITER:[%v]\n",
				a.Name, nextDirection, nextCityPtr.Name, currCityPtr.Name, i)

			nextCityPtr.mu.Unlock()
			currCityPtr.mu.Unlock()

		} else if occuantCount == 1 {
			//This will end in a Bloody fight between the 2 that can end both of them and the next city
			//Remove all references of the Alien from the current city

			a.CurrCityName = ""

			//Update the current cities occupants
			currCityPtr.RemoveOccupant(a.Name)

			//Add Alien to next cities occupants,Shouldnt make a difference
			//Add urself to next Cities occupants
			nextCityPtr.AddOccupant(a)
			log.Infof("Alien[%v] => **MOVED** %v:%v from '%v' === ITER:[%v]\n",
				a.Name, nextDirection, nextCityPtr.Name, currCityPtr.Name, i)

			fmt.Printf("%v has been destroyed by %v\n", nextCityPtr.Name, nextCityPtr.occupants)
			log.Infof("Alien[%v] are **DEAD** and took '%v' with them\n", nextCityPtr.occupants, nextCityPtr.Name)

			//Release the locks
			nextCityPtr.mu.Unlock()
			currCityPtr.mu.Unlock()

			//Initial Delete of City from XWorld
			XWorld.DeleteCity(nextCityPtr.Name)

		}

	}

	wg.Done()
}

//Instanitiate the aliens with count provided in cli Args
//Each alien will be placed in a loving and caring city
//It will be made sure that no other alien will be present in the
//same city
func SpawnAliens(alienCount int) []*Alien {

	if alienCount < 2 {
		msg := "Error: Alient count cannot be less then 2"
		log.Error(msg)
		panic(msg)
	}

	//Get fancy alien names from this NameGenerator library
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	//Keep a track of Spawned Aliens
	//This struct will be used by start simulation function to
	//Launch the aliens
	var result []*Alien = make([]*Alien, alienCount)
	for i := 0; i < alienCount; i++ {
		result[i] = &Alien{Name: util.Capitalise(nameGenerator.Generate())}
	}
	return result
}
