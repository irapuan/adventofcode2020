package secondpuzzle

import (
	"fmt"
	"testing"
)

func TestSecondPuzzle(t *testing.T) {

	t.Run("return three values that add 2020", func(t *testing.T) {
		input := `1721
979
366
299
675
1456`

		got := SecondPuzzleCalc(input)
		fmt.Printf("value returned : %d", got)
		if got != 241861950 {
			t.Errorf("expected 241861950, got %d", got)
		}
	})

}
