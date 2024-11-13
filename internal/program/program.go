package program

import (
	"github.com/oleksandrcherevkov/cryptography/internal/commands"
)

func Start() {
	first := commands.DirectionCommand{}
	err := commands.Cycle(first)
	if err != nil {
		panic(err)
	}
}
