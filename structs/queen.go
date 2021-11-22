package structs

import (
	"fmt"
	"strings"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/hani-q/alien-invasion/util"
	log "github.com/sirupsen/logrus"
)

const QUEEN_TAG = "QUEEN"

type AlienLanguage struct {
	AlienName   string
	AlienStatus Status
}

func (al AlienLanguage) String() string {
	return fmt.Sprintf("%v, status: %v", al.AlienName, al.AlienStatus)
}

type Queen struct {
	Children  map[string]*Alien
	QueenChan chan AlienLanguage
}

func (q *Queen) String() string {
	return fmt.Sprintf("%v[%v]", QUEEN_TAG, q.Children)
}

func (q *Queen) PrintStatus() {
	var data []string
	for _, child := range q.Children {
		data = append(data, fmt.Sprintf("%v={Moves: %v, Status: %v}", child.Name, child.moveCount, child.status))
	}
	fmt.Printf("Queen [%v]", strings.Join(data, ", "))
}

func (q *Queen) WaitChildren() {
	var total, pending int = len(q.Children), 0
	for d := range q.QueenChan {
		pending++
		log.Infof("%v: Child [%v] is Done (%v/%v)", QUEEN_TAG, d, pending, total)
		if total == pending {
			close(q.QueenChan)
		}
	}
}

func (q *Queen) HatchChildren(maxMoves int, world *World) {
	for alienName := range q.Children {
		q.Children[alienName].Hatch(maxMoves, world, q.QueenChan)
	}
}

func (q *Queen) LayEggs(childCount int, world *World) {

	if childCount < 2 {
		msg := "Error: Alient count cannot be less then 2"
		log.Error(msg)
		panic(msg)
	}

	if childCount > world.GetCityCount() {
		msg := fmt.Sprintf("aliens (%v) cannot be more then the cities(%v)", childCount, world.GetCityCount())
		_ = fmt.Errorf(msg)
		panic(msg)
	}

	//Get alien names from this NameGenerator library
	//Thats how they name themselves in thier world
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	for i := 0; i < childCount; i++ {
		alienName := util.Capitalise(nameGenerator.Generate())
		q.Children[alienName] = &Alien{Name: alienName, PersonalChan: make(chan string)}

		for cityName := range world.Data {
			//Add alien to cities which have 0 occupants
			if world.Data[cityName].Occupant == nil {

				world.Data[cityName].Occupant = q.Children[alienName]

				//Alien also about the name of its City
				//Wish they were more smarter
				q.Children[alienName].CurrCityName = cityName

				log.Infof("%v: Added [%v] to City %v", QUEEN_TAG, q.Children[alienName].Name, cityName)

				break //inner loop break
			}
		}
	}

	log.Infof("%v: layed %v Alien Eggs", QUEEN_TAG, len(q.Children))
}
