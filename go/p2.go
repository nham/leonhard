package main

import "fmt"

func main() {
    prev, sum, curr := 0, 0, 2

    for curr <= 4000000 {
        sum += curr
        prev, curr = curr, (prev + 4*curr)
    }

    fmt.Printf("Sum: %d\n", sum)

}

