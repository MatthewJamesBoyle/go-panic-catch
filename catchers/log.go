package catchers

import "fmt"

type Log struct{}

func (l Log) HandlePanic(message string) error {
	fmt.Print(message)
	return nil
}
