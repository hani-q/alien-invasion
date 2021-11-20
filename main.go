package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/hani-q/alien-invasion/structs"
	log "github.com/sirupsen/logrus"
)

func main() {
	//Parsing input arguments
	var alienCount int
	var printStdout bool
	var numIterations int
	flag.IntVar(&alienCount, "aliens", 4, "Count of the Aliens that will descen upon the world. Default is 4")
	flag.BoolVar(&printStdout, "stdout", false, "Print output to screen instead of logging. Default is false")
	flag.IntVar(&numIterations, "numIterations", 10000, "Number of Iterations of the simulation. Each Alien will move these many Moves. Default is 10000")
	inputMapPathStr := flag.String("world_file", "./test/world_tiny.txt", "Text file containing the world. Defaults to ./test/world_tiny.txt")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	//InitalizeLogging
	if !printStdout {
		var file, err = os.OpenFile("build/logs.txt", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Could Not Open Log File : " + err.Error())
		}
		log.SetOutput(file)
	}
	log.SetFormatter(&log.TextFormatter{})

	//One alien will roam the world endlessly, since no City will be destroyed
	if alienCount < 2 {
		msg := fmt.Sprintf("too few aliens (%v) for simulation", alienCount)
		_ = fmt.Errorf(msg)
		panic(msg)
	}

	//Load the map.txt file into the Map pf world
	log.Info("Loading world from Map file", *inputMapPathStr)
	xWorld := structs.LoadWorldMap(*inputMapPathStr)

	//Print Statistics about the cities
	log.Infof("%v cities have been loaded from map file", xWorld.GetCityCount())

	if alienCount > xWorld.GetCityCount() {
		msg := fmt.Sprintf("aliens (%v) cannot be more then the cities(%v)", alienCount, xWorld.GetCityCount())
		_ = fmt.Errorf(msg)
		panic(msg)
	}

	fmt.Printf("\nWorld in peacefull times\n\n")
	fmt.Println(xWorld)

	//Bring in the Queen via the inter-dimensional portal
	//She will arrive in the Eaths outer-orbit
	queen := structs.Queen{Children: make(map[string]*structs.Alien), QueenChan: make(chan structs.AlienLanguage)}

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)

	queen.HatchChildren(numIterations)
	queen.WaitChildren()

	//Log the state of the World after the mayhem
	fmt.Printf("\nWhats left after the Mayhem\n\n")
	fmt.Println(xWorld)

	fmt.Printf("\nWhats left of the Invading Queen and his children\n\n")
	queen.PrintStatus()
}
