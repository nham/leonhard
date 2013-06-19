package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func readseq() []uint {
    content, _ := ioutil.ReadFile("../p8.txt")
    lines := strings.Split(string(content), "\n")
    is := make([]uint, 1000)

    for i := 0; i < len(lines); i++ {
        for j := 0; j < len(lines[i]); j++ {
            n, _ := strconv.Atoi(string(lines[i][j]))
            is[50*i + j] = uint(n)
        }
    }

    return is

}


func maxseqinc(a []uint, width, inc int) int {
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
    fmt.Println(maxseqinc(readseq(), 5, 1))
}
