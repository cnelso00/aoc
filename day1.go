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
        bytes, err := ioutil.ReadFile("resources/day1.txt")
        if err != nil {
                return
        }
        contents := string(bytes)
        lines := strings.Split(contents, "\n")
        //part 1

        var increases = 0;
        var prev = -1;

        for _, i := range lines {
                value, err := strconv.Atoi(i)
                if (err != nil) {
                        fmt.Println("hit an error")
                        fmt.Println(err)
                } else {
                        if prev != -1 {
                                if prev < value {
                                        increases++
                                }
                        }
                        prev = value
                }
        }

        fmt.Println(increases)


        //part 1
        increases = 0;
        nums := []int{0, 0, 0}
        prev = 0
        for i, content  := range lines {
                value, err := strconv.Atoi(content)
                if (err != nil) {
                        fmt.Println("hit an error")
                        fmt.Println(err)
                } else {
                        curr := prev + value - nums[0]
                        nums = append(nums[1:], value)
                        if (i > 3 && prev < curr) {
                                increases++;
                        }
                        prev = curr

                }
        }
        fmt.Println(increases)
}

