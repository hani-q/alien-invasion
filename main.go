package main

import (
	"alien-invasion/structs"
	"flag"
	"fmt"
	"os"
)

func main() {

	//Parse all input args
	var alien_count int
	flag.IntVar(&alien_count, "aliens", 4, "Count of the Aliens that will descen upon the world")
	filePathPtr := flag.String("world_file", "./test/world.txt", "Text file containing the world")
	flag.Parse()

	if alien_count < 2 {
		fmt.Println("Too few Aliens for simulation, what will a lonely Alien do...")
		os.Exit(1)
	}

	xWorld := structs.LoadWorldMap(*filePathPtr, alien_count)
	xWorld.BringInTheAliens(alien_count)
	start_simulation(xWorld)
}

func start_simulation(xWorld structs.World) {
	for _, alienData := range structs.Ayp {
		alienData.Wander()
		break
	}

}
