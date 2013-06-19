package main

import "fmt"

    /* a^2 + b^2 = c^2 & a + b + c = M
     * => M(M - 2c) = 2ab
     * Since a, b > 0, M > 2c, or c < M/2
     *
     * Also, a^2 + (M - a - c)^2 = c^2
     * => 1/2 M^2 = (M-a)*(a+c)
     *
     * Note that this implies that M can never be odd. (M odd iff M^2 odd)
     * So M is even and 1/2 M^2 is even, which implies that if c is even,
     * a cannot be odd, since (a+c) would be odd, as would (M-a), and
     * odd*odd = even
     *
     * Perhaps the latter analysis is not really worthwhile, since it only
     * cuts down operations by 1/4 in the worst case (for even c we only need
     * examine the even a)
     */

func main() {
    M := 1000

    if M % 2 != 0 {
        return
    }

    var inc int
    for i := 1; i < M/2; i++ {
        if i % 2 == 0 {
            inc = 2
        } else {
            inc = 1
        }

        for a := inc; a <= (M-i)/2; a += inc {
            if M*M/2 == (M-a)*(a+i) {
                fmt.Printf("(%d, %d, %d) => %d\n", a, (M-a-i), i, a*i*(M-a-i))
            }
        }
    }

}
