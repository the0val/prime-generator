package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func chkNum(n int64, primes []int64) bool {
	iSqrt := int64(math.Sqrt(float64(n)))
	for _, v := range primes {
		if n%v == 0 {
			return false
		}
		if v > iSqrt {
			break
		}
	}
	return true && (len(primes) < cap(primes))
}

func chkAll(primes *[]int64, c chan int64) {
	var routine = func(n int64) {
		if chkNum(n, *primes) {
			c <- n
		}
	}
	for n := (*primes)[len(*primes)-1] + 2; len(*primes) < cap(*primes); n += 2 {
		go routine(n)
	}
	close(c)
}

func main() {
	start := time.Now()
	var primArray = [1e6]int64{2, 3, 5, 7, 11}
	primes := primArray[:5]

	ch := make(chan int64, 1000)
	go chkAll(&primes, ch)
	for n := range ch {
		primes = append(primes, n)
	}
	dt := time.Since(start)
	fmt.Printf("Found %v primes in %v", len(primes), dt)
	f, _ := os.Create("primes-out")
	defer f.Close()
	for _, v := range primes {
		f.WriteString(fmt.Sprintf("%v, ", v))
	}
	f.WriteString(fmt.Sprintf("\nlen, %v\ndt, %v", len(primes), dt))
}
