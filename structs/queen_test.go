package structs

import (
	"fmt"
	"testing"
)

// Test when aliens are with less count
func TestLayEggsAliensZeroPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestLayEggsAliensZeroPanic")
		}
	}()

	alienCount := 0
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//We know the test file has 10 cities
	testMapFilePath := "../test/world_test_ten.txt"

	xWorld := LoadWorldMap(testMapFilePath)

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)
}

func TestLayEggsAliensOnePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestLayEggsAliensOnePanic")
		}
	}()

	alienCount := 1
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//We know the test file has 10 cities
	testMapFilePath := "../test/world_test_ten.txt"

	xWorld := LoadWorldMap(testMapFilePath)

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)
}

func TestLayEggsAliensMoreThenCitiesPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestLayEggsAliensMoreThenCitiesPanic")
		}
	}()

	alienCount := 11 //Should Panic as Cities are 10 and Aliens 11
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//We know the test file has 10 cities
	testMapFilePath := "../test/world_test_ten.txt"
	xWorld := LoadWorldMap(testMapFilePath)

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)
}

func TestHatchedChildCount(t *testing.T) {
	testMapFilePath := "../test/world_test_ten.txt"

	alienCount := 10
	xWorld := LoadWorldMap(testMapFilePath)
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)
	queen.HatchChildren(10, xWorld)

	hatchedCount := 0

	for _, cityData := range xWorld.Data {
		if cityData.Occupant != nil {
			if cityData.Occupant.status == HATCHED {
				hatchedCount++
			}
		}
	}

	if alienCount != hatchedCount {
		t.Fatalf(`TestHatchedChildCount(%v) = %v, %v, want %v, error`,
			alienCount, hatchedCount,
			fmt.Sprintf("Aliens HATCHED in world should be %v",
				alienCount), alienCount)
	}
}

func TestExhaustedChildCount(t *testing.T) {
	testMapFilePath := "../test/world_test_non_connected.txt"

	alienCount := 2
	xWorld := LoadWorldMap(testMapFilePath)
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)
	queen.HatchChildren(1, xWorld)
	queen.WaitChildren()

	exhaustedCount := 0

	for _, cityData := range xWorld.Data {
		if cityData.Occupant != nil {
			if cityData.Occupant.status == EXHASUTED {
				exhaustedCount++
			}
		}
	}

	if alienCount != exhaustedCount {
		t.Fatalf(`TestExhaustedChildCount(%v) = %v, %v, want %v, error`,
			alienCount, exhaustedCount,
			fmt.Sprintf("Aliens EXHAUSTED in world should be %v",
				alienCount), alienCount)
	}
}

func TestTrappedChildCount(t *testing.T) {
	testMapFilePath := "../test/world_test_trapped.txt"

	alienCount := 2
	xWorld := LoadWorldMap(testMapFilePath)
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	//Manually destory cities
	//To generate a trapped world
	xWorld.DeleteCity("5")
	xWorld.DeleteCity("4")
	xWorld.DeleteCity("2")

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)
	queen.HatchChildren(5, xWorld)
	queen.WaitChildren()

	trappedCount := 0

	for _, cityData := range xWorld.Data {
		if cityData.Occupant != nil {
			if cityData.Occupant.status == TRAPPED {
				trappedCount++
			}
		}
	}

	if alienCount != trappedCount {
		t.Fatalf(`TestTrappedChildCount(%v) = %v, %v, want %v, error`,
			alienCount, trappedCount,
			fmt.Sprintf("Aliens TRAPPED in world should be %v",
				alienCount), alienCount)
	}
}

func TestDeadChildCount(t *testing.T) {
	testMapFilePath := "../test/world_test_trapped.txt"

	alienCount := 2
	xWorld := LoadWorldMap(testMapFilePath)
	queen := Queen{Children: make(map[string]*Alien), QueenChan: make(chan AlienLanguage)}

	xWorld.DeleteCity("3")
	xWorld.DeleteCity("4")
	xWorld.DeleteCity("5")
	xWorld.DeleteCity("6")

	//Spew the Queen Mothers Eggs the Entire X-World
	queen.LayEggs(alienCount, xWorld)
	queen.HatchChildren(10, xWorld)
	queen.WaitChildren()

	deadCount := 0

	for _, alien := range queen.Children {
		if alien.status == DEAD {
			deadCount++
		}
	}

	if alienCount != deadCount {
		t.Fatalf(`TestDeadChildCount(%v) = %v, %v, want %v, error`,
			alienCount, deadCount,
			fmt.Sprintf("Aliens DEAD in world should be %v",
				alienCount), alienCount)
	}
}
