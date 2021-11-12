package main

import (
	"alien-invasion/world"
	"flag"
	"fmt"
)

func main() {

	//Parse all input args
	var alien_count int
	flag.IntVar(&alien_count, "aliens", 4, "Count of the Aliens that will descen upon the world")
	filePathPtr := flag.String("world_file", "./test/world.txt", "Text file containing the world")
	flag.Parse()

	xWorld := world.LoadWorldMap(*filePathPtr)
	// fmt.Println(xWorld)

	// xWorld.DeleteCity("Foo")

	fmt.Println(xWorld)
}
