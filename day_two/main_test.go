package main

import (
	"testing"
)

type TestItem struct {
	valueForTest   string
	expectedResult bool
}

func TestPasswordPolicy(t *testing.T) {

	t.Run("Password inside the range should succedd", func(t *testing.T) {
		input := "1-3 a: abcde"

		got := SledRentalPasswordCheck(input)

		if got == false {
			t.Error("Policy didn't work")
		}

	})

	t.Run("Bellow the minimum number of letters should fail", func(t *testing.T) {
		input := "1-3 a: bcde"

		got := SledRentalPasswordCheck(input)

		if got == true {
			t.Error("Policy didn't work")
		}

	})

	t.Run("Password surpassing the number of letters should fail", func(t *testing.T) {
		input := "2-6 e: eeeeeeee"

		got := SledRentalPasswordCheck(input)

		if got == true {
			t.Error("Policy didn't work")
		}

	})

	t.Run("Wrong format should fail", func(t *testing.T) {
		input := "444444444"

		got := SledRentalPasswordCheck(input)

		if got == true {
			t.Error("Policy didn't work")
		}

	})

	t.Run("Empty string should fail", func(t *testing.T) {
		input := ""

		got := SledRentalPasswordCheck(input)

		if got == true {
			t.Error("Policy didn't work")
		}

	})

	t.Run("Multiple checks", func(t *testing.T) {
		multipleItems := []TestItem{
			{"9-14 b: bbbbbbbbbbbbbvb", true},
			{"3-5 r: trhrrrrrgrrn", false},
			{"7-20 r: rxrqlrwrnrsrrvbvtrrm", true},
			{"1-5 t: sttttttttvt", false},
		}

		for _, item := range multipleItems {
			got := SledRentalPasswordCheck(item.valueForTest)
			if got != item.expectedResult {
				t.Errorf("for the policy %s was expected %t but got %t", item.valueForTest, item.expectedResult, got)
			}
		}

	})
}

func TestPasswordPolicyByPosition(t *testing.T) {

	t.Run("One occurence of the right character on the right position should return true", func(t *testing.T) {
		input := "1-3 a: abcde"

		got := TobogganPasswordCheck(input)

		if got == false {
			t.Error("Policy didn't work")
		}

	})

	t.Run("Two occurencies of the right character should return false", func(t *testing.T) {
		input := "1-3 a: abade"

		got := TobogganPasswordCheck(input)

		if got == true {
			t.Error("Policy didn't work")
		}

	})

	t.Run("Zero occurencies of the right character should return false", func(t *testing.T) {
		input := "1-3 a: eeeee"

		got := TobogganPasswordCheck(input)

		if got == true {
			t.Error("Policy didn't work")
		}

	})

	t.Run("wrong format should return false", func(t *testing.T) {
		input := "wrong format"

		got := TobogganPasswordCheck(input)

		if got == true {
			t.Error("Policy didn't work")
		}

	})

}
