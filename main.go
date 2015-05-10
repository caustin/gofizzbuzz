package main

import (
	"flag"
	"fmt"
	"strconv"
)

// FizzBuzz extends int types to implement the fizz buzz protocol
type FizzBuzz int

func (x FizzBuzz) String() string {
	switch {
	case x%3 == 0 && x%5 == 0:
		return "FizzBuzz"
	case x%3 == 0:
		return "Fizz"
	case x%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(int(x))

}

func main() {
	startPtr := flag.Int("start", 1, "The starting value.")
	lastPtr := flag.Int("last", 100, "The last integer value")

	flag.Parse()
	for x := FizzBuzz(*startPtr); x <= FizzBuzz(*lastPtr); x++ {
		fmt.Println(x.String())
	}

}
