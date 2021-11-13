package structs

import (
	"alien-invasion/util"
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

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

type World map[string]*City

func (w World) String() string {

	var printData string

	if w == nil || len(w) == 0 {
		return "World is Empty!...Generate World Map first"
	} else {
		for _, cityData := range w {
			printData = printData + fmt.Sprintln(cityData)
		}

		return printData
	}
}

func LoadWorldMap(fileName string, alien_count int) World {
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

	world := make(World)

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

				addCityToWorld(cityName, world, roadData)
			}
		}
	}

	return world

}

func (world World) BringInTheAliens(alien_count int) {

	if len(world) < alien_count {
		fmt.Println("Aliens cannot live in such a congested (simulated) world..\n Add more cities to World file")
		os.Exit(1)
	}

	//Open the dimensional portal and let Zentardi descend from the Sky
	aliens := SpawnAliens(alien_count)

	//Rehabilitate displaced Zentardi in EMPTY cities
	for _, alien := range aliens {
		for cityName := range world {
			if world[cityName].Invaders[0] == nil {
				alien.CurrentCity = world[cityName]
				world[cityName].Invaders[0] = alien
				break
			}
		}

		Ayp[alien.Name] = alien
	}

}

func addCityToWorld(cityName string, w World, roadData []string) {

	neighbouringCityName, neighbourDirection := roadData[1], roadData[0]

	var currentCity, neighbourCity *City

	//Add the Neigbhour City first in World Map if not already added
	_, ok := w[neighbouringCityName]
	if !ok {
		w[neighbouringCityName] = &City{Name: neighbouringCityName}
	}
	neighbourCity = w[neighbouringCityName]

	//Add the City the World Map if not already added
	if entry, ok := w[cityName]; ok {

		currentCity = entry

	} else {
		var cityData City

		cityData.Name = cityName
		w[cityName] = &cityData
		currentCity = w[cityName]
	}

	//Update neighbour info
	//we Will re-add the Directional neigbhour if already added
	//If file has issues and direrction is repeated for a city
	//the last most value will be considered this way
	addNeighboutInfo(currentCity, neighbourCity, neighbourDirection)

	//Add reverse neigbourInfo
	// if Foo is to the South of Baz THEN Baz is to the North of Foo
	// if Baz is to the West of Bee THEN Bee is to the East of Baz
	addNeighboutInfo(neighbourCity, currentCity, ReverseStringDirecton(neighbourDirection))

}

func addNeighboutInfo(c *City, neigbourCity *City, neighboutDirection string) {
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

func (w World) DeleteCity(cityName string) {

	if entry, ok := w[cityName]; ok {

		//Check all of the Cities Roads and go to those Cities
		//and Delete the reverse road links

		if entry.North != nil {
			entry.North.DestCity = nil
		}

		if entry.South != nil {
			entry.South.DestCity = nil
		}

		if entry.East != nil {
			entry.East.DestCity = nil
		}

		if entry.West != nil {
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
