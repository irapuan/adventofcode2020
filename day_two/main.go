package main

import (
	"regexp"
	"strconv"
	"strings"
)

func PasswordCheck(input string) bool {
	rgx := regexp.MustCompile(`(?P<minimum>\d{1})-(?P<maximum>\d{1}) (?P<letter>\w{1}): (?P<password>\w+)`)

	policy := rgx.FindStringSubmatch(input)

	policyMinimum, _ := strconv.Atoi(policy[1])
	policyMaximum, _ := strconv.Atoi(policy[2])
	policyLetter := policy[3]
	password := policy[4]

	lettersInPassword := strings.Count(password, policyLetter)

	if lettersInPassword >= policyMinimum && lettersInPassword <= policyMaximum {
		return true
	}

	return false

}
