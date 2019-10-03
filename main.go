package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed  = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start = flag.Int("start", 1, "minimum value of random")
	end   = flag.Int("end", 6, "maximum value of random")
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Фукнция должна вернуть число из интервала [l,r]
func randInterval(l, r int) int {
	return rand.Intn(abs(r-l)+1) + min(l, r)
}

func main() {
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	// Dice roll 1..6
	if *start > *end {
		fmt.Println("Ошибка! значение start должно быть меньше, чем значение end.")
	} else {
		fmt.Println(randInterval(*start, *end))
	}
}
