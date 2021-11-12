package main

import (
	"alien-invasion/structs"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	MAX_MOVES = 1000
)

func main() {

	//Parse all input args
	var alien_count int
	var fps int64
	flag.IntVar(&alien_count, "aliens", 4, "Count of the Aliens that will descen upon the world")
	flag.Int64Var(&fps, "fps", 3, "How fast similaution should run per iteration in ms")
	filePathPtr := flag.String("world_file", "./test/world.txt", "Text file containing the world")
	flag.Parse()

	if alien_count < 2 {
		fmt.Println("Too few Aliens for simulation, what will a lonely Alien do...")
		os.Exit(1)
	}

	xWorld := structs.LoadWorldMap(*filePathPtr, alien_count)
	xWorld.BringInTheAliens(alien_count)
	start_simulation(xWorld, fps)
}

func start_simulation(xWorld structs.World, fps int64) {
	for i := 0; i <= MAX_MOVES; i++ {
		time.Sleep(time.Duration(fps) * time.Second)

		//Move Aliens

		fmt.Println(xWorld)
	}
}
