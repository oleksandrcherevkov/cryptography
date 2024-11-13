package commands

type Command interface {
	Exec() (Command, error)
}

type ExitCommand struct {
}

func (c ExitCommand) Exec() (Command, error) {
	return nil, nil
}

type FunctionCommand func() (Command, error)

func (c FunctionCommand) Exec() (Command, error) {
	return c()
}

func Cycle(command Command) error {
	for {
		var err error
		command, err = command.Exec()
		if err != nil {
			return err
		}
		if command == nil {
			return nil
		}
	}
}
