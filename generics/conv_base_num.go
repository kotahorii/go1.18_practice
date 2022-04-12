package generics

import (
	"log"
	"math"
	"strconv"
)

func ConvBinaryNum(n int) (answer string) {
	for n >= 1 {
		if n%2 == 0 {
			answer = "0" + answer
		} else {
			answer = "1" + answer
		}
		n /= 2
	}
	return
}

func ConvToTernary(n int) (answer string) {
	for n >= 1 {
		if n%3 == 0 {
			answer = "0" + answer
		} else if n%3 == 1 {
			answer = "1" + answer
		} else {
			answer = "2" + answer
		}
		n /= 3
	}
	return
}

func ConvToDecimal(n string) (answer int) {
	for i := 0; i < len(n); i++ {
		num, err := strconv.Atoi(string(n[i]))
		if err != nil {
			log.Fatal(err)
		}
		answer += num * int(math.Pow(float64(2), float64(len(n)-1-i)))
	}
	return
}
