package play

import "errors"

func Publish(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("no args")
	}
	return args[1], nil
}