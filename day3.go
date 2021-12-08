package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// code for stub mostly taken from LizTheGrey, but modified to remove bits I did not have context for,
// and to use my preferred naming conventions
func main() {
	bytes, err := ioutil.ReadFile("resources/day3.txt")
	if err != nil {
		return
	}
	contents := string(bytes)
	contents = strings.Trim(contents, "\n")
	lines := strings.Split(contents, "\n")

	bits := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, val := range lines {
		for i, char := range val {
			if string(char) == "0" {
				bits[i]--
			} else {
				bits[i]++
			}
		}
	}
	gamma := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	epsilon := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i, val := range bits {
		if val > 0 {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	gammaRate := 0
	epsilonRate := 0
	pow := 1
	for i := len(bits) - 1; i >= 0; i-- {
		gammaRate = gammaRate + (gamma[i] * pow)
		epsilonRate = epsilonRate + (epsilon[i] * pow)
		pow = pow * 2
	}
	result := epsilonRate * gammaRate
	fmt.Println(gammaRate)
	fmt.Println(epsilonRate)
	fmt.Println(result)

	//num 2

	//get most common

	common := findMostCommon(lines, "0")
	fmt.Println(common)

}

func findMostCommon(list []string, char string) []string {
	if len(list) > 1 {
		common, otherList := splitByLeadingChar(list, char)
		if len(common) < len(otherList) {
			common = otherList
		}
		char := string(common[0][0])
		fmt.Println(common)
		list = findMostCommon(shortenList(common), char)
	}

	return list
}

func splitByLeadingChar(list []string, char string) ([]string, []string) {
	var charList []string
	var otherList []string

	for _, line := range list {
		if string(line[0]) == char {
			charList = append(charList, line)
		} else {
			otherList = append(otherList, line)
		}
	}

	return charList, otherList
}

//slice off the first value in each line
func shortenList(list []string) []string {
	for i, val := range list {
		list[i] = val[1:]
	}
	return list
}
