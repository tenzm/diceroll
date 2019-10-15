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
	n     = flag.Int("n", 1, "number of repeat")
	norepeat = flag.Bool("norepeat", false, "May repeat?")
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
	var repeated[100000] bool
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
		if (*norepeat && *n > *end-*start + 1){
			fmt.Println("Ошибка! Не возможно вывести", *n, "чисел в диапазоне от", *start,"до", *end,"без повторений.")
			return
		}
		for i := 0; i < *n; i++ {
			val := randInterval(*start, *end)
			if *norepeat && repeated[val] != true{
				repeated[val] = true
				fmt.Print(val)
				if i != *n-1 {
					fmt.Print(", ")
				}
			} else if *norepeat {
				i--
			} else {
				fmt.Print(val)
				if i != *n-1 {
					fmt.Print(", ")
				}
			}
		}
		fmt.Print("\n")
	}
}
