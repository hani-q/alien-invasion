package structs

import (
	"testing"
)

//Test if an Alien gets exhausted
func TestExhaustedAlien(t *testing.T) {
	testMapFilePath := "../test/world_test_trapped.txt"

	alien := Alien{Name: "Test", moveCount: 0, PersonalChan: make(chan string), CurrCityName: "1", status: HATCHED}
	xWorld := LoadWorldMap(testMapFilePath)
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	xWorld.DeleteCity("3")
	xWorld.DeleteCity("4")
	xWorld.DeleteCity("5")
	xWorld.DeleteCity("6")

	alien.Hatch(10, xWorld, queen.QueenChan)

	<-queen.QueenChan
	close(queen.QueenChan)
	//One alien should just be EXHASUTED

	if alien.status != EXHASUTED {
		t.Fatalf(`TestExhaustedAlien() = %v, %v, want %v, error`,
			EXHASUTED, alien.status,
			"Alien should be exhausted")
	}
}

//Test if an Alien gets trapped
func TestTrappedAlien(t *testing.T) {
	testMapFilePath := "../test/world_test_trapped.txt"

	alien := Alien{Name: "Test", moveCount: 0, PersonalChan: make(chan string), CurrCityName: "1", status: HATCHED}
	xWorld := LoadWorldMap(testMapFilePath)
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	xWorld.DeleteCity("5")
	xWorld.DeleteCity("4")
	xWorld.DeleteCity("2")

	alien.Hatch(10, xWorld, queen.QueenChan)

	<-queen.QueenChan
	close(queen.QueenChan)
	//One alien should just be EXHASUTED

	if alien.status != TRAPPED {
		t.Fatalf(`TestTrappedAlien() = %v, %v, want %v, error`,
			TRAPPED, alien.status,
			"Alien should be trapped")
	}
}

//Test if an Alien Dies
func TestDeadAlien(t *testing.T) {
	testMapFilePath := "../test/world_test_trapped.txt"

	alien := Alien{Name: "Test", moveCount: 0, PersonalChan: make(chan string), CurrCityName: "1", status: HATCHED}
	alien2 := Alien{Name: "Killer", moveCount: 0, PersonalChan: make(chan string), CurrCityName: "1", status: HATCHED}

	xWorld := LoadWorldMap(testMapFilePath)

	xWorld.DeleteCity("3")
	xWorld.DeleteCity("4")
	xWorld.DeleteCity("5")
	xWorld.DeleteCity("6")

	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	alien.Hatch(10, xWorld, queen.QueenChan)
	alien2.Hatch(10, xWorld, queen.QueenChan)

	<-queen.QueenChan
	<-queen.QueenChan
	close(queen.QueenChan)
	//One alien should just be EXHASUTED

	if alien.status != DEAD {
		t.Fatalf(`TestDeadAlien() = %v, %v, want %v, error`,
			EXHASUTED, alien.status,
			"Alien should be Dead")
	}

	if alien2.status != DEAD {
		t.Fatalf(`TestDeadAlien() = %v, %v, want %v, error`,
			EXHASUTED, alien.status,
			"Alien should be Dead")
	}
}

//Test if an alien dies when recving on channel
func TestSendDieToAlien(t *testing.T) {
	testMapFilePath := "../test/world_test_trapped.txt"

	alien := Alien{Name: "Test", moveCount: 0, PersonalChan: make(chan string), CurrCityName: "1", status: HATCHED}
	xWorld := LoadWorldMap(testMapFilePath)
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	xWorld.DeleteCity("5")
	xWorld.DeleteCity("4")
	xWorld.DeleteCity("2")

	alien.Hatch(50, xWorld, queen.QueenChan)

	alien.PersonalChan <- "die"
	<-queen.QueenChan
	close(queen.QueenChan)
	//One alien should just be EXHASUTED

	if alien.status != DEAD {
		t.Fatalf(`TestSendDieToAlien() = %v, %v, want %v, error`,
			DEAD, alien.status,
			"Alien should be dead")
	}
}
