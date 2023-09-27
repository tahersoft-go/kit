package exp

import "testing"

func TestTerIf(t *testing.T) {
	t.Run("When condition is true", func(t *testing.T) {
		cond := true
		result := TerIf(cond, 1, 2)
		if result != 1 {
			t.Error("Result is not valid")
		}
	})

	t.Run("When condition is false", func(t *testing.T) {
		cond := false
		result := TerIf(cond, 1, 2)
		if result != 2 {
			t.Error("Result is not valid")
		}
	})
}
