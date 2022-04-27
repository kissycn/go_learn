package variable

import (
	"fmt"
	"testing"
)

type Weekday int

const (
	_      = iota // 0
	Sunday        // 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func TestIota2(t *testing.T) {
	fmt.Println(Sunday)
}

const (
	Apple, Banana     = iota, iota + 10 // 0, 10 (iota=0)
	Strawberry, Grape                   // 1, 11 (iota=1)
	Pear, Watermelon                    // 2, 12 (iota=2)
)

func TestIota1(t *testing.T) {
	fmt.Println(Apple, Banana)
	fmt.Println(Strawberry, Grape)
	fmt.Println(Pear, Watermelon)
}

func TestIotaConst(t *testing.T) {
	a := Weekday(Saturday)
	println(a)
	println(Sunday)

	var workday Weekday = Sunday
	println(workday)
}

type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}

	return "N/A"
}

func TestIota(t *testing.T) {
	var chipType ChipType = CPU
	fmt.Printf("%s %d", chipType, chipType)
}
