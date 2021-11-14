package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

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
		msg := fmt.Sprintf("too few aliens (%v) for simulation, what will a lonely alien do", alienCount)
		_ = fmt.Errorf(msg)
		panic(msg)
	}

	//Load the map.txt file into the Map pf world
	log.Info("Loading world from Map file", *inputMapPathStr)
	xWorld := structs.LoadWorldMap(*inputMapPathStr)

	//Spawn the Aliens and assign each to a unique and random city
	xWorld.PlaceTheAliens(alienCount)

	//Log the state of the Wold after the mayhem
	log.Info(xWorld)
	fmt.Printf("\nWorld in peacefull times\n\n")
	fmt.Println(xWorld)

	//Start the simulation
	startSimulation(numIterations)

	//Log the state of the Wold after the mayhem
	fmt.Printf("\nWhats left after the Mayhem\n\n")
	fmt.Println(xWorld)
	log.Info(xWorld)

	//Save it as final_map.txt

	//Printing Alien stats after the simulation ends
	var trappedAliens []string
	for alienName, alienData := range structs.Ayp {
		if alienData.Trapped {
			trappedAliens = append(trappedAliens, alienName)
		}
	}

	log.Infof("%v Aliens are still trapped and need rescue\nNames=%v", len(trappedAliens), strings.Join(trappedAliens, ", "))
}

//Starts the simulation by creating a WaitGroup for all the
//stored in Alien Yello Pages map
func startSimulation(maxMoves int) {
	var wg sync.WaitGroup

	for _, alienData := range structs.Ayp {
		log.Infof("Launching Alien=%v", alienData.Name)
		wg.Add(1)

		//Calling wander will make the Alien explore its
		//neighbouring cities
		go alienData.Wander(&wg, maxMoves)
	}

	//Wait for all Aliens to finish (Dear or Alive or Trapped)
	wg.Wait()
}
