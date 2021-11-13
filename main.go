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
	flag.IntVar(&alienCount, "aliens", 4, "Count of the Aliens that will descen upon the world")
	filePathPtr := flag.String("world_file", "./test/world_tiny.txt", "Text file containing the world")
	flag.Parse()

	//InitalizeLogging
	// var file, err = os.OpenFile("build/logs.txt", os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println("Could Not Open Log File : " + err.Error())
	// }
	// log.SetOutput(file)
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})

	//One alien will roam the world endlessly, since no City will be destroyed
	if alienCount < 2 {
		_ = fmt.Errorf("too few aliens (%v) for simulation, what will a lonely alien do", alienCount)
		os.Exit(1)
	}

	//Load the map.txt file into the Map pf world
	log.Info("Loading world from Map file", *filePathPtr)
	xWorld := structs.LoadWorldMap(*filePathPtr, alienCount)

	//Spawn the Aliens and assign each to a unique and random city
	xWorld.BringInTheAliens(alienCount)

	//Log the state of the Wold before the mayhem
	log.Info(xWorld)

	//Start the simulation
	startSimulation()

	//Log the state of the Wold after the mayhem
	log.Info(xWorld)

	//Printing Alien stats after the simulation ends
	var deadAliens []string
	var trappedAliens []string
	for alienName, alienData := range structs.Ayp {
		if alienData.Dead {
			deadAliens = append(deadAliens, alienName)
		} else if alienData.Trapped {
			trappedAliens = append(trappedAliens, alienName)
		}
	}

	log.Infof("%v Aliens died in active combat\nName=%v", len(deadAliens), strings.Join(deadAliens, ", "))
	log.Infof("%v Aliens are still trapped and need rescue\nNames=%v", len(trappedAliens), strings.Join(trappedAliens, ", "))
}

//Starts the simulation by creating a WaitGroup for all the
//stored in Alien Yello Pages map
func startSimulation() {
	var wg sync.WaitGroup

	for _, alienData := range structs.Ayp {
		log.Infof("Launching Alien=%v", alienData.Name)
		wg.Add(1)

		//Calling wander will make the Alien explore its
		//neighbouring cities
		go alienData.Wander(&wg)
	}

	//Wait for all Aliens to finish (Dear or Alive or Trapped)
	wg.Wait()
}
