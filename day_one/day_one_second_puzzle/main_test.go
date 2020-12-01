package main

import "testing"

func SecondPuzzleTest(t *testing.T) {

	t.Run("return three values that add 2020", func(t *testing.T) {
		input := `1721
979
366
299
675
1456`

		result := SecondPuzzleCalc(input)
		if result != 241861950 {
			t.Errorf("expected 241861950, got %d", result)
		}
	})

}
