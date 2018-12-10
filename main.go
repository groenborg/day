package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"time"
)

func NewTimer(duration time.Duration) {

	startTime := time.Now()
	endTime := startTime.Add(duration)

	h := big.NewInt(int64(time.Until(endTime).Hours()))
	m := big.NewInt(int64(time.Until(endTime).Minutes()))
	s := big.NewInt(int64(time.Until(endTime).Seconds()))
	fmt.Printf("\rTimer:\t %02d:%02d:%02d", h, m, s)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			fmt.Printf("\rTimer:\t %02d:%02d:%02d", h, m, s)
		}
	}()

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var input string

	for input != "exit\n" {
		fmt.Println("again" + input)
		input, _ = reader.ReadString('\n')
		fmt.Println(input)
	}

	startTime := time.Now()
	endTime := startTime.Add(20)

	fmt.Printf("Start:\t %s \n", startTime.Format(time.ANSIC))
	fmt.Printf("End:\t %s\n", endTime.Format(time.ANSIC))

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {

			h := big.NewInt(int64(time.Until(endTime).Hours()))
			m := big.NewInt(int64(time.Until(endTime).Minutes()))
			s := big.NewInt(int64(time.Until(endTime).Seconds()))
			fmt.Printf("\rTimer:\t %02d:%02d:%02d", h, m, s)
		}
	}()

	time.Sleep(time.Second * 10)

	fmt.Println("Ticker stopped")

}
