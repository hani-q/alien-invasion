package structs

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/hani-q/alien-invasion/util"
	log "github.com/sirupsen/logrus"
)

//This is used only to keep trakc of All spawn Aliens
//so that we can loop over these and instruct them to Wander
type AlienYellowPages map[string]*Alien

var Ayp AlienYellowPages = make(AlienYellowPages)

func (Ayp AlienYellowPages) String() string {
	keys := reflect.ValueOf(Ayp).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strings.Join(strkeys, ",")
}

//Map of the World.. in every sense of the word
type World map[string]*City

var XWorld = make(World)

//Prints the Map world in the same format as Input file
func (w World) String() string {

	var printData string

	if w == nil || len(w) == 0 {
		return "World is Empty!...Generate World Map first"
	} else {
		for _, cityData := range w {
			printData = printData + cityData.String()
		}

		return printData
	}
}

//Open the Map file path provided and populatre the World struct
func LoadWorldMap(fileName string, alien_count int) *World {
	file, err := os.Open(fileName)
	if isError(err) {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	for _, line := range fileLines {
		cityData := strings.Split(line, " ")

		// Making sure City name in file is always begings with a Cap
		cityName := util.Capitalise(cityData[0])

		for i := range cityData {
			if i > 0 {
				roadData := strings.Split(cityData[i], "=")

				//direction should always be lower case
				roadData[0] = strings.ToLower(roadData[0])

				// Making sure City name in file is always begins with a Cap
				roadData[1] = util.Capitalise(roadData[1])

				addCityToWorld(cityName, XWorld, roadData)
			}
		}
	}

	//Print Statistics about the cities
	log.Infof("%v cities have been loaded from map file", len(XWorld))

	return &XWorld
}

//Aliens are habiliated in the world. Each alien is provided with a caring city
//Each city is only one given one Alien
func (world World) PlaceTheAliens(alien_count int) {
	if len(world) < alien_count {
		msg := "Aliens cannot live in such a congested (simulated) world. Add more cities to World file"
		fmt.Println(msg)
		log.Error(msg)
		os.Exit(1)
	}

	//Open the dimensional portal and let Aliens descend from the Sky
	aliens := SpawnAliens(alien_count)

	//Rehabilitate displaced Aliens in EMPTY cities
	for _, alien := range aliens {
		for cityName := range world {
			//Add alien to cities which have 0 occupants
			if len(world[cityName].occupants) == 0 {

				world[cityName].AddOccupant(alien)
				log.Infof("Added %v to City %v", alien.Name, cityName)
				//Tell Alien also about the name of its City
				//Wish they were more smarter
				alien.CurrCityName = cityName

				//Update in YellowPages too
				//This will be used to tell all aliens to start
				Ayp[alien.Name] = alien
				break //inner loop break
			}
		}

	}
}

//The parsed cities that are read from map are added to the Map struct
func addCityToWorld(cityName string, w World, roadData []string) {

	neighbouringCityName, neighbourDirection := roadData[1], roadData[0]

	var currentCity, neighbourCity *City

	//Add the Neigbhour City first in World Map if not already added
	_, ok := w[neighbouringCityName]
	if !ok {
		w[neighbouringCityName] = &City{Name: neighbouringCityName, occupants: make(map[string]*Alien)}
	}
	neighbourCity = w[neighbouringCityName]

	//Add the City to the World Map if not already added
	//Repeated cities will be updated with latest info
	if entry, ok := w[cityName]; ok {
		currentCity = entry
	} else {
		var cityData City = City{Name: cityName, occupants: make(map[string]*Alien)}
		w[cityName] = &cityData
		currentCity = w[cityName]
	}

	//Update neighbour info
	//we Will re-add the Directional neigbhour if already added
	//If file has issues and direrction is repeated for a city
	//the last most value will be considered this way
	addNeighbourInfo(currentCity, neighbourCity, neighbourDirection)

	//Add reverse neigbourInfo
	// if Foo is to the South of Baz THEN Baz is to the North of Foo
	// if Baz is to the West of Bee THEN Bee is to the East of Baz

	addNeighbourInfo(neighbourCity, currentCity, ReverseStringDirecton(neighbourDirection))
}

//Adds the Neighout Info as Road struct for the passed Cardinal direction
func addNeighbourInfo(c *City, neigbourCity *City, neighboutDirection string) {
	switch neighboutDirection {
	case "north":
		c.North = &Road{DirName: North, DestCity: neigbourCity}
	case "south":
		c.South = &Road{DirName: South, DestCity: neigbourCity}
	case "east":
		c.East = &Road{DirName: East, DestCity: neigbourCity}
	case "west":
		c.West = &Road{DirName: West, DestCity: neigbourCity}
	}
}

//Wipes the world of the City. Does it kill the X-Worldings living there. Not sure
func (w World) DeleteCity(cityName string) {

	if entry, ok := w[cityName]; ok {

		//Check all of the Cities Roads and go to those Cities
		//and Delete the reverse road links

		if entry.North != nil {
			entry.North.DestCity.South = nil
			entry.North.DestCity = nil
		}

		if entry.South != nil {
			entry.South.DestCity.North = nil
			entry.South.DestCity = nil
		}

		if entry.East != nil {
			entry.East.DestCity.West = nil
			entry.East.DestCity = nil
		}

		if entry.West != nil {
			entry.West.DestCity.East = nil
			entry.West.DestCity = nil
		}
	}

	//Delete the City itSelf
	defer delete(w, cityName)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
