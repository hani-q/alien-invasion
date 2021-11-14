package structs

import (
	"testing"
)

// Test when aliens are with less count
func TestSpawnAliensZeroPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	SpawnAliens(0)
}
