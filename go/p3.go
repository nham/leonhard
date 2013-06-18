package main

import "fmt"

func divides(a, b int) bool {
    return a%b == 0
}

func isPrime(x int) bool {
    for i := 2; i*i <= x; i += 2 {
        if divides(x, i) {
            return false
        }
    }

    return true
}

func main() {
    // Alternatively, using coreutils `factor`, type "factor 600851475143" ;)
    N := 600851475143
    n := N
    var bpf int

    if divides(N, 2) {
        n /= 2
        bpf = 2
    }

    for i := 3; n > 1 && i <= n; i += 2 {
        if divides(n, i) && isPrime(i) {
            n /= i
            bpf = i
        }

    }

    fmt.Printf("biggest prime factor: %d\n", bpf)
}
