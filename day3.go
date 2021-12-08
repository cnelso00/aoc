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

        bits := []int{0,0,0,0,0,0,0,0,0,0,0,0}
        for _, val := range lines {
                for i, char := range val {
                        if (string(char) == "0") {
                                bits[i]--
                        } else {
                                bits[i]++
                        }
                }
        }
       gamma := []int{0,0,0,0,0,0,0,0,0,0,0,0}
       epsilon := []int{0,0,0,0,0,0,0,0,0,0,0,0}

        for i, val := range bits {
                if val >  0 {
                        gamma[i] = 1
                        epsilon[i] = 0
                } else {
                        gamma[i] = 0
                        epsilon[i] = 1
                }
        }

        gamma_rate := 0
        epsilon_rate := 0
        pow := 1
        for i := len(bits) - 1; i >= 0; i-- {
                gamma_rate = gamma_rate + (gamma[i] * pow)
                epsilon_rate = epsilon_rate + (epsilon[i] * pow)
                pow = pow * 2
        }
        result := epsilon_rate * gamma_rate
        fmt.Println(gamma_rate)
        fmt.Println(epsilon_rate)
        fmt.Println(result)

        //num 2

        //get most common

        common := find_most_common(lines, "0")
        fmt.Println(common)

}

func find_most_common(list []string, char string) ([]string) {
        if (len(list) > 1) {
                common, other_list := split_by_leading_char(list, char)
                if (len(common) < len(other_list)) {
                        common = other_list
                }
                char := string(common[0][0])
                fmt.Println(char)
                list = find_most_common(common, char)
        }

        return list
}

func split_by_leading_char(list []string, char string) ([]string, []string) {
        char_list := []string{}
        other_list := []string{}

        for _, line := range list {
                if (string(line[0]) == char) {
                        char_list = append(char_list, line)
                } else {
                        other_list = append(other_list, line)
                }
        }

        return char_list, other_list
}

//slice off the first value in each line
func shorten_list(list []string) ([]string) {
        for i, val := range list {
                list[i] = val[1:]
        }
        return list
}
