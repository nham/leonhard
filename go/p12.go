package main

import "fmt"

func divides(a, b int) bool {
    return a%b == 0
}

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

func numFactors(n int) int {
    factors := prime_factorize(n)
    tally := tallyPowers(factors)

    total := 1
    for _, v := range tally {
        total *= int(v + 1)
    }

    return total

}

func main() {
    total := 1
    var num int

    for i := 1; total < 500; i++ {
        if i % 2 == 0 {
            total = numFactors(i/2) * numFactors(i+1)
        } else {
            total = numFactors(i) * numFactors((i+1)/2)
        }

        num = i*(i+1)/2
    }
        fmt.Println(num)
        fmt.Println(total, "\n")

}
