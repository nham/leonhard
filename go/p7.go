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

    count, i := 1, 3

    for {
        if isPrime(i) {
            count += 1
        }

        if count == 10001 {
            break
        } else {
            i += 2
        }
    }

    fmt.Println(i)
}

