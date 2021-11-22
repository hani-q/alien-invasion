package structs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/hani-q/alien-invasion/util"
	log "github.com/sirupsen/logrus"
)

const WORLD_TAG = "WORLD"

//Map of the World.. in every sense of the word
type World struct {
	Data map[string]*City
	mu   sync.Mutex // guards
}

//Prints the Map world in the same format as Input file
func (w *World) String() string {
	var printData string
	if w.Data == nil || len(w.Data) == 0 {
		msg := "World is Empty!...Generate World Map first"
		log.Warnf("%v: %v", QUEEN_TAG, msg)
		return msg
	} else {
		for _, cityData := range w.Data {
			printData = printData + cityData.String()
		}

		return printData
	}
}

//Count the cities. Lock first to avoid deletion during
//print phase
func (w *World) GetCityCount() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return len(w.Data)
}

//Get City Ptr by name
func (w *World) GetCity(name string) *City {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.Data[name]
}

//Open the Map file path provided and populatre the World struct
func LoadWorldMap(filePath string) *World {
	file, err := os.Open(filePath)
	if util.IsError(err) {
		msg := fmt.Sprintf("cannot open %v", filePath)
		_ = fmt.Errorf(msg)
		panic(msg)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	var xWorld = World{Data: make(map[string]*City)}

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

				addCityToWorld(cityName, &xWorld, roadData)
			}
		}
	}

	return &xWorld
}

//The parsed cities that are read from map are added to the Map struct
func addCityToWorld(cityName string, w *World, roadData []string) {

	w.mu.Lock()
	defer w.mu.Unlock()
	neighbouringCityName, neighbourDirection := roadData[1], roadData[0]

	var currentCity, neighbourCity *City

	//Add the Neigbhour City first in World Map if not already added
	_, ok := w.Data[neighbouringCityName]
	if !ok {
		w.Data[neighbouringCityName] = &City{Name: neighbouringCityName, Occupant: nil}
	}
	neighbourCity = w.Data[neighbouringCityName]

	//Add the City to the World Map if not already added
	//Repeated cities will be updated with latest info
	if entry, ok := w.Data[cityName]; ok {
		currentCity = entry
	} else {
		var cityData City = City{Name: cityName, Occupant: nil}
		w.Data[cityName] = &cityData
		currentCity = w.Data[cityName]
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
func (w *World) DeleteCity(cityName string) {

	w.mu.Lock()
	defer w.mu.Unlock()

	if entry, ok := w.Data[cityName]; ok {

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

	log.Infof("%v: Deleting city %v", QUEEN_TAG, cityName)

	//Delete the City itSelf
	delete(w.Data, cityName)
}
