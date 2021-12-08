package main

import "fmt"
import "os"
import "bufio"
import "log"
import "strconv"

func main() {
        input, err := os.Open("resources/day1.txt")
        if err != nil {
                log.Fatal(err)
        }
        defer input.Close()

        var scanner = bufio.NewScanner(input)

        var increases = 0;
        var prev = -1;

        for scanner.Scan() {
                var value, err = strconv.Atoi(scanner.Text())
                if err != nil {
                        log.Fatal(err)
                }
                if prev != -1 {
                        if prev < value {
                                increases++
                        }
                }
                prev = value;
        }
        fmt.Println(increases)
}

