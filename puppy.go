package puppy

import (
	"fmt"

	"github.com/Leonardo-almd/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	fmt.Print("Barks: ")
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp("Woof!")
}
