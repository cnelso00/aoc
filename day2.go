package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// code for stub mostly taken from LizTheGrey, but modified to remove bits I did not have context for,
// and to use my preferred naming conventions
func main() {
        bytes, err := ioutil.ReadFile("resources/day2.txt")
        if err != nil {
                return
        }
        contents := string(bytes)
        contents = strings.Trim(contents, "\n")
        lines := strings.Split(contents, "\n")

        var FORWARD = "forward"
        var DOWN = "down"
        var UP = "up"

        x_pos := 0
        y_pos := 0

        for _, line := range lines {
                vals := strings.Split(line, " ")
                cmd := vals[0]
                num, err := strconv.Atoi(vals[1])
                if (err != nil) {
                        fmt.Println(err)
                } else {
                        if (cmd == FORWARD) {
                                x_pos = x_pos + num
                        }
                        if (cmd == DOWN) {
                                y_pos = y_pos + num
                        }
                        if (cmd == UP) {
                                y_pos = y_pos - num
                        }
                }
        }
        final := x_pos * y_pos
        fmt.Println(final)

        x_pos = 0
        y_pos = 0
        aim := 0

        for _, line := range lines {
                vals := strings.Split(line, " ")
                cmd := vals[0]
                num, err := strconv.Atoi(vals[1])
                if (err != nil) {
                        fmt.Println(err)
                } else {
                        if (cmd == FORWARD) {
                                x_pos = x_pos + num
                                y_pos = y_pos + (aim * num)
                        }
                        if (cmd == DOWN) {
                                aim = aim + num
                        }
                        if (cmd == UP) {
                                aim = aim  - num
                        }
                }
        }
        final = x_pos * y_pos
        fmt.Println(final)
}

