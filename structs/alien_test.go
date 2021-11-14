package structs

import (
	"testing"
)

// Capitalises the first character of a string
func TestSpawnAliensZeroPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	SpawnAliens(0)
}
