package main

import "fmt"

func maxseqinc(a []uint, length, width, inc int) int {
    max := 0
    var tmp int

    for i := 0; i <= len(a) - 1 - inc*(width-1); i += inc {
        tmp = 1
        for off := 0; off < width; off += inc {
            tmp *= int(a[i+off])
        }
        if tmp > max {
            max = tmp
            fmt.Println(a[i:i+width])
        }
    }

    return max

}

func main() {
    a := readseq()
    max := 0

    for i := 0; i < 20; i++ {
        max = math.Max(max, maxseqinc(a, 20, 4, 1))
        max = math.Max(max, maxseqinc(a, 20, 4, 20))
        max = math.Max(max, maxseqinc(a, 20-i, 4, 21))
        max = math.Max(max, maxseqinc(a, i+1, 4, 19))

    }
}
