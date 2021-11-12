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

	neighbouringCity, neighboutDirection := roadData[1], roadData[0]

	//Add the Neigbhour City first in World Map if not already added
	_, ok := w[neighbouringCity]
	if !ok {
		w[neighbouringCity] = &city.City{Name: neighbouringCity}
	}
	neighbourEntry := w[neighbouringCity]

	//Add the City the World Map if not already added
	if entry, ok := w[cityName]; ok {

		//Update neighbour info
		//we Will re-add the Directional neigbhour if already added
		//If file has issues and direrction is repeated for a city
		//the last most value will be considered this way
		switch neighboutDirection {
		case "north":
			entry.North = city.Road{DirName: city.North, DestCity: neighbourEntry}
		case "south":
			entry.South = city.Road{DirName: city.South, DestCity: neighbourEntry}
		case "east":
			entry.East = city.Road{DirName: city.East, DestCity: neighbourEntry}
		case "west":
			entry.West = city.Road{DirName: city.West, DestCity: neighbourEntry}
		}
	} else {

		var cityData city.City
		neighbourEntry := w[neighbouringCity] // we have already made sure that neighbour is added
		switch neighboutDirection {
		case "north":
			cityData.North = city.Road{DirName: city.North, DestCity: neighbourEntry}
		case "south":
			cityData.South = city.Road{DirName: city.South, DestCity: neighbourEntry}
		case "east":
			cityData.East = city.Road{DirName: city.East, DestCity: neighbourEntry}
		case "west":
			cityData.West = city.Road{DirName: city.West, DestCity: neighbourEntry}
		}
		cityData.Name = cityName

		w[cityName] = &cityData
	}

}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}
