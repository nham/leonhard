package main

import (
    "fmt"
    "math"
    "io/ioutil"
    "strings"
    "strconv"
)

func readseq() []uint {
    content, _ := ioutil.ReadFile("../p11.txt")
    lines := strings.Split(string(content), "\n")
    is := make([]uint, 400)

    for i := 0; i < len(lines) - 1; i++ {
        line := strings.Split(lines[i], " ")
        for j := 0; j < len(line); j++ {
            n, _ := strconv.Atoi(string(line[j]))
            is[20*i + j] = uint(n)
        }
    }

    return is

}

func maxseqinc(a []uint, length, width, inc int) int {
    max := 0
    var tmp int

    if length < width {
        return 0
    }

    for i := 0; i < length - width; i++ {
        tmp = 1
        for off := 0; off < width; off++ {
            tmp *= int(a[(i+off)*inc])
        }
        if tmp > max {
            max = tmp
            //fmt.Println(a[i*inc:(i+width)*inc])
            //fmt.Println(" max -- ", max)
        }
    }

    return max

}

func main() {
    a := readseq()
    max := 0.0

    for i := 0; i < 20; i++ {
        max = math.Max(max, float64(maxseqinc(a[i:], 20, 4, 20)))
        max = math.Max(max, float64(maxseqinc(a[i*20:], 20, 4, 1)))

        max = math.Max(max, float64(maxseqinc(a[i:], 20-i, 4, 21)))
        max = math.Max(max, float64(maxseqinc(a[i:], i+1, 4, 19)))
    }
    fmt.Println(int(max))
}
