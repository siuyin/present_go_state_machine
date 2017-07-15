package main

import "fmt"

func Example_btnUpStopped() {
	env := elevatorEnv{currentState: elStopped, currentLoad: 100}
	inp := "BtnUpPressed"
	opState := nextStateTry2(inp, env)
	fmt.Println(opState)

	// Output:
	// apply brakes
	// release brakes
	// motor direction up
	// MovingUp
}
func Example_btnUpMovingUp() {
	env := elevatorEnv{currentState: elMovingUp, currentLoad: 100}
	inp := "BtnUpPressed"
	opState := nextStateTry2(inp, env)
	fmt.Println(opState)

	// Output:
	// MovingUp
}
func Example_btnUpOverload() {
	env := elevatorEnv{currentState: elStopped, currentLoad: 1100}
	inp := "BtnUpPressed"
	opState := nextStateTry2(inp, env)
	fmt.Println(opState)

	// Output:
	// Overload
	// Stopped
}
