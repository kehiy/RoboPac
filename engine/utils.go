package engine

import "fmt"

func CheckArgs(requiredArg int, args []string) error {
	if len(args) != requiredArg {
		return fmt.Errorf("incorrect number of arguments, expected %d but got %d", requiredArg, len(args))
	}
	return nil
}
