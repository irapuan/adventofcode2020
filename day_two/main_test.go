package main

import (
	"testing"
)

func TestPasswordPolicy(t *testing.T) {

	t.Run("Check password policy with 1-3 a's should succedd", func(t *testing.T) {
		input := "1-3 a: abcde"

		got := PasswordCheck(input)

		if got == false {
			t.Error("Policy didn't work")
		}

	})

	t.Run("Check password policy with 1-3 a's should fail", func(t *testing.T) {
		input := "1-3 a: bcde"

		got := PasswordCheck(input)

		if got == true {
			t.Error("Policy didn't work")
		}

	})
}
