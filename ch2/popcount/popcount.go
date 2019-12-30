package popcount

import (
	"fmt"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {
	var ret byte
	var i uint64
	for i = 0; i < 8; i++ {
		ret += pc[byte(x>>(i*8))]
	}
	return int(ret)
}

// PopCount returns the population count (number of set bits) of x.
func PopCount2(x uint64) byte {
	var ret byte
	var i uint8
	for i = 0; i < 64; i++ {
		ret += byte((x >> i) & 1)
	}
	return ret
}

func PopCount3(x uint64) byte {
	var ret byte
	for x > 0 {
		x = x & (x - 1)
		ret++
	}
	return ret
}

func Show(num uint64) {
	start := time.Now()
	fmt.Print(PopCount(num))
	fmt.Print(time.Since(start).Nanoseconds())
	fmt.Print("\n")
	start = time.Now()
	fmt.Print(PopCount(num))
	fmt.Print("\n")
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	PopCount(num)
	fmt.Print(time.Since(start).Nanoseconds())
	fmt.Print("\n")
	start = time.Now()
	fmt.Print(PopCount1(num))
	fmt.Print("\n")
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	PopCount1(num)
	fmt.Print(time.Since(start).Nanoseconds())
	fmt.Print("\n")
	start = time.Now()
	fmt.Print(PopCount2(num))
	fmt.Print("\n")
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	PopCount2(num)
	fmt.Print(time.Since(start).Nanoseconds())
	fmt.Print("\n")
	start = time.Now()
	fmt.Print(PopCount3(num))
	fmt.Print("\n")
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	PopCount3(num)
	fmt.Print(time.Since(start).Nanoseconds())
	fmt.Print("\n")
}
