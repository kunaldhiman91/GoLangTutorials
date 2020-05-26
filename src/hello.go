package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("hello, world\n")
	resultAdd := addNum(10, 20)
	fmt.Print("addition ==", resultAdd, "\n")
	resultDivide, err := divideNum(22.0, 11.0)
	fmt.Print("division == ", resultDivide, " error == ", err, "\n")
	resultErrorDivide, err := divideNum(11.0, 0.0)
	fmt.Print("division == ", resultErrorDivide, " error == ", err, "\n")

	dhawalDeveloper := developer{name: "DHawal Harami !!!!!", experience: 10}
	fmt.Println(dhawalDeveloper.name)
	fmt.Println(dhawalDeveloper.experience)

}

func addNum(num1 int, num2 int) int {
	return num1 + num2
}

//Similar to returning tuples in swift.
func divideNum(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Division by zero not permitted")
	}
	return (num1 / num2), nil
}

// Declaring struct
type developer struct {
	name       string
	experience int
}
