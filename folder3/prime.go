package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	f, _ := os.OpenFile("primes.txt", os.O_APPEND|os.O_WRONLY, 0644)
	for i := 1; i < 100000000; i++ {
		if checkPrime(uint32(i)) {
			defer f.Close()
			line := fmt.Sprintf("%v \n", i)
			_, _ = f.WriteString(line)
		}
	}
}
func checkPrime(n uint32) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%uint32(i) == 0 {
			return false
		}
	}
	return true
}
