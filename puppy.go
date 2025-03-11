package puppy

import "fmt"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	fmt.Print("Barks: ")
	return "Woof! Woof! Woof!"
}
