package main

import (
    "fmt"
    "math"
)

func divides(a, b int) bool {
    return a%b == 0
}

// A slice of all prime factors, with repetitions
// prime_factorize(60) = [2, 2, 3, 5]
func prime_factorize(n int) []uint {
    factors := make([]uint, 20)
    i := 0

    for divides(n, 2) {
        n /= 2
        factors[i] = 2
        i += 1
    }

    for j := 3; n > 1 && j <= n; {
        if divides(n, j) && isPrime(j) {
            n /= int(j)
            factors[i] = uint(j)
            i += 1
        } else {
            j += 2
        }
    }

    return factors

}

func isPrime(x int) bool {
    for i := 2; i*i <= x; i += 2 {
        if divides(x, i) {
            return false
        }
    }

    return true
}

func tallyPowers(a []uint) map[int]uint {
    powers := make(map[int]uint)
    var p int

    for i := 0; a[i] != 0; i++ {
        p = int(a[i])
        _, ok := powers[p]

        if ok == false {
            powers[p] = 1
        } else {
            powers[p] += 1
        }
    }

    return powers
}

func main() {
    factorPowers := map[int]uint{
        2: 0, 3: 0, 5: 0, 7: 0, 11: 0, 13: 0, 17: 0, 19: 0}

    for i := 2; i <= 20; i++ {
        factors := prime_factorize(i)

        tally := tallyPowers(factors)
        for j := 0; j < 20; j++ {
            if tally[j] > factorPowers[j] {
                factorPowers[j] = tally[j]
            }

        }
    }

    fmt.Println(factorPowers)

    sum := 1
    for i := 0; i < 20; i++ {
        if factorPowers[i] != 0 {
            sum *= int(math.Pow(float64(i), float64(factorPowers[i])))
        }
    }

    fmt.Println(sum)
}
