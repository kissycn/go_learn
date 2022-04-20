package variable

import (
	"fmt"
	"testing"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

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
