package main

import "fmt"

//010 OMIT
type elevatorState int

const (
	elStopped = elevatorState(iota)
	elMovingUp
	elMovingDown
)

func (s elevatorState) String() string {
	switch s {
	case elStopped:
		return "Stopped"
	case elMovingUp:
		return "MovingUp"
	case elMovingDown:
		return "MovingDown"
	}
	return "UnknownState"
}

//020 OMIT
//030 OMIT
const elevatorMaxLoad = 1000

type elevatorEnv struct {
	currentState elevatorState
	currentLoad  int
}

func senseEnv() elevatorEnv {
	return elevatorEnv{}
}

//040 OMIT
//050 OMIT
// nextStateFirstTry returns the next state
func nextStateFirstTry(inp string, env elevatorEnv) elevatorState {
	switch inp {
	case "BtnUpPressed":
		return elMovingUp // HL
	default:
		return elStopped
	}
}

//060 OMIT
//070 OMIT
// nextStateTry2 returns the next state
func nextStateTry2(inp string, env elevatorEnv) elevatorState {
	if env.currentLoad > elevatorMaxLoad {
		alertError("Overload")
		return elStopped
	}

	switch inp {
	case "BtnUpPressed":
		if env.currentState != elMovingUp {
			cmdMoveUp()
			return elMovingUp
		}
		return elMovingUp
	default:
		alertError("Unknown Input")
		return elStopped
	}
	return elStopped
}

func cmdMoveUp() {
	fmt.Println("apply brakes")
	fmt.Println("release brakes")
	fmt.Println("motor direction up")
}
func alertError(e string) {
	fmt.Println(e)
}

//080 OMIT

func main() {
}
