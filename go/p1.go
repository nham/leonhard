package main

import "fmt"

func divides(a, b int) bool {
    return a%b == 0

}

func main() {
    d1, d2 := 3, 5
    sum := 0

    for i := 1; i < 1000; i++ {
        if divides(i, d1) || divides(i, d2) {
            sum += i
        }
    }
    fmt.Printf("Sum of numbers < 1000 that are multiples of 3 and 5: %d\n", sum)
}
