package catchers

import (
	"fmt"
	"os"
)

//Log is a an empty struct that satisfies the PanicHandler interface.
type Log struct{}

//HandlePanic takes the message and logs it to fmt.println if an unhandled panic occurs within your server.
func (l Log) HandlePanic(message string) error {
	fmt.Fprintln(os.Stderr, message)
	return nil
}
