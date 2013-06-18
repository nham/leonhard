package main

import "fmt"

func isPal(n int) bool {
    s := fmt.Sprintf("%d", n)
    runes := []rune(s)
    rev := make([]rune, len(runes))
    copy(rev, runes)

    for i, j := 0, len(rev) - 1; i < j; i, j = i+1, j-1 {
        rev[i], rev[j] = rev[j], rev[i]
    }
    return string(runes) == string(rev)

}

func has3DigitFactors(n int) (bool, int) {
    for i := 999; i > 100; i-- {
        if n % i == 0 && n/i < 1000 && n/i > 99 {
            return true, i
        }

    }

    return false, 0

}

func main() {
    for c := 999*999; c > 9900; c-- {
        if isPal(c) {
            h3DF, factor := has3DigitFactors(c)
            if h3DF {
                fmt.Printf("winner, winner: %d\n", factor)
                break
            }
        }
    }

}
