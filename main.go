package main

import (
	"alien-invasion/structs"
	"flag"
	"fmt"
	"os"
	"sync"
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
	start_simulation()
	fmt.Println(xWorld)

	//print Alien Stats

	var deadAliens []string
	var trappedAliens []string
	for alienName, alienData := range structs.Ayp {
		if alienData.Dead {
			deadAliens = append(deadAliens, alienName)
		}
		if alienData.Trapped {
			trappedAliens = append(trappedAliens, alienName)
		}
	}

	fmt.Println(deadAliens)
	fmt.Println(trappedAliens)
}

func start_simulation() {
	var wg sync.WaitGroup

	count := 0
	for _, alienData := range structs.Ayp {

		fmt.Println("Launching Alien", alienData.Name)
		wg.Add(1)
		go alienData.Wander(&wg)
		count++

	}

	wg.Wait()

}
