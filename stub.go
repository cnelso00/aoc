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
        contents = strings.Trim(contents, "\n")
        lines := strings.Split(contents, "\n")

}

