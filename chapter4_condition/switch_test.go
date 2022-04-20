package chapter4_condition

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	var day string = "Monday"

	switch day {
	case "Monday":
		fmt.Println("Monday")
	case "Tuesday":
		fmt.Println("Tuesday")
	case "Wednesday":
		fmt.Println("Wednesday")
	case "Thursday":
		fmt.Println("Thursday")
	case "Friday":
		fmt.Println("Friday")
	case "Saturday":
		fmt.Println("Saturday")
	case "Sunday":
		fmt.Println("Sunday")
	default:
		fmt.Println("Default")
	}
}

func TestSwitchMul(t *testing.T) {
	var day string = "2"

	switch day {
	case "Monday":
		fmt.Println("Monday")
	case "Tuesday", "2":
		fmt.Println("Tuesday")
	case "Wednesday":
		fmt.Println("Wednesday")
	case "Thursday":
		fmt.Println("Thursday")
	case "Friday":
		fmt.Println("Friday")
	case "Saturday":
		fmt.Println("Saturday")
	case "Sunday":
		fmt.Println("Sunday")
	default:
		fmt.Println("Default")
	}
}
