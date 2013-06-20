package main

import "fmt"

func divides(a, b int) bool {
    return a%b == 0
}

func isPrime(x int) bool {
    if divides(x, 2) {
        return false
    }

    for i := 3; i*i <= x; i += 2 {
        if divides(x, i) {
            return false
        }
    }

    return true
}

func main() {

    var sum int64 = 2
    for i := 3; i < 2000000; i += 2 {
        if isPrime(i) {
            sum += int64(i)
        }
    }

    fmt.Println(sum)
}
