package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(Door *Door) {
	PrintStr("Door Opening...\n")
	Door.state = true
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = false
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	return Door.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if !IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == true {
		CloseDoor(door)
	}
}
