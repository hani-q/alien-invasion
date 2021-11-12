package world

import (
	"alien-invasion/city"
	"alien-invasion/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type World map[string]*city.City

var XWorld World

func (w World) String() string {

	var printData string

	if w == nil || len(w) == 0 {
		return "World is Empty!..Generate World Map first"
	} else {
		for _, cityData := range w {

			printData = printData + fmt.Sprintln(cityData)
		}

		return printData
	}
}

func LoadWorldMap(fileName string) World {
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

				// Making sure City name in file is always begings with a Cap
				roadData[1] = util.Capitalise(roadData[1])

				addCityToWorld(cityName, world, roadData)
			}
		}
	}

	return world

}

func addCityToWorld(cityName string, w World, roadData []string) {

	neighbouringCityName, neighbourDirection := roadData[1], roadData[0]

	var currentCity, neighbourCity *city.City

	//Add the Neigbhour City first in World Map if not already added
	_, ok := w[neighbouringCityName]
	if !ok {
		w[neighbouringCityName] = &city.City{Name: neighbouringCityName}
	}
	neighbourCity = w[neighbouringCityName]

	//Add the City the World Map if not already added
	if entry, ok := w[cityName]; ok {

		currentCity = entry

	} else {
		var cityData city.City

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
	addNeighboutInfo(neighbourCity, currentCity, city.ReverseStringDirecton(neighbourDirection))

}

func addNeighboutInfo(c *city.City, neigbourCity *city.City, neighboutDirection string) {
	switch neighboutDirection {
	case "north":
		c.North = city.Road{DirName: city.North, DestCity: neigbourCity}
	case "south":
		c.South = city.Road{DirName: city.South, DestCity: neigbourCity}
	case "east":
		c.East = city.Road{DirName: city.East, DestCity: neigbourCity}
	case "west":
		c.West = city.Road{DirName: city.West, DestCity: neigbourCity}
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
